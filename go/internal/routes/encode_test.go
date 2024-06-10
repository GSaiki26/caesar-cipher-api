// Libs
package routes

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Functions
func TestPostEncode(t *testing.T) {
	t.Run("Encode 'Hello World!'", func(t *testing.T) {
		// Data
		body := ReqEncodeBody{
			Content: "Hello World!",
		}
		var b bytes.Buffer
		if err := json.NewEncoder(&b).Encode(body); err != nil {
			t.Error(err)
		}

		// Request
		res := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, "/encode", &b)
		req.Header.Set("Content-Type", "application/json")

		PostEncode(res, req)

		// Test
		if res.Code != http.StatusOK {
			t.Errorf("Expected status code %v, got %v", http.StatusOK, res.Code)
		}

		// Decode the response body
		var resBody ResEncodeBody
		if err := json.NewDecoder(res.Body).Decode(&resBody); err != nil {
			t.Error(err)
		}

		// Check the results
		if len(resBody.Results) != 26 {
			t.Errorf("Expected 26 results, got %v", len(resBody.Results))
		}
		if resBody.Results[0].Text != "Hello World!" {
			t.Errorf("Expected 'Hello World!', got %v", resBody.Results[0].Text)
		}
		if resBody.Results[1].Text != "Ifmmp Xpsme!" {
			t.Errorf("Expected 'Ifmmp Xpsme!', got %v", resBody.Results[1].Text)
		}
		if resBody.Results[25].Text != "Gdkkn Vnqkc!" {
			t.Errorf("Expected 'Gdkkn Vnqkc!', got %v", resBody.Results[25].Text)
		}
	})

	t.Run("Invalid content type", func(t *testing.T) {
		// Request
		res := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, "/encode", nil)

		PostEncode(res, req)

		// Test
		if res.Code != http.StatusUnsupportedMediaType {
			t.Errorf("Expected status code %v, got %v", http.StatusUnsupportedMediaType, res.Code)
		}
	})

	t.Run("Invalid request body", func(t *testing.T) {
		// Data
		var b bytes.Buffer
		b.WriteString("invalid")

		// Request
		res := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodPost, "/encode", &b)
		req.Header.Set("Content-Type", "application/json")

		PostEncode(res, req)

		// Test
		if res.Code != http.StatusBadRequest {
			t.Errorf("Expected status code %v, got %v", http.StatusBadRequest, res.Code)
		}
	})
}
