package unit

import (
	"math/rand"
)

const (
	minPort int = 50000
	maxPort int = 59999
)

func randomPort() int {
	portRange := maxPort - minPort + 1
	return rand.Intn(portRange) + minPort
}
