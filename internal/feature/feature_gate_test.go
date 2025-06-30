package feature

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/earaujoassis/space/internal/config"
	"github.com/earaujoassis/space/internal/gateways/redis"
)

type FeatureGateTestSuite struct {
	suite.Suite
	fg *FeatureGate
}

func (s *FeatureGateTestSuite) SetupSuite() {
	cfg := &config.Config{
		Environment: "test",
	}
	ms, _ := redis.NewMemoryService(cfg)
	s.fg = NewFeatureGate(ms)
}

func (s *FeatureGateTestSuite) TestIsActive() {
	s.False(s.fg.IsActive("no-feature"), "shouldn't have no-feature active")
}

func (s *FeatureGateTestSuite) TestEnable() {
	s.False(s.fg.IsActive("not-enabled"), "shouldn't have not-enabled active")
	s.fg.Enable("not-enabled")
	s.True(s.fg.IsActive("not-enabled"), "should have no-feature active")
	s.fg.Disable("not-enabled")
}

func (s *FeatureGateTestSuite) TestDisable() {
	s.False(s.fg.IsActive("to-disable"), "shouldn't have to-disable active")
	s.fg.Enable("to-disable")
	s.True(s.fg.IsActive("to-disable"), "should have to-disable active")
	s.fg.Disable("to-disable")
	s.False(s.fg.IsActive("to-disable"), "shouldn't have to-disable active")
}

func TestFeatureGateSuite(t *testing.T) {
	suite.Run(t, new(FeatureGateTestSuite))
}
