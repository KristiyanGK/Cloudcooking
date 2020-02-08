package responses

import (
	"encoding/json"
	"net/http"
)

//JSONResponse writes a json response 
func JSONResponse(w http.ResponseWriter, data interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	jsonData, _ := json.MarshalIndent(data, "", "    ")

	w.Write(jsonData)
}
