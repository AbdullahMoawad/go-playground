package models

type Log struct {
	Status  int
	Message string
	Error   error
}

func Logger(status int, message string, error error) Log {
	logger := Log{status, message, error}
	return logger
}
