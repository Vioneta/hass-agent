// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package agent

import (
	"context"
	"github.com/joshuar/go-hass-agent/internal/hass/sensor"
	"sync"
)

// Ensure, that SensorControllerMock does implement SensorController.
// If this is not the case, regenerate this file with moq.
var _ SensorController = &SensorControllerMock{}

// SensorControllerMock is a mock implementation of SensorController.
//
//	func TestSomethingThatUsesSensorController(t *testing.T) {
//
//		// make and configure a mocked SensorController
//		mockedSensorController := &SensorControllerMock{
//			ActiveWorkersFunc: func() []string {
//				panic("mock out the ActiveWorkers method")
//			},
//			InactiveWorkersFunc: func() []string {
//				panic("mock out the InactiveWorkers method")
//			},
//			StartFunc: func(ctx context.Context, name string) (<-chan sensor.Details, error) {
//				panic("mock out the Start method")
//			},
//			StartAllFunc: func(ctx context.Context) (<-chan sensor.Details, error) {
//				panic("mock out the StartAll method")
//			},
//			StopFunc: func(name string) error {
//				panic("mock out the Stop method")
//			},
//			StopAllFunc: func() error {
//				panic("mock out the StopAll method")
//			},
//		}
//
//		// use mockedSensorController in code that requires SensorController
//		// and then make assertions.
//
//	}
type SensorControllerMock struct {
	// ActiveWorkersFunc mocks the ActiveWorkers method.
	ActiveWorkersFunc func() []string

	// InactiveWorkersFunc mocks the InactiveWorkers method.
	InactiveWorkersFunc func() []string

	// StartFunc mocks the Start method.
	StartFunc func(ctx context.Context, name string) (<-chan sensor.Details, error)

	// StartAllFunc mocks the StartAll method.
	StartAllFunc func(ctx context.Context) (<-chan sensor.Details, error)

	// StopFunc mocks the Stop method.
	StopFunc func(name string) error

	// StopAllFunc mocks the StopAll method.
	StopAllFunc func() error

	// calls tracks calls to the methods.
	calls struct {
		// ActiveWorkers holds details about calls to the ActiveWorkers method.
		ActiveWorkers []struct {
		}
		// InactiveWorkers holds details about calls to the InactiveWorkers method.
		InactiveWorkers []struct {
		}
		// Start holds details about calls to the Start method.
		Start []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Name is the name argument value.
			Name string
		}
		// StartAll holds details about calls to the StartAll method.
		StartAll []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
		}
		// Stop holds details about calls to the Stop method.
		Stop []struct {
			// Name is the name argument value.
			Name string
		}
		// StopAll holds details about calls to the StopAll method.
		StopAll []struct {
		}
	}
	lockActiveWorkers   sync.RWMutex
	lockInactiveWorkers sync.RWMutex
	lockStart           sync.RWMutex
	lockStartAll        sync.RWMutex
	lockStop            sync.RWMutex
	lockStopAll         sync.RWMutex
}

// ActiveWorkers calls ActiveWorkersFunc.
func (mock *SensorControllerMock) ActiveWorkers() []string {
	if mock.ActiveWorkersFunc == nil {
		panic("SensorControllerMock.ActiveWorkersFunc: method is nil but SensorController.ActiveWorkers was just called")
	}
	callInfo := struct {
	}{}
	mock.lockActiveWorkers.Lock()
	mock.calls.ActiveWorkers = append(mock.calls.ActiveWorkers, callInfo)
	mock.lockActiveWorkers.Unlock()
	return mock.ActiveWorkersFunc()
}

// ActiveWorkersCalls gets all the calls that were made to ActiveWorkers.
// Check the length with:
//
//	len(mockedSensorController.ActiveWorkersCalls())
func (mock *SensorControllerMock) ActiveWorkersCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockActiveWorkers.RLock()
	calls = mock.calls.ActiveWorkers
	mock.lockActiveWorkers.RUnlock()
	return calls
}

// InactiveWorkers calls InactiveWorkersFunc.
func (mock *SensorControllerMock) InactiveWorkers() []string {
	if mock.InactiveWorkersFunc == nil {
		panic("SensorControllerMock.InactiveWorkersFunc: method is nil but SensorController.InactiveWorkers was just called")
	}
	callInfo := struct {
	}{}
	mock.lockInactiveWorkers.Lock()
	mock.calls.InactiveWorkers = append(mock.calls.InactiveWorkers, callInfo)
	mock.lockInactiveWorkers.Unlock()
	return mock.InactiveWorkersFunc()
}

// InactiveWorkersCalls gets all the calls that were made to InactiveWorkers.
// Check the length with:
//
//	len(mockedSensorController.InactiveWorkersCalls())
func (mock *SensorControllerMock) InactiveWorkersCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockInactiveWorkers.RLock()
	calls = mock.calls.InactiveWorkers
	mock.lockInactiveWorkers.RUnlock()
	return calls
}

// Start calls StartFunc.
func (mock *SensorControllerMock) Start(ctx context.Context, name string) (<-chan sensor.Details, error) {
	if mock.StartFunc == nil {
		panic("SensorControllerMock.StartFunc: method is nil but SensorController.Start was just called")
	}
	callInfo := struct {
		Ctx  context.Context
		Name string
	}{
		Ctx:  ctx,
		Name: name,
	}
	mock.lockStart.Lock()
	mock.calls.Start = append(mock.calls.Start, callInfo)
	mock.lockStart.Unlock()
	return mock.StartFunc(ctx, name)
}

// StartCalls gets all the calls that were made to Start.
// Check the length with:
//
//	len(mockedSensorController.StartCalls())
func (mock *SensorControllerMock) StartCalls() []struct {
	Ctx  context.Context
	Name string
} {
	var calls []struct {
		Ctx  context.Context
		Name string
	}
	mock.lockStart.RLock()
	calls = mock.calls.Start
	mock.lockStart.RUnlock()
	return calls
}

// StartAll calls StartAllFunc.
func (mock *SensorControllerMock) StartAll(ctx context.Context) (<-chan sensor.Details, error) {
	if mock.StartAllFunc == nil {
		panic("SensorControllerMock.StartAllFunc: method is nil but SensorController.StartAll was just called")
	}
	callInfo := struct {
		Ctx context.Context
	}{
		Ctx: ctx,
	}
	mock.lockStartAll.Lock()
	mock.calls.StartAll = append(mock.calls.StartAll, callInfo)
	mock.lockStartAll.Unlock()
	return mock.StartAllFunc(ctx)
}

// StartAllCalls gets all the calls that were made to StartAll.
// Check the length with:
//
//	len(mockedSensorController.StartAllCalls())
func (mock *SensorControllerMock) StartAllCalls() []struct {
	Ctx context.Context
} {
	var calls []struct {
		Ctx context.Context
	}
	mock.lockStartAll.RLock()
	calls = mock.calls.StartAll
	mock.lockStartAll.RUnlock()
	return calls
}

// Stop calls StopFunc.
func (mock *SensorControllerMock) Stop(name string) error {
	if mock.StopFunc == nil {
		panic("SensorControllerMock.StopFunc: method is nil but SensorController.Stop was just called")
	}
	callInfo := struct {
		Name string
	}{
		Name: name,
	}
	mock.lockStop.Lock()
	mock.calls.Stop = append(mock.calls.Stop, callInfo)
	mock.lockStop.Unlock()
	return mock.StopFunc(name)
}

// StopCalls gets all the calls that were made to Stop.
// Check the length with:
//
//	len(mockedSensorController.StopCalls())
func (mock *SensorControllerMock) StopCalls() []struct {
	Name string
} {
	var calls []struct {
		Name string
	}
	mock.lockStop.RLock()
	calls = mock.calls.Stop
	mock.lockStop.RUnlock()
	return calls
}

// StopAll calls StopAllFunc.
func (mock *SensorControllerMock) StopAll() error {
	if mock.StopAllFunc == nil {
		panic("SensorControllerMock.StopAllFunc: method is nil but SensorController.StopAll was just called")
	}
	callInfo := struct {
	}{}
	mock.lockStopAll.Lock()
	mock.calls.StopAll = append(mock.calls.StopAll, callInfo)
	mock.lockStopAll.Unlock()
	return mock.StopAllFunc()
}

// StopAllCalls gets all the calls that were made to StopAll.
// Check the length with:
//
//	len(mockedSensorController.StopAllCalls())
func (mock *SensorControllerMock) StopAllCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockStopAll.RLock()
	calls = mock.calls.StopAll
	mock.lockStopAll.RUnlock()
	return calls
}

// Ensure, that WorkerMock does implement Worker.
// If this is not the case, regenerate this file with moq.
var _ Worker = &WorkerMock{}

// WorkerMock is a mock implementation of Worker.
//
//	func TestSomethingThatUsesWorker(t *testing.T) {
//
//		// make and configure a mocked Worker
//		mockedWorker := &WorkerMock{
//			IDFunc: func() string {
//				panic("mock out the ID method")
//			},
//			SensorsFunc: func(ctx context.Context) ([]sensor.Details, error) {
//				panic("mock out the Sensors method")
//			},
//			StopFunc: func() error {
//				panic("mock out the Stop method")
//			},
//			UpdatesFunc: func(ctx context.Context) (<-chan sensor.Details, error) {
//				panic("mock out the Updates method")
//			},
//		}
//
//		// use mockedWorker in code that requires Worker
//		// and then make assertions.
//
//	}
type WorkerMock struct {
	// IDFunc mocks the ID method.
	IDFunc func() string

	// SensorsFunc mocks the Sensors method.
	SensorsFunc func(ctx context.Context) ([]sensor.Details, error)

	// StopFunc mocks the Stop method.
	StopFunc func() error

	// UpdatesFunc mocks the Updates method.
	UpdatesFunc func(ctx context.Context) (<-chan sensor.Details, error)

	// calls tracks calls to the methods.
	calls struct {
		// ID holds details about calls to the ID method.
		ID []struct {
		}
		// Sensors holds details about calls to the Sensors method.
		Sensors []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
		}
		// Stop holds details about calls to the Stop method.
		Stop []struct {
		}
		// Updates holds details about calls to the Updates method.
		Updates []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
		}
	}
	lockID      sync.RWMutex
	lockSensors sync.RWMutex
	lockStop    sync.RWMutex
	lockUpdates sync.RWMutex
}

// ID calls IDFunc.
func (mock *WorkerMock) ID() string {
	if mock.IDFunc == nil {
		panic("WorkerMock.IDFunc: method is nil but Worker.ID was just called")
	}
	callInfo := struct {
	}{}
	mock.lockID.Lock()
	mock.calls.ID = append(mock.calls.ID, callInfo)
	mock.lockID.Unlock()
	return mock.IDFunc()
}

// IDCalls gets all the calls that were made to ID.
// Check the length with:
//
//	len(mockedWorker.IDCalls())
func (mock *WorkerMock) IDCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockID.RLock()
	calls = mock.calls.ID
	mock.lockID.RUnlock()
	return calls
}

// Sensors calls SensorsFunc.
func (mock *WorkerMock) Sensors(ctx context.Context) ([]sensor.Details, error) {
	if mock.SensorsFunc == nil {
		panic("WorkerMock.SensorsFunc: method is nil but Worker.Sensors was just called")
	}
	callInfo := struct {
		Ctx context.Context
	}{
		Ctx: ctx,
	}
	mock.lockSensors.Lock()
	mock.calls.Sensors = append(mock.calls.Sensors, callInfo)
	mock.lockSensors.Unlock()
	return mock.SensorsFunc(ctx)
}

// SensorsCalls gets all the calls that were made to Sensors.
// Check the length with:
//
//	len(mockedWorker.SensorsCalls())
func (mock *WorkerMock) SensorsCalls() []struct {
	Ctx context.Context
} {
	var calls []struct {
		Ctx context.Context
	}
	mock.lockSensors.RLock()
	calls = mock.calls.Sensors
	mock.lockSensors.RUnlock()
	return calls
}

// Stop calls StopFunc.
func (mock *WorkerMock) Stop() error {
	if mock.StopFunc == nil {
		panic("WorkerMock.StopFunc: method is nil but Worker.Stop was just called")
	}
	callInfo := struct {
	}{}
	mock.lockStop.Lock()
	mock.calls.Stop = append(mock.calls.Stop, callInfo)
	mock.lockStop.Unlock()
	return mock.StopFunc()
}

// StopCalls gets all the calls that were made to Stop.
// Check the length with:
//
//	len(mockedWorker.StopCalls())
func (mock *WorkerMock) StopCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockStop.RLock()
	calls = mock.calls.Stop
	mock.lockStop.RUnlock()
	return calls
}

// Updates calls UpdatesFunc.
func (mock *WorkerMock) Updates(ctx context.Context) (<-chan sensor.Details, error) {
	if mock.UpdatesFunc == nil {
		panic("WorkerMock.UpdatesFunc: method is nil but Worker.Updates was just called")
	}
	callInfo := struct {
		Ctx context.Context
	}{
		Ctx: ctx,
	}
	mock.lockUpdates.Lock()
	mock.calls.Updates = append(mock.calls.Updates, callInfo)
	mock.lockUpdates.Unlock()
	return mock.UpdatesFunc(ctx)
}

// UpdatesCalls gets all the calls that were made to Updates.
// Check the length with:
//
//	len(mockedWorker.UpdatesCalls())
func (mock *WorkerMock) UpdatesCalls() []struct {
	Ctx context.Context
} {
	var calls []struct {
		Ctx context.Context
	}
	mock.lockUpdates.RLock()
	calls = mock.calls.Updates
	mock.lockUpdates.RUnlock()
	return calls
}

// Ensure, that ScriptMock does implement Script.
// If this is not the case, regenerate this file with moq.
var _ Script = &ScriptMock{}

// ScriptMock is a mock implementation of Script.
//
//	func TestSomethingThatUsesScript(t *testing.T) {
//
//		// make and configure a mocked Script
//		mockedScript := &ScriptMock{
//			ExecuteFunc: func() ([]sensor.Details, error) {
//				panic("mock out the Execute method")
//			},
//			ScheduleFunc: func() string {
//				panic("mock out the Schedule method")
//			},
//		}
//
//		// use mockedScript in code that requires Script
//		// and then make assertions.
//
//	}
type ScriptMock struct {
	// ExecuteFunc mocks the Execute method.
	ExecuteFunc func() ([]sensor.Details, error)

	// ScheduleFunc mocks the Schedule method.
	ScheduleFunc func() string

	// calls tracks calls to the methods.
	calls struct {
		// Execute holds details about calls to the Execute method.
		Execute []struct {
		}
		// Schedule holds details about calls to the Schedule method.
		Schedule []struct {
		}
	}
	lockExecute  sync.RWMutex
	lockSchedule sync.RWMutex
}

// Execute calls ExecuteFunc.
func (mock *ScriptMock) Execute() ([]sensor.Details, error) {
	if mock.ExecuteFunc == nil {
		panic("ScriptMock.ExecuteFunc: method is nil but Script.Execute was just called")
	}
	callInfo := struct {
	}{}
	mock.lockExecute.Lock()
	mock.calls.Execute = append(mock.calls.Execute, callInfo)
	mock.lockExecute.Unlock()
	return mock.ExecuteFunc()
}

// ExecuteCalls gets all the calls that were made to Execute.
// Check the length with:
//
//	len(mockedScript.ExecuteCalls())
func (mock *ScriptMock) ExecuteCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockExecute.RLock()
	calls = mock.calls.Execute
	mock.lockExecute.RUnlock()
	return calls
}

// Schedule calls ScheduleFunc.
func (mock *ScriptMock) Schedule() string {
	if mock.ScheduleFunc == nil {
		panic("ScriptMock.ScheduleFunc: method is nil but Script.Schedule was just called")
	}
	callInfo := struct {
	}{}
	mock.lockSchedule.Lock()
	mock.calls.Schedule = append(mock.calls.Schedule, callInfo)
	mock.lockSchedule.Unlock()
	return mock.ScheduleFunc()
}

// ScheduleCalls gets all the calls that were made to Schedule.
// Check the length with:
//
//	len(mockedScript.ScheduleCalls())
func (mock *ScriptMock) ScheduleCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockSchedule.RLock()
	calls = mock.calls.Schedule
	mock.lockSchedule.RUnlock()
	return calls
}