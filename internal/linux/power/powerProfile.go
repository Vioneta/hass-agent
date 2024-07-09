// Copyright (c) 2024 Joshua Rich <joshua.rich@gmail.com>
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

//nolint:exhaustruct
//revive:disable:unused-receiver
package power

import (
	"context"
	"fmt"

	"github.com/godbus/dbus/v5"

	"github.com/joshuar/go-hass-agent/internal/hass/sensor"
	"github.com/joshuar/go-hass-agent/internal/linux"
	"github.com/joshuar/go-hass-agent/internal/logging"
	"github.com/joshuar/go-hass-agent/pkg/linux/dbusx"
)

const (
	powerProfilesPath      = "/net/hadess/PowerProfiles"
	powerProfilesDest      = "net.hadess.PowerProfiles"
	powerProfilesInterface = "org.freedesktop.Upower.PowerProfiles"
	activeProfileProp      = "ActiveProfile"
)

type powerSensor struct {
	linux.Sensor
}

//nolint:exhaustruct
func newPowerSensor(sensorType linux.SensorTypeValue, sensorValue dbus.Variant) *powerSensor {
	newSensor := &powerSensor{}

	value, err := dbusx.VariantToValue[string](sensorValue)
	if err != nil {
		newSensor.Value = sensor.StateUnknown
	} else {
		newSensor.Value = value
	}

	newSensor.SensorTypeValue = sensorType
	newSensor.IconString = "mdi:flash"
	newSensor.SensorSrc = linux.DataSrcDbus
	newSensor.IsDiagnostic = true

	return newSensor
}

type profileWorker struct{}

//nolint:exhaustruct
func (w *profileWorker) Events(ctx context.Context) (chan sensor.Details, error) {
	sensorCh := make(chan sensor.Details)

	// Check for power profile support, exit if not available. Otherwise, send
	// an initial update.
	sensors, err := w.Sensors(ctx)
	if err != nil {
		close(sensorCh)

		return sensorCh, fmt.Errorf("cannot retrieve power profile: %w", err)
	}

	go func() {
		sensorCh <- sensors[0]
	}()

	triggerCh, err := dbusx.WatchBus(ctx, &dbusx.Watch{
		Bus:       dbusx.SystemBus,
		Names:     []string{dbusx.PropChangedSignal},
		Interface: dbusx.PropInterface,
		Path:      powerProfilesPath,
	})
	if err != nil {
		close(sensorCh)

		return sensorCh, fmt.Errorf("could not watch D-Bus for power profile updates: %w", err)
	}

	// Watch for power profile changes.
	go func() {
		logging.FromContext(ctx).Debug("Monitoring power profile.")

		defer close(sensorCh)

		for {
			select {
			case <-ctx.Done():
				logging.FromContext(ctx).Debug("Unmonitoring power profile.")

				return
			case event := <-triggerCh:
				props, err := dbusx.ParsePropertiesChanged(event.Content)
				if err != nil {
					logging.FromContext(ctx).Warn("Received unknown event from D-Bus.", "error", err.Error())

					continue
				}

				if profile, profileChanged := props.Changed[activeProfileProp]; profileChanged {
					sensorCh <- newPowerSensor(linux.SensorPowerProfile, profile)
				}
			}
		}
	}()

	return sensorCh, nil
}

func (w *profileWorker) Sensors(ctx context.Context) ([]sensor.Details, error) {
	profile, err := dbusx.GetProp[dbus.Variant](ctx,
		dbusx.SystemBus,
		powerProfilesPath,
		powerProfilesDest,
		powerProfilesDest+"."+activeProfileProp)
	if err != nil {
		return nil, fmt.Errorf("cannot retrieve a power profile from D-Bus: %w", err)
	}

	return []sensor.Details{newPowerSensor(linux.SensorPowerProfile, profile)}, nil
}

func NewProfileWorker() (*linux.SensorWorker, error) {
	return &linux.SensorWorker{
			WorkerName: "Power Profile Sensor",
			WorkerDesc: "Sensor to track the current power profile.",
			Value:      &profileWorker{},
		},
		nil
}
