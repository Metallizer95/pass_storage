package logging

type Level string

const (
	DEBUG   Level = "debug"
	WARNING Level = "warning"
	INFO    Level = "info"
	ERROR   Level = "error"
	FATAL   Level = "fatal"
)
