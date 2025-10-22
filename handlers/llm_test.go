package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

// TestMain loads the .env file before running any tests.
// It skips tests if OPENAI_API_KEY is not set.
func TestMain(m *testing.M) {
	_ = godotenv.Load("../.env")

	if os.Getenv("OPENAI_API_KEY") == "" {
		println("Skipping tests: OPENAI_API_KEY not set")
		return
	}

	os.Exit(m.Run())
}

// TestHandleQuery tests the /query handler with a real OpenAI API call.
func TestHandleQuery(t *testing.T) {
	req := newTestRequest(t, "Hello, test!")
	rr := httptest.NewRecorder()

	HandleQuery(rr, req)

	assertStatusOK(t, rr)
	respData := decodeResponseJSON(t, rr)
	assertNonEmptyResponse(t, respData)
	t.Logf("OpenAI response: %s", respData["response"])
}

// newTestRequest creates an HTTP POST request with the given message.
func newTestRequest(t *testing.T, message string) *http.Request {
	t.Helper()

	input := QueryInput{Message: message}
	body, err := json.Marshal(input)
	if err != nil {
		t.Fatalf("failed to marshal input: %v", err)
	}

	req := httptest.NewRequest(http.MethodPost, "/query", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	return req
}

// assertStatusOK asserts that the ResponseRecorder returned HTTP 200.
func assertStatusOK(t *testing.T, rr *httptest.ResponseRecorder) {
	t.Helper()
	if rr.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", rr.Code)
	}
}

// decodeResponseJSON decodes the JSON body from the ResponseRecorder.
func decodeResponseJSON(t *testing.T, rr *httptest.ResponseRecorder) map[string]string {
	t.Helper()
	var resp map[string]string
	if err := json.NewDecoder(rr.Body).Decode(&resp); err != nil {
		t.Fatalf("failed to decode response JSON: %v", err)
	}
	return resp
}

// assertNonEmptyResponse asserts that the OpenAI "response" field is not empty.
func assertNonEmptyResponse(t *testing.T, resp map[string]string) {
	t.Helper()
	if resp["response"] == "" {
		t.Errorf("expected non-empty response from OpenAI")
	}
}
