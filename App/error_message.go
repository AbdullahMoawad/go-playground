package App

import (
	"encoding/json"
	"go.uber.org/zap"
	"net/http"
)

type Controller struct{}

// @todo extract the file to app.go

type Log struct {
	Status  int
	Message string
	Error   interface{}
}

// return json error
func (self Controller) JsonLogger(res http.ResponseWriter, status int, msg string, error interface{}) {
	log := Log{
		Status:  status,
		Message: msg,
		Error:   error,
	}
	response, _ := json.Marshal(log)
	res.Header().Set("Content-Type", "application/json")
	_, _ = res.Write(response)
}

func (self Controller) Json(res http.ResponseWriter, payload interface{}, statusCode int) {

	response, _ := json.Marshal(payload)
	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(statusCode)
	_, _ = res.Write(response)
}

// return cli error using zap logger
func (self Controller) Logger(msg, errType string) *zap.Logger {
	log, _ := zap.NewDevelopment()
	//@todo use switch

	switch errType {
	case "debug":
		log.Debug(msg)
		return log
	case "info":
		log.Info(msg)
		return log
	default:
		log.Error(msg)
		return log

	}
}

func Logger(msg, errType string) *zap.Logger {
	log, _ := zap.NewDevelopment()
	//@todo use switch

	switch errType {
	case "debug":
		log.Debug(msg)
		return log
	case "info":
		log.Info(msg)
		return log
	default:
		log.Error(msg)
		return log

	}
}
