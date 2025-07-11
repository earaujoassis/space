package policy

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/earaujoassis/space/internal/config"
	"github.com/earaujoassis/space/internal/gateways/memory"
	"github.com/earaujoassis/space/test/utils"
)

type PolicyTestSuite struct {
	suite.Suite
	ms  *memory.MemoryService
	rls *RateLimitService
}

func (s *PolicyTestSuite) SetupSuite() {
	utils.SetupConfigEnv()
	config, _ := config.Load()
	ms, err := memory.NewMemoryService(config)
	s.Require().NoError(err)
	if err != nil {
		s.T().Fatalf("Could not create new memory service: %v", err)
	}
	s.rls = NewRateLimitService(ms)
	s.ms = ms
}

func (s *PolicyTestSuite) SetupTest() {
	s.cleanupRedis()
}

func (s *PolicyTestSuite) cleanupRedis() {
	s.ms.Do("FLUSHDB")
}

func (s *PolicyTestSuite) TearDownSuite() {
	s.ms.Close()
}

func TestPolicySuite(t *testing.T) {
	suite.Run(t, new(PolicyTestSuite))
}
