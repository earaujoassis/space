package policy

import (
	"github.com/earaujoassis/space/internal/gateways/redis"
)

type RateLimitService struct {
	ms *redis.MemoryService
}

func NewRateLimitService(ms *redis.MemoryService) *RateLimitService {
	return &RateLimitService{ms: ms}
}

// SignInAttemptStatus checks and controls sign-in attempts from a Web browser/User
func (rls *RateLimitService) SignInAttemptStatus(id string) string {
	var result string

	rls.ms.Transaction(func(c *redis.Commands) {
		if c.CheckFieldExistence("sign-in.blocked", id) {
			result = Blocked
		} else if c.CheckFieldExistence("sign-in.attempt", id) {
			numberOfAttempts := c.GetFieldAtKey("sign-in.attempt", id).ToInt()
			switch {
			case numberOfAttempts > 0 && numberOfAttempts <= attemptsUntilPreblock:
				result = Clear
			case numberOfAttempts > attemptsUntilPreblock && numberOfAttempts <= attemptsUntilBlock:
				result = Preblocked
			case numberOfAttempts > attemptsUntilBlock:
				result = Blocked
			}
		} else {
			result = Clear
		}
	})

	return result
}

// SignUpAttemptStatus checks and controls sign-up attempts from a Web browser/User
func (rls *RateLimitService) SignUpAttemptStatus(id string) string {
	var result string

	rls.ms.Transaction(func(c *redis.Commands) {
		if c.CheckFieldExistence("sign-up.blocked", id) {
			result = Blocked
		} else if c.CheckFieldExistence("sign-up.attempt", id) {
			numberOfAttempts := c.GetFieldAtKey("sign-up.attempt", id).ToInt()
			switch {
			case numberOfAttempts > 0 && numberOfAttempts <= attemptsUntilPreblock:
				result = Clear
			case numberOfAttempts > attemptsUntilPreblock && numberOfAttempts <= attemptsUntilBlock:
				result = Preblocked
			case numberOfAttempts > attemptsUntilBlock:
				result = Blocked
			}
		} else {
			result = Clear
		}
	})

	return result
}
