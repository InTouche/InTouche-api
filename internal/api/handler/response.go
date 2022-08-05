package handler

import (
	"encoding/json"
	"net/http"

	"go.uber.org/zap"
)

const (
	HeaderContentType   = "Content-Type"
	MimeApplicationJSON = "application/json"
	MimeTextYAML        = "text/yaml"
)

type ResponseManager struct {
	logger *zap.SugaredLogger
}

func NewResponseManager(logger *zap.SugaredLogger) *ResponseManager {
	return &ResponseManager{logger: logger}
}

func (rm *ResponseManager) Write(w http.ResponseWriter, code int, mime string, data []byte) {
	if data == nil {
		w.WriteHeader(code)
		return
	}

	w.Header().Set(HeaderContentType, mime)
	w.WriteHeader(code)

	if _, err := w.Write(data); err != nil {
		rm.logger.Error("failed to write response", err)
	}
}

func (rm *ResponseManager) JSON(w http.ResponseWriter, code int, data interface{}) {
	if data == nil {
		w.WriteHeader(code)
		return
	}

	w.Header().Set(HeaderContentType, MimeApplicationJSON)
	w.WriteHeader(code)

	if err := json.NewEncoder(w).Encode(&data); err != nil {
		rm.logger.Error("failed to respond with JSON", err)
	}
}

type ApiError struct {
	Error string `json:"error"`
}

func (rm *ResponseManager) Error(w http.ResponseWriter, code int, err error) {

	rm.logger.Error("got an api error", err)

	rm.JSON(w, code, err.Error())
}

func (rm *ResponseManager) OK(w http.ResponseWriter, data interface{}) {
	rm.JSON(w, http.StatusOK, data)
}

func (rm *ResponseManager) Created(w http.ResponseWriter, data interface{}) {
	rm.JSON(w, http.StatusCreated, data)
}

func (rm *ResponseManager) NoContent(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNoContent)
}
