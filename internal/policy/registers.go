package policy

import (
	"time"

	"github.com/earaujoassis/space/internal/services/volatile"
)

// RegisterSignInAttempt records a sign-in attempt (in order to control it)
func RegisterSignInAttempt(id string) {
	volatile.TransactionWrapper(func() {
		nowMoment := time.Now().UTC().Unix()
		if volatile.CheckFieldExistence("sign-in.blocked", id) {
			blockMoment := volatile.GetFieldAtKey("sign-in.blocked", id).ToInt64()
			if (nowMoment - blockMoment) >= blockPeriodFailedSignIn {
				volatile.DeleteFieldAtKey("sign-in.blocked", id)
				volatile.SetFieldAtKey("sign-in.attempt", id, 1)
			}
			return
		}
		if !volatile.CheckFieldExistence("sign-in.attempt", id) {
			volatile.SetFieldAtKey("sign-in.attempt", id, 1)
		} else {
			volatile.IncrementFieldAtKeyBy("sign-in.attempt", id, 1)
			numberOfAttempts := volatile.GetFieldAtKey("sign-in.attempt", id).ToInt()
			if numberOfAttempts >= attemptsUntilBlock {
				volatile.SetFieldAtKey("sign-in.blocked", id, nowMoment)
			}
		}
	})
}

// RegisterSuccessfulSignIn records a successful sign-in attempt (in order to clear it)
func RegisterSuccessfulSignIn(id string) {
	volatile.TransactionWrapper(func() {
		volatile.DeleteFieldAtKey("sign-in.attempt", id)
		volatile.DeleteFieldAtKey("sign-in.blocked", id)
	})
}

// RegisterSignUpAttempt records a sign-up attempt (in order to control it)
func RegisterSignUpAttempt(id string) {
	volatile.TransactionWrapper(func() {
		nowMoment := time.Now().UTC().Unix()
		if volatile.CheckFieldExistence("sign-up.blocked", id) {
			blockMoment := volatile.GetFieldAtKey("sign-up.blocked", id).ToInt64()
			if (nowMoment - blockMoment) >= blockPeriodFailedSignUp {
				volatile.DeleteFieldAtKey("sign-up.blocked", id)
				volatile.SetFieldAtKey("sign-up.attempt", id, 1)
			}
			return
		}
		if !volatile.CheckFieldExistence("sign-up.attempt", id) {
			volatile.SetFieldAtKey("sign-up.attempt", id, 1)
		} else {
			volatile.IncrementFieldAtKeyBy("sign-up.attempt", id, 1)
			numberOfAttempts := volatile.GetFieldAtKey("sign-up.attempt", id).ToInt()
			if numberOfAttempts >= attemptsUntilBlock {
				volatile.SetFieldAtKey("sign-up.blocked", id, nowMoment)
			}
		}
	})
}

// RegisterSuccessfulSignUp records a successful sign-up attempt (in order to clear it)
func RegisterSuccessfulSignUp(id string) {
	volatile.TransactionWrapper(func() {
		volatile.DeleteFieldAtKey("sign-up.attempt", id)
		volatile.DeleteFieldAtKey("sign-up.blocked", id)
	})
}
