package rest

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func (b *BaseHandler) GetAllModels(writer http.ResponseWriter, request *http.Request) {
	makeId, _ := strconv.Atoi(mux.Vars(request)["make"])
	b.HandlerLog.Println("Somebody requested all models")
	makeService := b.getService("make_service")
	models, e := makeService.GetAllModels(makeId)
	if e != nil {
		b.handleError(writer, request, requestError{
			Header:  "Not found",
			Message: e.Error(),
		})
	} else {
		bytes, _ := json.Marshal(models)
		_, _ = writer.Write(bytes)
	}
}

func (b *BaseHandler) GetModel(writer http.ResponseWriter, request *http.Request) {
	makeId, _ := strconv.Atoi(mux.Vars(request)["make"])
	modelId, _ := strconv.Atoi(mux.Vars(request)["model"])
	b.HandlerLog.Println("Somebody requested model with makeId =", makeId, "and modelId =", modelId)

	makeService := b.getService("make_service")
	amodel, err := makeService.GetModelById(makeId, modelId)

	if err != nil {
		b.handleError(writer, request, requestError{
			Header:  "Not found",
			Message: err.Error(),
		})
	} else {
		j, _ := json.Marshal(amodel)
		_, _ = writer.Write(j)
	}
}
