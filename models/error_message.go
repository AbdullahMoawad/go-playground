package models

type Log struct {
	Status  int
	Message string
}

func Logger(status int,message string) Log  {
		logger := Log{status ,message }
		return logger
}