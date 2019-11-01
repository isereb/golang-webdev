package rest

import (
	"../services"
	"encoding/json"
	"log"
	"net/http"
)

type BaseHandler struct {
	Services   *map[string]services.BaseService
	HandlerLog *log.Logger
}

func (b *BaseHandler) getService(name string) services.BaseService {
	return (*b.Services)[name]
}

type requestError struct {
	Header  string `json:"header"`
	Message string `json:"message"`
}

func (b *BaseHandler) handleError(writer http.ResponseWriter, request *http.Request, rerr requestError) {
	b.HandlerLog.Println("Failed to respond to request: ", request.URL.RequestURI(), "\n\t",
		"Header:", rerr.Header, ", Message:", rerr.Message)
	output, _ := json.Marshal(rerr)
	_, _ = writer.Write(output)
}
