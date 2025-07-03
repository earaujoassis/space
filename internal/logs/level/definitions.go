package level

type Level int8

const (
	Info Level = iota
	Error
	Critical
	Panic
)
