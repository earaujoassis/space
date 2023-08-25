package models

const (
	eternalExpirationLength   int64 = 0
	largestExpirationLength   int64 = 3600 // 60 min
	defaultExpirationLength   int64 = 1800 // 30 min
	shortestExpirationLength  int64 = 900  // 15 min
	machineryExpirationLength int64 = 150  // 2,5 min
)

// Tokens interface defines methods/actions for checking session-tokens
//
//	time-based validity
type Tokens interface {
	WithinExpirationWindow()
}
