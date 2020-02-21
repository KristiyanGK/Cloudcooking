package middlewares

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"fmt"
	"net/http"
)

// LoggingMiddleware logs information about every request to the console
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("url: ", r.URL, " with method: ", r.Method)

		writer := httptest.NewRecorder()
		next.ServeHTTP(writer, r)

		var prettyJSON bytes.Buffer
		json.Indent(&prettyJSON, writer.Body.Bytes(), "", "\t")

		fmt.Println("Response: ")
		fmt.Println(string(prettyJSON.Bytes()))

		for h, hvalues := range writer.Result().Header {
			for _, hvalue := range hvalues {
				w.Header().Set(h, hvalue)
			}
		}

		w.WriteHeader(writer.Code)
		w.Write(writer.Body.Bytes())
	})
}
