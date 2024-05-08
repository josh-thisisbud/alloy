// This service provides debug stream APIs for components.
package xray

import (
	"context"
	"fmt"

	"github.com/grafana/alloy/internal/featuregate"
	"github.com/grafana/alloy/internal/service"
)

// ServiceName defines the name used for the xray service.
const ServiceName = "xray"

type Service struct {
	debugStreamManager DebugStreamHandler
}

var _ service.Service = (*Service)(nil)

func New() *Service {
	return &Service{
		debugStreamManager: NewDebugStreamManager(),
	}
}

// Data implements service.Service.
// It returns the debugStreamManager for the components to stream.
func (s *Service) Data() any {
	return s.debugStreamManager
}

// Definition implements service.Service.
func (*Service) Definition() service.Definition {
	return service.Definition{
		Name:       ServiceName,
		ConfigType: nil, // xray does not accept configuration
		DependsOn:  []string{},
		Stability:  featuregate.StabilityExperimental,
	}
}

// Run implements service.Service.
func (s *Service) Run(ctx context.Context, _ service.Host) error {
	<-ctx.Done()
	return nil
}

// Update implements service.Service.
func (*Service) Update(_ any) error {
	return fmt.Errorf("xray service does not support configuration")
}
