package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/jamie-belanger/personal-quotes/internal/app"
	"github.com/jamie-belanger/personal-quotes/internal/models"
)

type QuoteCreateResponse struct {
	Id int64 `json:"id"`
}


/*
	Retrieves a random quote from the database

	# Returns
	
	QuoteRetrieveResponse
*/
func GetRandomQuote(a *app.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		if quote, err := a.Database.GetRandomQuote(); err == nil {
	 		json.NewEncoder(w).Encode(quote)
		} else {
		 	sendJsonErrorMessage(a, w, http.StatusNotFound, err.Error(), err)
		}
	}
}


/*
	Retrieves a stored quote using the id provided in the URL

	# Parameters (URL)

	- id (string) = a quote to retrieve

	# Returns
	
	QuoteRetrieveResponse
*/
func GetQuote(a *app.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		// First ensure the link ID exists in path; otherwise it's a 400 Bad Request
		id, err := tryParseId(r.PathValue("id"))
		if nil != err {
			sendJsonErrorMessage(a, w, http.StatusBadRequest, "Invalid or out of range ID", err)
			return
		}

		if quote, err := a.Database.GetQuote(id); err == nil {
			json.NewEncoder(w).Encode(quote)
		} else {
			sendJsonErrorMessage(a, w, http.StatusNotFound, err.Error(), err)
		}
	}
}



/*
	Saves the given data to the database

	# Parameters (FORM)

	- body (string) = text to save for quote body

	- author (string) = text to save for quote author

	# Returns
	
	QuoteCreateResponse = the newly saved quote Id
*/
func CreateQuote(a *app.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		// Load data
		body, err := sanitizeHTML(r.FormValue("body"))
		if nil != err {
			sendJsonErrorMessage(a, w, http.StatusBadRequest, "Error parsing body text", err)
			return
		}

		author, err := sanitizeHTML(r.FormValue("author"))
		if nil != err {
			sendJsonErrorMessage(a, w, http.StatusBadRequest, "Error parsing author text", err)
			return
		}

		quote := &models.Quote{
			Id: 0,
			Body: body,
			Author: author,
		}

		if quoteId, err := a.Database.SaveQuote(quote); err == nil {
			json.NewEncoder(w).Encode(QuoteCreateResponse{ Id: quoteId })
		} else {
			sendJsonErrorMessage(a, w, http.StatusNotFound, err.Error(), err)
		}
	}
}



/*
	Saves the given data to the database

	# Parameters (URL)

	- id (string) = a quote to retrieve

	# Parameters (FORM)

	- body (string) = text to save for quote body

	- author (string) = text to save for quote author

	# Returns
	
	HTML response code and/or error message
*/
func UpdateQuote(a *app.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		// First ensure the link ID exists in path; otherwise it's a 400 Bad Request
		id, err := tryParseId(r.PathValue("id"))
		if nil != err {
			sendJsonErrorMessage(a, w, http.StatusBadRequest, "Invalid or out of range ID", err)
			return
		}

		// Load data
		body, err := sanitizeHTML(r.FormValue("body"))
		if nil != err {
			sendJsonErrorMessage(a, w, http.StatusBadRequest, "Error parsing body text", err)
			return
		}

		author, err := sanitizeHTML(r.FormValue("author"))
		if nil != err {
			sendJsonErrorMessage(a, w, http.StatusBadRequest, "Error parsing author text", err)
			return
		}

		quote := &models.Quote{
			Id: id,
			Body: body,
			Author: author,
		}

		if _, err := a.Database.SaveQuote(quote); err == nil {
			w.WriteHeader(http.StatusNoContent)
		} else {
			sendJsonErrorMessage(a, w, http.StatusNotFound, err.Error(), err)
		}
	}
}


/*
	Deletes the given quote from the database

	# Parameters (URL)

	- id (string) = a quote to retrieve

	# Returns
	
	HTML response code and/or error message
*/
func DeleteQuote(a *app.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		// First ensure the link ID exists in path; otherwise it's a 400 Bad Request
		id, err := tryParseId(r.PathValue("id"))
		if nil != err {
			sendJsonErrorMessage(a, w, http.StatusBadRequest, "Invalid or out of range ID", err)
			return
		}

		if err := a.Database.DeleteQuote(id); err == nil {
			w.WriteHeader(http.StatusNoContent)
		} else {
			sendJsonErrorMessage(a, w, http.StatusNotFound, err.Error(), err)
		}
	}
}