package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jm-borges/llm-voice-agent-api/config"
	openai "github.com/openai/openai-go/v3"
	"github.com/openai/openai-go/v3/option"
)

// QueryInput represents the expected JSON body for the /query endpoint.
// Example: { "message": "Hello!" }
type QueryInput struct {
	Message string `json:"message"`
}

// HandleQuery handles POST requests containing a JSON "message" field.
// It sends the received text to the OpenAI API and returns the response in JSON format.
func HandleQuery(w http.ResponseWriter, r *http.Request) {
	env := config.LoadEnv()

	input, err := decodeRequest(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	response, err := getOpenAIResponse(env, input.Message)
	if err != nil {
		http.Error(w, fmt.Sprintf("error calling OpenAI model: %v", err), http.StatusInternalServerError)
		return
	}

	respondJSON(w, map[string]string{"response": response})
}

// decodeRequest reads the HTTP request body and extracts the "message" field.
// Returns a pointer to QueryInput or an error if JSON decoding fails.
func decodeRequest(r *http.Request) (*QueryInput, error) {
	var input QueryInput

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&input); err != nil {
		return nil, fmt.Errorf("failed to decode JSON: %w", err)
	}

	if input.Message == "" {
		return nil, fmt.Errorf("the 'message' field is required")
	}

	return &input, nil
}

// getOpenAIResponse sends the message to the OpenAI model defined in .env (e.g., GPT-4o-mini)
// and returns the model's text response.
func getOpenAIResponse(env *config.EnvConfig, message string) (string, error) {
	client := openai.NewClient(option.WithAPIKey(env.OpenAIKey))
	ctx := context.Background()

	params := openai.ChatCompletionNewParams{
		Model: env.OpenAIModel,
		Messages: []openai.ChatCompletionMessageParamUnion{
			openai.UserMessage(message),
		},
	}

	resp, err := client.Chat.Completions.New(ctx, params)
	if err != nil {
		return "", fmt.Errorf("failed to get response from OpenAI: %w", err)
	}

	if len(resp.Choices) == 0 {
		return "(no response from the model)", nil
	}

	return resp.Choices[0].Message.Content, nil
}

// respondJSON sends a standardized JSON response with Content-Type header and HTTP status 200.
// If encoding fails, returns HTTP 500.
func respondJSON(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, fmt.Sprintf("failed to serialize response: %v", err), http.StatusInternalServerError)
	}
}
