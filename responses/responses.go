package responses

import (
	"encoding"
	"net/http"

	"example.com/p-service/models"
)

func Ok(w http.ResponseWriter, body encoding.BinaryMarshaler) {
	writeResponse(w, http.StatusOK, body)
}

func Forbidden(w http.ResponseWriter, message string) {
	writeResponse(w, http.StatusForbidden, models.NewErrorResponse(message))
}

func InternalServerError(w http.ResponseWriter) {
	writeResponse(w, http.StatusInternalServerError, models.NewErrorResponse(
		"Internal server error",
	))
}

func writeResponse(
	w http.ResponseWriter,
	code int,
	response encoding.BinaryMarshaler,
) {
	encodedMessage, err := response.MarshalBinary()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("{\"error\":\"Internal server error\"}"))
	}

	w.WriteHeader(code)
	w.Write(encodedMessage)
}
