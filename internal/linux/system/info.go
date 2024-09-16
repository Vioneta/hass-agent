// Copyright (c) 2024 Joshua Rich <joshua.rich@gmail.com>
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

//revive:disable:unused-receiver
package system

import (
	"context"
	"log/slog"

	"github.com/joshuar/go-hass-agent/internal/device"
	"github.com/joshuar/go-hass-agent/internal/hass/sensor"
	"github.com/joshuar/go-hass-agent/internal/linux"
	"github.com/joshuar/go-hass-agent/internal/logging"
)

const (
	infoWorkerID = "system_info_sensors"
)

type infoWorker struct{}

func (w *infoWorker) Sensors(ctx context.Context) ([]sensor.Details, error) {
	var sensors []sensor.Details

	// Get distribution name and version.
	distro, version, err := device.GetOSDetails()
	if err != nil {
		logging.FromContext(ctx).
			With(slog.String("worker", infoWorkerID)).
			Warn("Could not retrieve distro details.", slog.Any("error", err))
	} else {
		sensors = append(sensors,
			&linux.Sensor{
				DisplayName:  "Distribution Name",
				UniqueID:     "distribution_name",
				Value:        distro,
				IsDiagnostic: true,
				IconString:   "mdi:linux",
				DataSource:   linux.DataSrcProcfs,
			},
			&linux.Sensor{
				DisplayName:  "Distribution Version",
				UniqueID:     "distribution_version",
				Value:        version,
				IsDiagnostic: true,
				IconString:   "mdi:numeric",
				DataSource:   linux.DataSrcProcfs,
			},
		)
	}

	// Get kernel version.
	kernelVersion, err := device.GetKernelVersion()
	if err != nil {
		logging.FromContext(ctx).
			With(slog.String("worker", infoWorkerID)).
			Warn("Could not retrieve kernel version.", slog.Any("error", err))
	} else {
		sensors = append(sensors,
			&linux.Sensor{
				DisplayName:  "Kernel Version",
				UniqueID:     "kernel_version",
				Value:        kernelVersion,
				IsDiagnostic: true,
				IconString:   "mdi:chip",
				DataSource:   linux.DataSrcProcfs,
			},
		)
	}

	return sensors, nil
}

func NewInfoWorker(_ context.Context) (*linux.OneShotSensorWorker, error) {
	worker := linux.NewOneShotWorker(infoWorkerID)
	worker.OneShotType = &infoWorker{}

	return worker, nil
}
