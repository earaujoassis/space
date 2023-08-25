package policy

const (
	attemptsUntilPreblock int = 2
	attemptsUntilBlock    int = 5

	blockPeriodFailedSignIn int64 = 43200 // 12 hours
	blockPeriodFailedSignUp int64 = 720   // 12 minutes

	// Preblocked policy state
	Preblocked string = "preblocked"
	// Blocked policy state
	Blocked string = "blocked"
	// Clear policy state
	Clear string = "clear"
)
