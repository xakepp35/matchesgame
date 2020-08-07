package controller

import (
	"net/http"

	log "github.com/sirupsen/logrus"
)

// SendErrorResponse синтаксический сахар проставляет код ошибки в заголовок и текст ошибки в тело.
func SendErrorResponse(w http.ResponseWriter, resCode int, errorDesc string) {
	log.Errorf(errorDesc)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(resCode)
	w.Write([]byte("{error:\"" + errorDesc + "\"}"))
}
