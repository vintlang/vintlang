package module

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/vintlang/vintlang/object"
)

const (
	openaiChatEndpoint       = "https://api.openai.com/v1/chat/completions"
	openaiCompletionEndpoint = "https://api.openai.com/v1/completions"
	defaultModel             = "gpt-3.5-turbo"
)

// Message represents a message for OpenAI chat API
// Role: "system", "user", or "assistant"
type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type ChatRequest struct {
	Model       string    `json:"model"`
	Messages    []Message `json:"messages"`
	Stream      bool      `json:"stream,omitempty"`
	MaxTokens   int       `json:"max_tokens,omitempty"`
	Temperature float64   `json:"temperature,omitempty"`
}

type ChatChoice struct {
	Index        int     `json:"index"`
	Message      Message `json:"message"`
	FinishReason string  `json:"finish_reason"`
}

type ChatResponse struct {
	Choices []ChatChoice `json:"choices"`
}

type CompletionRequest struct {
	Model       string  `json:"model"`
	Prompt      string  `json:"prompt"`
	MaxTokens   int     `json:"max_tokens,omitempty"`
	Temperature float64 `json:"temperature,omitempty"`
}

type CompletionChoice struct {
	Text         string `json:"text"`
	Index        int    `json:"index"`
	FinishReason string `json:"finish_reason"`
}

type CompletionResponse struct {
	Choices []CompletionChoice `json:"choices"`
}

// LLMFunctions exposes OpenAI functions to VintLang
var LLMFunctions = map[string]object.ModuleFunction{
	"chat":       llmChat,
	"completion": llmCompletion,
}

// llmChat is the VintLang wrapper for OpenAI chat
func llmChat(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) < 1 {
		return &object.Error{Message: "llm.chat expects at least 1 argument (messages)"}
	}
	msgs, ok := args[0].(*object.Array)
	if !ok {
		return &object.Error{Message: "llm.chat expects an array of messages as the first argument"}
	}
	var messages []Message
	for _, m := range msgs.Elements {
		msgMap, ok := m.(*object.Dict)
		if !ok {
			return &object.Error{Message: "Each message must be a dictionary with 'role' and 'content'"}
		}
		roleObj, ok1 := msgMap.Pairs[(&object.String{Value: "role"}).HashKey()]
		contentObj, ok2 := msgMap.Pairs[(&object.String{Value: "content"}).HashKey()]
		if !ok1 || !ok2 {
			return &object.Error{Message: "Each message must have 'role' and 'content'"}
		}
		role, _ := roleObj.Value.(*object.String)
		content, _ := contentObj.Value.(*object.String)
		messages = append(messages, Message{Role: role.Value, Content: content.Value})
	}
	model := defaultModel
	maxTokens := 128
	temperature := 0.7
	if len(args) > 1 {
		if s, ok := args[1].(*object.String); ok {
			model = s.Value
		}
	}
	if len(args) > 2 {
		if i, ok := args[2].(*object.Integer); ok {
			maxTokens = int(i.Value)
		}
	}
	if len(args) > 3 {
		if f, ok := args[3].(*object.Float); ok {
			temperature = f.Value
		}
	}
	resp, err := CallOpenAIChat(messages, model, maxTokens, temperature)
	if err != nil {
		return &object.Error{Message: err.Error()}
	}
	return &object.String{Value: resp}
}

// llmCompletion is the VintLang wrapper for OpenAI completion
func llmCompletion(args []object.VintObject, defs map[string]object.VintObject) object.VintObject {
	if len(args) < 1 {
		return &object.Error{Message: "llm.completion expects at least 1 argument (prompt)"}
	}
	promptObj, ok := args[0].(*object.String)
	if !ok {
		return &object.Error{Message: "llm.completion expects a string prompt as the first argument"}
	}
	prompt := promptObj.Value
	model := "text-davinci-003"
	maxTokens := 128
	temperature := 0.7
	if len(args) > 1 {
		if s, ok := args[1].(*object.String); ok {
			model = s.Value
		}
	}
	if len(args) > 2 {
		if i, ok := args[2].(*object.Integer); ok {
			maxTokens = int(i.Value)
		}
	}
	if len(args) > 3 {
		if f, ok := args[3].(*object.Float); ok {
			temperature = f.Value
		}
	}
	resp, err := CallOpenAICompletion(prompt, model, maxTokens, temperature)
	if err != nil {
		return &object.Error{Message: err.Error()}
	}
	return &object.String{Value: resp}
}

// getAPIKey fetches the OpenAI API key from env or config
func getAPIKey() (string, error) {
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		return "", errors.New("OPENAI_API_KEY not set in environment")
	}
	return apiKey, nil
}

// CallOpenAIChat sends a chat request to OpenAI and returns the assistant's reply
func CallOpenAIChat(messages []Message, model string, maxTokens int, temperature float64) (string, error) {
	apiKey, err := getAPIKey()
	if err != nil {
		return "", err
	}
	if model == "" {
		model = defaultModel
	}
	chatReq := ChatRequest{
		Model:       model,
		Messages:    messages,
		MaxTokens:   maxTokens,
		Temperature: temperature,
	}
	body, _ := json.Marshal(chatReq)
	client := &http.Client{Timeout: 30 * time.Second}
	req, _ := http.NewRequest("POST", openaiChatEndpoint, bytes.NewBuffer(body))
	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	respBody, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != 200 {
		return "", errors.New(string(respBody))
	}
	var chatResp ChatResponse
	err = json.Unmarshal(respBody, &chatResp)
	if err != nil {
		return "", err
	}
	if len(chatResp.Choices) > 0 {
		return chatResp.Choices[0].Message.Content, nil
	}
	return "", nil
}

// CallOpenAICompletion sends a completion request to OpenAI and returns the result
func CallOpenAICompletion(prompt, model string, maxTokens int, temperature float64) (string, error) {
	apiKey, err := getAPIKey()
	if err != nil {
		return "", err
	}
	if model == "" {
		model = "text-davinci-003"
	}
	compReq := CompletionRequest{
		Model:       model,
		Prompt:      prompt,
		MaxTokens:   maxTokens,
		Temperature: temperature,
	}
	body, _ := json.Marshal(compReq)
	client := &http.Client{Timeout: 30 * time.Second}
	req, _ := http.NewRequest("POST", openaiCompletionEndpoint, bytes.NewBuffer(body))
	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	respBody, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != 200 {
		return "", errors.New(string(respBody))
	}
	var compResp CompletionResponse
	err = json.Unmarshal(respBody, &compResp)
	if err != nil {
		return "", err
	}
	if len(compResp.Choices) > 0 {
		return compResp.Choices[0].Text, nil
	}
	return "", nil
}
