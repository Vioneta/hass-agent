// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package ui

import (
	"github.com/joshuar/go-hass-agent/internal/tracker"
	"sync"
)

// Ensure, that AgentMock does implement Agent.
// If this is not the case, regenerate this file with moq.
var _ Agent = &AgentMock{}

// AgentMock is a mock implementation of Agent.
//
//	func TestSomethingThatUsesAgent(t *testing.T) {
//
//		// make and configure a mocked Agent
//		mockedAgent := &AgentMock{
//			AppIDFunc: func() string {
//				panic("mock out the AppID method")
//			},
//			AppNameFunc: func() string {
//				panic("mock out the AppName method")
//			},
//			AppVersionFunc: func() string {
//				panic("mock out the AppVersion method")
//			},
//			GetConfigFunc: func(s string, ifaceVal interface{}) error {
//				panic("mock out the GetConfig method")
//			},
//			IsHeadlessFunc: func() bool {
//				panic("mock out the IsHeadless method")
//			},
//			SensorListFunc: func() []string {
//				panic("mock out the SensorList method")
//			},
//			SensorValueFunc: func(s string) (tracker.Sensor, error) {
//				panic("mock out the SensorValue method")
//			},
//			SetConfigFunc: func(s string, ifaceVal interface{}) error {
//				panic("mock out the SetConfig method")
//			},
//			StopFunc: func()  {
//				panic("mock out the Stop method")
//			},
//		}
//
//		// use mockedAgent in code that requires Agent
//		// and then make assertions.
//
//	}
type AgentMock struct {
	// AppIDFunc mocks the AppID method.
	AppIDFunc func() string

	// AppNameFunc mocks the AppName method.
	AppNameFunc func() string

	// AppVersionFunc mocks the AppVersion method.
	AppVersionFunc func() string

	// GetConfigFunc mocks the GetConfig method.
	GetConfigFunc func(s string, ifaceVal interface{}) error

	// IsHeadlessFunc mocks the IsHeadless method.
	IsHeadlessFunc func() bool

	// SensorListFunc mocks the SensorList method.
	SensorListFunc func() []string

	// SensorValueFunc mocks the SensorValue method.
	SensorValueFunc func(s string) (tracker.Sensor, error)

	// SetConfigFunc mocks the SetConfig method.
	SetConfigFunc func(s string, ifaceVal interface{}) error

	// StopFunc mocks the Stop method.
	StopFunc func()

	// calls tracks calls to the methods.
	calls struct {
		// AppID holds details about calls to the AppID method.
		AppID []struct {
		}
		// AppName holds details about calls to the AppName method.
		AppName []struct {
		}
		// AppVersion holds details about calls to the AppVersion method.
		AppVersion []struct {
		}
		// GetConfig holds details about calls to the GetConfig method.
		GetConfig []struct {
			// S is the s argument value.
			S string
			// IfaceVal is the ifaceVal argument value.
			IfaceVal interface{}
		}
		// IsHeadless holds details about calls to the IsHeadless method.
		IsHeadless []struct {
		}
		// SensorList holds details about calls to the SensorList method.
		SensorList []struct {
		}
		// SensorValue holds details about calls to the SensorValue method.
		SensorValue []struct {
			// S is the s argument value.
			S string
		}
		// SetConfig holds details about calls to the SetConfig method.
		SetConfig []struct {
			// S is the s argument value.
			S string
			// IfaceVal is the ifaceVal argument value.
			IfaceVal interface{}
		}
		// Stop holds details about calls to the Stop method.
		Stop []struct {
		}
	}
	lockAppID       sync.RWMutex
	lockAppName     sync.RWMutex
	lockAppVersion  sync.RWMutex
	lockGetConfig   sync.RWMutex
	lockIsHeadless  sync.RWMutex
	lockSensorList  sync.RWMutex
	lockSensorValue sync.RWMutex
	lockSetConfig   sync.RWMutex
	lockStop        sync.RWMutex
}

// AppID calls AppIDFunc.
func (mock *AgentMock) AppID() string {
	if mock.AppIDFunc == nil {
		panic("AgentMock.AppIDFunc: method is nil but Agent.AppID was just called")
	}
	callInfo := struct {
	}{}
	mock.lockAppID.Lock()
	mock.calls.AppID = append(mock.calls.AppID, callInfo)
	mock.lockAppID.Unlock()
	return mock.AppIDFunc()
}

// AppIDCalls gets all the calls that were made to AppID.
// Check the length with:
//
//	len(mockedAgent.AppIDCalls())
func (mock *AgentMock) AppIDCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockAppID.RLock()
	calls = mock.calls.AppID
	mock.lockAppID.RUnlock()
	return calls
}

// AppName calls AppNameFunc.
func (mock *AgentMock) AppName() string {
	if mock.AppNameFunc == nil {
		panic("AgentMock.AppNameFunc: method is nil but Agent.AppName was just called")
	}
	callInfo := struct {
	}{}
	mock.lockAppName.Lock()
	mock.calls.AppName = append(mock.calls.AppName, callInfo)
	mock.lockAppName.Unlock()
	return mock.AppNameFunc()
}

// AppNameCalls gets all the calls that were made to AppName.
// Check the length with:
//
//	len(mockedAgent.AppNameCalls())
func (mock *AgentMock) AppNameCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockAppName.RLock()
	calls = mock.calls.AppName
	mock.lockAppName.RUnlock()
	return calls
}

// AppVersion calls AppVersionFunc.
func (mock *AgentMock) AppVersion() string {
	if mock.AppVersionFunc == nil {
		panic("AgentMock.AppVersionFunc: method is nil but Agent.AppVersion was just called")
	}
	callInfo := struct {
	}{}
	mock.lockAppVersion.Lock()
	mock.calls.AppVersion = append(mock.calls.AppVersion, callInfo)
	mock.lockAppVersion.Unlock()
	return mock.AppVersionFunc()
}

// AppVersionCalls gets all the calls that were made to AppVersion.
// Check the length with:
//
//	len(mockedAgent.AppVersionCalls())
func (mock *AgentMock) AppVersionCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockAppVersion.RLock()
	calls = mock.calls.AppVersion
	mock.lockAppVersion.RUnlock()
	return calls
}

// GetConfig calls GetConfigFunc.
func (mock *AgentMock) GetConfig(s string, ifaceVal interface{}) error {
	if mock.GetConfigFunc == nil {
		panic("AgentMock.GetConfigFunc: method is nil but Agent.GetConfig was just called")
	}
	callInfo := struct {
		S        string
		IfaceVal interface{}
	}{
		S:        s,
		IfaceVal: ifaceVal,
	}
	mock.lockGetConfig.Lock()
	mock.calls.GetConfig = append(mock.calls.GetConfig, callInfo)
	mock.lockGetConfig.Unlock()
	return mock.GetConfigFunc(s, ifaceVal)
}

// GetConfigCalls gets all the calls that were made to GetConfig.
// Check the length with:
//
//	len(mockedAgent.GetConfigCalls())
func (mock *AgentMock) GetConfigCalls() []struct {
	S        string
	IfaceVal interface{}
} {
	var calls []struct {
		S        string
		IfaceVal interface{}
	}
	mock.lockGetConfig.RLock()
	calls = mock.calls.GetConfig
	mock.lockGetConfig.RUnlock()
	return calls
}

// IsHeadless calls IsHeadlessFunc.
func (mock *AgentMock) IsHeadless() bool {
	if mock.IsHeadlessFunc == nil {
		panic("AgentMock.IsHeadlessFunc: method is nil but Agent.IsHeadless was just called")
	}
	callInfo := struct {
	}{}
	mock.lockIsHeadless.Lock()
	mock.calls.IsHeadless = append(mock.calls.IsHeadless, callInfo)
	mock.lockIsHeadless.Unlock()
	return mock.IsHeadlessFunc()
}

// IsHeadlessCalls gets all the calls that were made to IsHeadless.
// Check the length with:
//
//	len(mockedAgent.IsHeadlessCalls())
func (mock *AgentMock) IsHeadlessCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockIsHeadless.RLock()
	calls = mock.calls.IsHeadless
	mock.lockIsHeadless.RUnlock()
	return calls
}

// SensorList calls SensorListFunc.
func (mock *AgentMock) SensorList() []string {
	if mock.SensorListFunc == nil {
		panic("AgentMock.SensorListFunc: method is nil but Agent.SensorList was just called")
	}
	callInfo := struct {
	}{}
	mock.lockSensorList.Lock()
	mock.calls.SensorList = append(mock.calls.SensorList, callInfo)
	mock.lockSensorList.Unlock()
	return mock.SensorListFunc()
}

// SensorListCalls gets all the calls that were made to SensorList.
// Check the length with:
//
//	len(mockedAgent.SensorListCalls())
func (mock *AgentMock) SensorListCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockSensorList.RLock()
	calls = mock.calls.SensorList
	mock.lockSensorList.RUnlock()
	return calls
}

// SensorValue calls SensorValueFunc.
func (mock *AgentMock) SensorValue(s string) (tracker.Sensor, error) {
	if mock.SensorValueFunc == nil {
		panic("AgentMock.SensorValueFunc: method is nil but Agent.SensorValue was just called")
	}
	callInfo := struct {
		S string
	}{
		S: s,
	}
	mock.lockSensorValue.Lock()
	mock.calls.SensorValue = append(mock.calls.SensorValue, callInfo)
	mock.lockSensorValue.Unlock()
	return mock.SensorValueFunc(s)
}

// SensorValueCalls gets all the calls that were made to SensorValue.
// Check the length with:
//
//	len(mockedAgent.SensorValueCalls())
func (mock *AgentMock) SensorValueCalls() []struct {
	S string
} {
	var calls []struct {
		S string
	}
	mock.lockSensorValue.RLock()
	calls = mock.calls.SensorValue
	mock.lockSensorValue.RUnlock()
	return calls
}

// SetConfig calls SetConfigFunc.
func (mock *AgentMock) SetConfig(s string, ifaceVal interface{}) error {
	if mock.SetConfigFunc == nil {
		panic("AgentMock.SetConfigFunc: method is nil but Agent.SetConfig was just called")
	}
	callInfo := struct {
		S        string
		IfaceVal interface{}
	}{
		S:        s,
		IfaceVal: ifaceVal,
	}
	mock.lockSetConfig.Lock()
	mock.calls.SetConfig = append(mock.calls.SetConfig, callInfo)
	mock.lockSetConfig.Unlock()
	return mock.SetConfigFunc(s, ifaceVal)
}

// SetConfigCalls gets all the calls that were made to SetConfig.
// Check the length with:
//
//	len(mockedAgent.SetConfigCalls())
func (mock *AgentMock) SetConfigCalls() []struct {
	S        string
	IfaceVal interface{}
} {
	var calls []struct {
		S        string
		IfaceVal interface{}
	}
	mock.lockSetConfig.RLock()
	calls = mock.calls.SetConfig
	mock.lockSetConfig.RUnlock()
	return calls
}

// Stop calls StopFunc.
func (mock *AgentMock) Stop() {
	if mock.StopFunc == nil {
		panic("AgentMock.StopFunc: method is nil but Agent.Stop was just called")
	}
	callInfo := struct {
	}{}
	mock.lockStop.Lock()
	mock.calls.Stop = append(mock.calls.Stop, callInfo)
	mock.lockStop.Unlock()
	mock.StopFunc()
}

// StopCalls gets all the calls that were made to Stop.
// Check the length with:
//
//	len(mockedAgent.StopCalls())
func (mock *AgentMock) StopCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockStop.RLock()
	calls = mock.calls.Stop
	mock.lockStop.RUnlock()
	return calls
}