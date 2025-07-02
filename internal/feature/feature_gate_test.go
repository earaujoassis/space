package feature

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/earaujoassis/space/internal/config"
	"github.com/earaujoassis/space/internal/gateways/memory"
)

type FeatureGateTestSuite struct {
	suite.Suite
	fg *FeatureGate
}

func (s *FeatureGateTestSuite) SetupSuite() {
	cfg := &config.Config{
		Environment: "test",
	}
	ms, _ := memory.NewMemoryService(cfg)
	s.fg = NewFeatureGate(ms)
}

func (s *FeatureGateTestSuite) TestIsActive() {
	s.False(s.fg.IsActive("no-feature"))
}

func (s *FeatureGateTestSuite) TestEnable() {
	s.False(s.fg.IsActive("not-enabled"))
	s.fg.Enable("not-enabled")
	s.True(s.fg.IsActive("not-enabled"))
	s.fg.Disable("not-enabled")
}

func (s *FeatureGateTestSuite) TestDisable() {
	s.False(s.fg.IsActive("to-disable"))
	s.fg.Enable("to-disable")
	s.True(s.fg.IsActive("to-disable"))
	s.fg.Disable("to-disable")
	s.False(s.fg.IsActive("to-disable"))
}

func TestFeatureGateSuite(t *testing.T) {
	suite.Run(t, new(FeatureGateTestSuite))
}
