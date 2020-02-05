package App

import (
	"encoding/json"
	"go.uber.org/zap"
	"net/http"
)

type Controller struct{}

type Log struct {
	Status  int
	Message string
	Error   interface{}
}

// return json error
func (self Controller) JsonLogger(res http.ResponseWriter, status int, msg string) {
	log := Log{
		Status:  status,
		Message: msg,
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

func (self Controller) Logger(errType, msg string, err interface{}) *zap.Logger {
	log, _ := zap.NewDevelopment()

	switch errType {
	case "debug":
		str := msg + (err).(string)
		log.Debug(msg + str)
		return log
	case "info":
		str := msg + (err).(string)
		log.Info(msg + str)
		return log
	default:
		str := msg + (err).(string)
		log.Error(msg + str)
		return log
	}
}

func Logger(errType, msg string, err string) *zap.Logger {
	log, _ := zap.NewDevelopment()

	switch errType {
	case "debug":
		str := msg + err
		log.Debug(msg + str)
		return log
	case "info":
		str := msg + err
		log.Info(msg + str)
		return log
	default:
		str := msg + err
		log.Error(msg + str)
		return log
	}
}
