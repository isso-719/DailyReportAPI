package response

import (
	"encoding/json"
	"log"
	"net/http"
)

// Success HTTPコード:200 正常終了を処理する
func Success(writer http.ResponseWriter, response interface{}) {
	if response == nil {
		return
	}
	data, err := json.Marshal(response)
	if err != nil {
		log.Println(err)
		InternalServerError(writer, "marshal error")
		return
	}
	if _, err := writer.Write(data); err != nil {
		log.Println(err)
	}
}

// BadRequest HTTPコード:400 BadRequestを処理する
func BadRequest(writer http.ResponseWriter, message string) {
	httpError(writer, http.StatusBadRequest, message)
}

// InternalServerError HTTPコード:500 InternalServerErrorを処理する
func InternalServerError(writer http.ResponseWriter, message string) {
	httpError(writer, http.StatusInternalServerError, message)
}

// Unauthorized HTTPコード:401 Unauthorizedを処理する
func Unauthorized(writer http.ResponseWriter, message string) {
	httpError(writer, http.StatusUnauthorized, message)
}

// httpError エラー用のレスポンス出力を行う
func httpError(writer http.ResponseWriter, code int, message string) {
	data, _ := json.Marshal(errorResponse{
		Code:    code,
		Message: message,
	})
	writer.WriteHeader(code)
	if data != nil {
		if _, err := writer.Write(data); err != nil {
			log.Println(err)
			return
		}
	}
}

type errorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
