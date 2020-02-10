package responses

import (
	"fmt"
	"net/http"
)

//CreatedResponse writes a 201 created response with location header of resource
func CreatedResponse(w http.ResponseWriter, base string, id uint) {
	location := fmt.Sprintf("/%s/%d", base, id)

	w.Header().Set("Location", location)
	w.WriteHeader(http.StatusCreated)

	w.Write([]byte("Resource created!"))
}
