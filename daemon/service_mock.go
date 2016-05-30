package daemon

import "os/exec"

// ServiceMock is the mock implementation of the Service interface.
type ServiceMock struct {
	Command             string
	processes           map[int]*exec.Cmd
	Args                []string
	ServiceStopResult   bool
	ServiceStopError    error
	ServiceList         map[int]*exec.Cmd
	ServiceStartCmd     *exec.Cmd
	ServicePort         int
	ServiceStopCount    int
	ServicesSetupCalled bool
}

// Setup the Management services.
func (s *ServiceMock) Setup() {
	s.ServicesSetupCalled = true
}

// Stop a Service and returns the exit status.
func (s *ServiceMock) Stop(pid int) (bool, error) {
	s.ServiceStopCount++
	return s.ServiceStopResult, s.ServiceStopError
}

// List all Service PIDs.
func (s *ServiceMock) List() map[int]*exec.Cmd {
	return s.ServiceList
}

// Start a Service and log its output.
func (s *ServiceMock) Start() *exec.Cmd {
	return s.ServiceStartCmd
}

// NewService creates a new MockService with default settings.
func (s *ServiceMock) NewService() (int, Service) {
	return s.ServicePort, s
}