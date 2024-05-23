package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

func Recovery(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Println(err)
				w.WriteHeader(500)
			}
		}()
		next.ServeHTTP(w, r)
	})
}

func Response(w http.ResponseWriter, result interface{}, errorCode int) {
	data := map[string]interface{}{
		"result": result,
		"error":  errorCode,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(errorCode)
	json.NewEncoder(w).Encode(data)
}
