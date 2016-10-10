package policy

const (
    attemptsUntilPreblock            int = 2
    attemptsUntilBlock               int = 5

    Preblocked                    string = "preblocked"
    Blocked                       string = "blocked"
    Clear                         string = "clear"

    blockPeriodFailedSignIn        int64 = 43200 // 12 hours
    blockPeriodFailedSignUp        int64 = 720   // 12 minutes
)
