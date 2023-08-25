package policy

import (
	"github.com/earaujoassis/space/internal/services/volatile"
)

// SignInAttemptStatus checks and controls sign-in attempts from a Web browser/User
func SignInAttemptStatus(id string) string {
	var result string

	volatile.TransactionsWrapper(func() {
		if volatile.CheckFieldExistence("sign-in.blocked", id) {
			result = Blocked
		} else if volatile.CheckFieldExistence("sign-in.attempt", id) {
			numberOfAttempts := volatile.GetFieldAtKey("sign-in.attempt", id).ToInt()
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
func SignUpAttemptStatus(id string) string {
	var result string

	volatile.TransactionsWrapper(func() {
		if volatile.CheckFieldExistence("sign-up.blocked", id) {
			result = Blocked
		} else if volatile.CheckFieldExistence("sign-up.attempt", id) {
			numberOfAttempts := volatile.GetFieldAtKey("sign-up.attempt", id).ToInt()
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
