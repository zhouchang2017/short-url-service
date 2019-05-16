package http

import (
	"encoding/json"
	"net/http"
)


func Json(w http.ResponseWriter, data interface{}, code int) {
	marshal, err := json.Marshal(data)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(marshal)

}
