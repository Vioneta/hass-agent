// Copyright (c) 2024 Joshua Rich <joshua.rich@gmail.com>
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

//revive:disable:unused-receiver
package agent

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net"
	"strings"
	"time"

	"github.com/go-resty/resty/v2"

	"github.com/joshuar/go-hass-agent/internal/device/helpers"
	"github.com/joshuar/go-hass-agent/internal/hass/sensor"
	"github.com/joshuar/go-hass-agent/internal/logging"
)

const (
	versionWorkerID    = "agent_version_sensor"
	externalIPWorkerID = "external_ip_sensor" //nolint:gosec // false positive
	deviceControllerID = "device_controller"

	ExternalIPUpdateInterval       = 5 * time.Minute
	ExternalIPUpdateJitter         = 10 * time.Second
	ExternalIPUpdateRequestTimeout = 15 * time.Second
)

var ipLookupHosts = map[string]map[int]string{
	"icanhazip": {4: "https://4.icanhazip.com", 6: "https://6.icanhazip.com"},
	"ipify":     {4: "https://api.ipify.org", 6: "https://api6.ipify.org"},
}

var (
	ErrInvalidIP     = errors.New("invalid IP address")
	ErrNoLookupHosts = errors.New("no IP lookup hosts found")
)

type VersionWorker struct {
	version
}

type ExternalIPWorker struct {
	client     *resty.Client
	cancelFunc context.CancelFunc
}

type deviceController map[string]*workerState

func (w deviceController) ID() string {
	return deviceControllerID
}

func (w deviceController) ActiveWorkers() []string {
	activeWorkers := make([]string, 0, len(w))

	for id, worker := range w {
		if worker.started {
			activeWorkers = append(activeWorkers, id)
		}
	}

	return activeWorkers
}

func (w deviceController) InactiveWorkers() []string {
	inactiveWorkers := make([]string, 0, len(w))

	for id, worker := range w {
		if !worker.started {
			inactiveWorkers = append(inactiveWorkers, id)
		}
	}

	return inactiveWorkers
}

func (w deviceController) Start(ctx context.Context, name string) (<-chan sensor.Details, error) {
	worker, exists := w[name]
	if !exists {
		return nil, ErrUnknownWorker
	}

	if worker.started {
		return nil, ErrWorkerAlreadyStarted
	}

	workerCh, err := w[name].Start(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not start worker: %w", err)
	}

	w[name].started = true

	return workerCh, nil
}

func (w deviceController) Stop(name string) error {
	// Check if the given worker ID exists.
	worker, exists := w[name]
	if !exists {
		return ErrUnknownWorker
	}
	// Stop the worker. Report any errors.
	if err := worker.Stop(); err != nil {
		return fmt.Errorf("error stopping worker: %w", err)
	}

	return nil
}

func (w deviceController) States(ctx context.Context) []sensor.Details {
	var sensors []sensor.Details

	for _, worker := range w.ActiveWorkers() {
		workerSensors, err := w[worker].Sensors(ctx)
		if err != nil {
			logging.FromContext(ctx).
				With(slog.String("controller", w.ID())).
				Debug("Could not retrieve worker sensors",
					slog.String("worker", w[worker].ID()),
					slog.Any("error", err))
		}

		sensors = append(sensors, workerSensors...)
	}

	return sensors
}

func (agent *Agent) newDeviceController(_ context.Context, prefs agentPreferences) SensorController {
	var worker worker

	controller := make(deviceController)

	// Set up sensor workers.
	worker = agent.newVersionWorker(prefs.AgentVersion())
	controller[worker.ID()] = &workerState{worker: worker}
	worker = agent.newExternalIPUpdaterWorker()
	controller[worker.ID()] = &workerState{worker: worker}

	return controller
}

func (w *VersionWorker) ID() string { return versionWorkerID }

func (w *VersionWorker) Stop() error { return nil }

func (w *VersionWorker) Start(_ context.Context) (<-chan sensor.Details, error) {
	sensorCh := make(chan sensor.Details)

	go func() {
		defer close(sensorCh)
		sensorCh <- w
	}()

	return sensorCh, nil
}

func (w *VersionWorker) Sensors(_ context.Context) ([]sensor.Details, error) {
	return []sensor.Details{&w.version}, nil
}

func (agent *Agent) newVersionWorker(value string) *VersionWorker {
	return &VersionWorker{version: version(value)}
}

// ID returns the unique string to represent this worker and its sensors.
func (w *ExternalIPWorker) ID() string { return externalIPWorkerID }

// Stop will stop any processing of sensors controlled by this worker.
func (w *ExternalIPWorker) Stop() error {
	w.cancelFunc()

	return nil
}

//nolint:mnd
func (w *ExternalIPWorker) Sensors(ctx context.Context) ([]sensor.Details, error) {
	sensors := make([]sensor.Details, 0, 2)

	for _, ver := range []int{4, 6} {
		ipAddr, err := w.lookupExternalIPs(ctx, w.client, ver)
		if err != nil || ipAddr == nil {
			logging.FromContext(ctx).
				With(slog.String("worker", externalIPWorkerID)).
				Log(ctx, logging.LevelTrace, "Looking up external IP failed.", slog.Any("error", err))

			continue
		}

		sensors = append(sensors, ipAddr)
	}

	return sensors, nil
}

func (w *ExternalIPWorker) Start(ctx context.Context) (<-chan sensor.Details, error) {
	sensorCh := make(chan sensor.Details)

	// Create a new context for the updates scope.
	updatesCtx, cancelFunc := context.WithCancel(ctx)
	// Save the context cancelFunc in the worker to be used as part of its
	// Stop() method.
	w.cancelFunc = cancelFunc

	updater := func(_ time.Duration) {
		sensors, err := w.Sensors(updatesCtx)
		if err != nil {
			logging.FromContext(ctx).
				With(slog.String("worker", externalIPWorkerID)).
				Debug("Could not get external IP.", slog.Any("error", err))
		}

		for _, s := range sensors {
			sensorCh <- s
		}
	}
	go func() {
		defer close(sensorCh)
		helpers.PollSensors(updatesCtx, updater, ExternalIPUpdateInterval, ExternalIPUpdateJitter)
	}()

	return sensorCh, nil
}

func (w *ExternalIPWorker) lookupExternalIPs(ctx context.Context, client *resty.Client, ver int) (*address, error) {
	for host, addr := range ipLookupHosts {
		logging.FromContext(ctx).With(slog.String("worker", externalIPWorkerID)).
			LogAttrs(ctx, logging.LevelTrace,
				"Fetching external IP.",
				slog.String("host", host),
				slog.String("method", "GET"),
				slog.String("url", addr[ver]),
				slog.Time("sent_at", time.Now()))

		resp, err := client.R().Get(addr[ver])
		if err != nil || resp.IsError() {
			return nil, fmt.Errorf("could not retrieve external v%d address with %s: %w", ver, addr[ver], err)
		}

		logging.FromContext(ctx).With(slog.String("worker", externalIPWorkerID)).
			LogAttrs(ctx, logging.LevelTrace,
				"Received external IP.",
				slog.Int("statuscode", resp.StatusCode()),
				slog.String("status", resp.Status()),
				slog.String("protocol", resp.Proto()),
				slog.Duration("time", resp.Time()),
				slog.String("body", string(resp.Body())))

		cleanResp := strings.TrimSpace(string(resp.Body()))

		a := net.ParseIP(cleanResp)
		if a == nil {
			return nil, ErrInvalidIP
		}

		return &address{addr: a}, nil
	}

	return nil, ErrNoLookupHosts
}

func (agent *Agent) newExternalIPUpdaterWorker() *ExternalIPWorker {
	return &ExternalIPWorker{
		client: resty.New().SetTimeout(ExternalIPUpdateRequestTimeout),
	}
}
