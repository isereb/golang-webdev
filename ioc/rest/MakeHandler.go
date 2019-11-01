package rest

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func (b *BaseHandler) GetAllMakes(writer http.ResponseWriter, request *http.Request) {
	b.HandlerLog.Println("Somebody requested all makes")
	makeService := b.getService("make_service")
	bytes, _ := json.Marshal(makeService.GetAllMakes())
	_, _ = writer.Write(bytes)
}

func (b *BaseHandler) GetMake(writer http.ResponseWriter, request *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(request)["make"])
	b.HandlerLog.Println("Somebody requested make with id =", id)

	makeService := b.getService("make_service")
	amake, err := makeService.GetMakeById(id)

	if err != nil {
		b.handleError(writer, request, requestError{
			Header:  "Not found",
			Message: fmt.Sprint("Model with id ", id, " not found"),
		})
	} else {
		j, _ := json.Marshal(amake)
		_, _ = writer.Write(j)
	}
}
