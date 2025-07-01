package policy

import (
	"time"

	"github.com/earaujoassis/space/internal/gateways/memory"
)

// RegisterSignInAttempt records a sign-in attempt (in order to control it)
func (rls *RateLimitService) RegisterSignInAttempt(id string) {
	rls.ms.Transaction(func(c *memory.Commands) {
		nowMoment := time.Now().UTC().Unix()

		if c.CheckFieldExistence("sign-in.blocked", id) {
			blockMoment := c.GetFieldAtKey("sign-in.blocked", id).ToInt64()
			if (nowMoment - blockMoment) >= blockPeriodFailedSignIn {
				c.DeleteFieldAtKey("sign-in.blocked", id)
				c.SetFieldAtKey("sign-in.attempt", id, 1)
			}
			return
		}
		if !c.CheckFieldExistence("sign-in.attempt", id) {
			c.SetFieldAtKey("sign-in.attempt", id, 1)
		} else {
			c.IncrementFieldAtKeyBy("sign-in.attempt", id, 1)
			numberOfAttempts := c.GetFieldAtKey("sign-in.attempt", id).ToInt()
			if numberOfAttempts >= attemptsUntilBlock {
				c.SetFieldAtKey("sign-in.blocked", id, nowMoment)
			}
		}
	})
}

// RegisterSuccessfulSignIn records a successful sign-in attempt (in order to clear it)
func (rls *RateLimitService) RegisterSuccessfulSignIn(id string) {
	rls.ms.Transaction(func(c *memory.Commands) {
		c.DeleteFieldAtKey("sign-in.attempt", id)
		c.DeleteFieldAtKey("sign-in.blocked", id)
	})
}

// RegisterSignUpAttempt records a sign-up attempt (in order to control it)
func (rls *RateLimitService) RegisterSignUpAttempt(id string) {
	rls.ms.Transaction(func(c *memory.Commands) {
		nowMoment := time.Now().UTC().Unix()
		if c.CheckFieldExistence("sign-up.blocked", id) {
			blockMoment := c.GetFieldAtKey("sign-up.blocked", id).ToInt64()
			if (nowMoment - blockMoment) >= blockPeriodFailedSignUp {
				c.DeleteFieldAtKey("sign-up.blocked", id)
				c.SetFieldAtKey("sign-up.attempt", id, 1)
			}
			return
		}
		if !c.CheckFieldExistence("sign-up.attempt", id) {
			c.SetFieldAtKey("sign-up.attempt", id, 1)
		} else {
			c.IncrementFieldAtKeyBy("sign-up.attempt", id, 1)
			numberOfAttempts := c.GetFieldAtKey("sign-up.attempt", id).ToInt()
			if numberOfAttempts >= attemptsUntilBlock {
				c.SetFieldAtKey("sign-up.blocked", id, nowMoment)
			}
		}
	})
}

// RegisterSuccessfulSignUp records a successful sign-up attempt (in order to clear it)
func (rls *RateLimitService) RegisterSuccessfulSignUp(id string) {
	rls.ms.Transaction(func(c *memory.Commands) {
		c.DeleteFieldAtKey("sign-up.attempt", id)
		c.DeleteFieldAtKey("sign-up.blocked", id)
	})
}
