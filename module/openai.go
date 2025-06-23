package module

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/vintlang/vintlang/object"
)

var OpenAIFunctions = map[string]object.ModuleFunction{}

func init() {
	OpenAIFunctions["chat"] = chat
	OpenAIFunctions["complete"] = complete
}

// chat sends a chat completion request to OpenAI's API
func chat(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 2 {
		return &object.Error{Message: "Usage: chat(api_key, message)"}
	}

	apiKey, ok1 := args[0].(*object.String)
	message, ok2 := args[1].(*object.String)

	if !ok1 || !ok2 {
		return &object.Error{Message: "Both api_key and message must be strings"}
	}

	requestBody := map[string]interface{}{
		"model": "gpt-3.5-turbo",
		"messages": []map[string]string{
			{"role": "user", "content": message.Value},
		},
	}

	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		return &object.Error{Message: fmt.Sprintf("Failed to encode JSON: %s", err)}
	}

	req, err := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", bytes.NewBuffer(jsonData))
	if err != nil {
		return &object.Error{Message: fmt.Sprintf("Failed to create request: %s", err)}
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey.Value)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return &object.Error{Message: fmt.Sprintf("Request failed: %s", err)}
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != 200 {
		return &object.Error{Message: fmt.Sprintf("OpenAI API error (%d): %s", resp.StatusCode, string(body))}
	}

	var responseData map[string]interface{}
	if err := json.Unmarshal(body, &responseData); err != nil {
		return &object.Error{Message: fmt.Sprintf("Failed to parse response: %s", err)}
	}

	choices, ok := responseData["choices"].([]interface{})
	if !ok || len(choices) == 0 {
		return &object.Error{Message: "Unexpected response format or empty response"}
	}

	firstChoice := choices[0].(map[string]interface{})
	content := firstChoice["message"].(map[string]interface{})["content"].(string)

	return &object.String{Value: content}
}

// complete sends a simple text completion request to OpenAI's API
func complete(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 2 {
		return &object.Error{Message: "Usage: complete(api_key, prompt)"}
	}

	apiKey, ok1 := args[0].(*object.String)
	prompt, ok2 := args[1].(*object.String)

	if !ok1 || !ok2 {
		return &object.Error{Message: "Both api_key and prompt must be strings"}
	}

	requestBody := map[string]interface{}{
		"model":  "text-davinci-003",
		"prompt": prompt.Value,
		"max_tokens": 100,
	}

	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		return &object.Error{Message: fmt.Sprintf("Failed to encode JSON: %s", err)}
	}

	req, err := http.NewRequest("POST", "https://api.openai.com/v1/completions", bytes.NewBuffer(jsonData))
	if err != nil {
		return &object.Error{Message: fmt.Sprintf("Failed to create request: %s", err)}
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey.Value)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return &object.Error{Message: fmt.Sprintf("Request failed: %s", err)}
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != 200 {
		return &object.Error{Message: fmt.Sprintf("OpenAI API error (%d): %s", resp.StatusCode, string(body))}
	}

	var responseData map[string]interface{}
	if err := json.Unmarshal(body, &responseData); err != nil {
		return &object.Error{Message: fmt.Sprintf("Failed to parse response: %s", err)}
	}

	choices, ok := responseData["choices"].([]interface{})
	if !ok || len(choices) == 0 {
		return &object.Error{Message: "Unexpected response format or empty response"}
	}

	firstChoice := choices[0].(map[string]interface{})
	text := firstChoice["text"].(string)

	return &object.String{Value: text}
}
