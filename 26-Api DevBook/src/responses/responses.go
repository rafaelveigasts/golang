package responses

import (
	"encoding/json"
	"log"
	"net/http"
)

type IInput interface{}

func JSON(w http.ResponseWriter, statusCode int, input IInput){
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if input != nil {
		if err := json.NewEncoder(w).Encode(input); err != nil {
			log.Fatal(err)
		}
	}

}

func Error(w http.ResponseWriter, statusCode int, err error){
	JSON(w, statusCode, struct {
		Error string `json:"error"`
	}{
		Error: err.Error(),
	})
}