package handlers

import (
	"encoding/json"
	"net/http"
	"log/slog"
	"github.com/jamie-belanger/personal-quotes/internal/app"
)

type ErrorResponse struct {
	Message string `json:"message"`
}


/*
	Writes a JSON error message to the output stream, along with HTTP status code
	
	# Parameters
	
	- w (http.ResponseWriter) = reference to the current response writer

	- statusCode (int) = status code to write (ie 400, 500)

	- message (string) = what to write to the output (ie "Bad request", "Malformed input", etc)
*/
func sendJsonErrorMessage(a *app.Application, w http.ResponseWriter, statusCode int, message string) {
	a.Logger.Info("sendJsonErrorMessage", slog.Int("statusCode", statusCode), slog.String("message", message))

	// Input validation
	if "" == message {
		message = "Unknown error"
	}
	// Note that I'm not testing every single possible value here
	switch {
	case statusCode >= 100 && statusCode < 400:
		a.Logger.Warn("sendJsonErrorMessage", slog.String("reason", "called with non-error code"), slog.Int("statusCode", statusCode))
	case statusCode >= 400 && statusCode <= 599:
		// ok! well, usually
	default:
		a.Logger.Error("sendJsonErrorMessage", slog.String("reason", "called with unknown or invalid code"), slog.Int("statusCode", statusCode))
		statusCode = http.StatusInternalServerError
	}

	response := ErrorResponse{
		Message: message,
	}

	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(response)
}
