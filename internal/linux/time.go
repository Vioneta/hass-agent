// Copyright (c) 2023 Joshua Rich <joshua.rich@gmail.com>
//
// This software is released under the MIT License.
// https://opensource.org/licenses/MIT

package linux

import (
	"context"
	"time"

	"github.com/iancoleman/strcase"
	"github.com/joshuar/go-hass-agent/internal/hass"
	"github.com/lthibault/jitterbug/v2"
	"github.com/rs/zerolog/log"
	"github.com/shirou/gopsutil/v3/host"
)

//go:generate stringer -type=timeProp -output time_props.go -linecomment
const (
	boottime timeProp = iota + 1 // Last Reboot
	uptime                       // Uptime
)

type timeProp int

type timeSensor struct {
	prop  timeProp
	value string
}

func (m *timeSensor) Name() string {
	return m.prop.String()
}

func (m *timeSensor) ID() string {
	return strcase.ToSnake(m.prop.String())
}

func (m *timeSensor) Icon() string {
	return "mdi:restart"
}

func (m *timeSensor) SensorType() hass.SensorType {
	return hass.TypeSensor
}

func (m *timeSensor) DeviceClass() hass.SensorDeviceClass {
	switch m.prop {
	case uptime:
		return hass.Duration
	case boottime:
		return hass.Timestamp
	}
	return 0
}

func (m *timeSensor) StateClass() hass.SensorStateClass {
	return 0
}

func (m *timeSensor) State() interface{} {
	return m.value
}

func (m *timeSensor) Units() string {
	return ""
}

func (m *timeSensor) Category() string {
	return "diagnostic"
}

func (m *timeSensor) Attributes() interface{} {
	return nil
}

func TimeUpdater(ctx context.Context, status chan interface{}) {
	update := func() {
		status <- &timeSensor{
			prop:  uptime,
			value: getUptime(ctx),
		}

		status <- &timeSensor{
			prop:  boottime,
			value: getBoottime(ctx),
		}
	}

	update()
	ticker := jitterbug.New(
		time.Minute*15,
		&jitterbug.Norm{Stdev: time.Minute},
	)
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				update()
			}
		}
	}()
}

func getUptime(ctx context.Context) string {
	u, err := host.UptimeWithContext(ctx)
	if err != nil {
		log.Debug().Caller().Err(err).
			Msg("Failed to retrieve uptime.")
		return "Unknown"
	}
	epoch := time.Unix(0, 0)
	uptime := time.Unix(int64(u), 0)
	return uptime.Sub(epoch).String()
}

func getBoottime(ctx context.Context) string {
	u, err := host.BootTimeWithContext(ctx)
	if err != nil {
		log.Debug().Caller().Err(err).
			Msg("Failed to retrieve boottime.")
		return "Unknown"
	}
	return time.Unix(int64(u), 0).Format(time.RFC3339)
}