package handlers

import (
	"encoding/json"
	"net/http"
	"github.com/jamie-belanger/personal-quotes/internal/app"
)

type QuoteCreateResponse struct {
	Id string `json:"id"`
}
type QuoteRetrieveResponse struct {
	QuoteBody string `json:"quote"`
	QuoteAuthor string `json:"author"`
}


/*
	Retrieves a random quote from the database

	# Returns
	
	QuoteRetrieveResponse
*/
func GetRandomQuote(a *app.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		if _, err := a.Database.GetRandomQuote(); err == nil {
	 		json.NewEncoder(w).Encode(QuoteRetrieveResponse{ QuoteBody: "", QuoteAuthor: "" })
		} else {
		 	sendJsonErrorMessage(a, w, http.StatusNotFound, err.Error())
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
		quoteId := r.PathValue("id")
		if "" == quoteId {
			sendJsonErrorMessage(a, w, http.StatusBadRequest, "ID not provided in URL")
			return
		}

		if _, err := a.Database.GetQuote(quoteId); err == nil {
			json.NewEncoder(w).Encode(QuoteRetrieveResponse{ QuoteBody: "", QuoteAuthor: "" })
		} else {
			sendJsonErrorMessage(a, w, http.StatusNotFound, err.Error())
		}
	}
}



// /*
// 	Creates a shorted link
	
// 	# Parameters (form)

// 	- link (string) = the link to shorten and store
	
// 	# Returns
	
// 	string = the short slug you can use to retrieve the link later
// */
// func (a *application) shortLinkCreate(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")

// 	// Validate link in payload
// 	userLink := r.FormValue("link")
// 	if 0 == len(userLink) {
// 		a.sendJsonErrorMessage(w, http.StatusBadRequest, "Payload form does not contain definition for 'link' value")
// 		return
// 	}

// 	// Since this is intended to store hyperlinks, let's ensure we have one:
// 	if "http://" != userLink[:7] && "https://" != userLink[:8] {
// 		a.sendJsonErrorMessage(w, http.StatusBadRequest, "Provided link does not look like a valid URL")
// 		return
// 	}

// 	// Generate short slug. Passing a GUID in here instead of the link would make this a lot more random.
// 	slug := a.generateSlug(userLink)

// 	// Add to database
// 	if err := a.InsertLink(slug, userLink); err != nil {
// 		a.sendJsonErrorMessage(w, http.StatusConflict, err.Error())
// 		return
// 	}

// 	// return slug to caller
// 	w.WriteHeader(http.StatusCreated)
// 	json.NewEncoder(w).Encode(createResponse{ Slug: slug })
// }

// /*
// 	Retrieves a stored hyperlink using the slug provided in the URL

// 	# Parameters (URL)

// 	- id (string) = a slug to retrieve

// 	# Returns
	
// 	string = the URL that is stored for this slug
// */
// func (a *application) shortLinkGet(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")

// 	// First ensure the link ID exists in path; otherwise it's a 400 Bad Request
// 	linkId := r.PathValue("id")
// 	if "" == linkId {
// 		a.sendJsonErrorMessage(w, http.StatusBadRequest, "ID not provided in URL")
// 		return
// 	}

// 	// Check the database for this ID
// 	if link, err := a.GetLink(linkId); err == nil {
// 		json.NewEncoder(w).Encode(retrieveResponse{ Link: link })
// 	} else {
// 		a.sendJsonErrorMessage(w, http.StatusNotFound, err.Error())
// 	}
// }

// /*
// 	Deletes a stored hyperlink using the slug provided in the URL

// 	# Parameters (URL)

// 	- id (string) = a slug to retrieve
// */
// func (a *application) shortLinkDelete(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")

// 	// First ensure the link ID exists in path; otherwise it's a 400 Bad Request
// 	linkId := r.PathValue("id")
// 	if "" == linkId {
// 		a.sendJsonErrorMessage(w, http.StatusBadRequest, "ID not provided in URL")
// 		return
// 	}

// 	// Try to delete link
// 	if err := a.RemoveLink(linkId); err == nil {
// 		w.WriteHeader(http.StatusNoContent)
// 	} else {
// 		a.sendJsonErrorMessage(w, http.StatusNotFound, err.Error())
// 	}
// }