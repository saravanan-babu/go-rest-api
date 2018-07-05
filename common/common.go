package common

import (
	"encoding/json"
	"net/http"
)

func ToJsonHttpResponse(w http.ResponseWriter, code int, data interface{}) {
	response, _ := json.Marshal(data)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
