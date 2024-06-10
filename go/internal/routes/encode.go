// Libs
package routes

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/GSaiki26/caesar-cipher-api/internal/utils"
)

// Structs
type ReqEncodeBody struct {
	Content string `json:"content"`
}

type ResEncodeBody struct {
	Results []ResEncodeBodyResults `json:"results"`
}

type ResEncodeBodyResults struct {
	Shift int    `json:"shift"`
	Text  string `json:"text"`
}

// Functions
func PostEncode(
	rw http.ResponseWriter,
	r *http.Request,
) {
	log.Println("Request received on: POST /encode")

	// Check if the media type is JSON
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(rw, "The request's content type must be 'application/json'.", http.StatusUnsupportedMediaType)
		return
	}

	// Decode the request body (JSON)
	var body ReqEncodeBody
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(rw, "Couldn't parse the request's body.", http.StatusBadRequest)
		return
	}

	// Encode
	var resBody ResEncodeBody
	for i := 0; i < 26; i++ {
		result := ResEncodeBodyResults{
			Shift: i,
			Text:  utils.ConvertToCesar(body.Content, i),
		}
		resBody.Results = append(resBody.Results, result)
	}

	// Return the response
	rw.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(rw).Encode(resBody); err != nil {
		http.Error(rw, "Couldn't encode the response.", http.StatusInternalServerError)
	}
}
