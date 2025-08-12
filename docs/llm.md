# LLM & OpenAI Module

This module provides access to OpenAI's GPT models for chat and text completion from VintLang scripts.

## Setup

1. **Get an OpenAI API Key:**
   - Sign up at https://platform.openai.com/ and create an API key.
2. **Set the API Key in your environment:**
   - On macOS/Linux: `export OPENAI_API_KEY=sk-...`
   - On Windows: `set OPENAI_API_KEY=sk-...`

## Functions

### `llm.chat(messages, model="gpt-3.5-turbo", max_tokens=128, temperature=0.7)`
- **messages:** List of message objects (`{"role": "user", "content": "..."}`)
- **model:** (optional) Model name (default: gpt-3.5-turbo)
- **max_tokens:** (optional) Max tokens in response
- **temperature:** (optional) Sampling temperature
- **Returns:** (response, error)

### `llm.completion(prompt, model="text-davinci-003", max_tokens=128, temperature=0.7)`
- **prompt:** String prompt
- **model:** (optional) Model name (default: text-davinci-003)
- **max_tokens:** (optional) Max tokens in response
- **temperature:** (optional) Sampling temperature
- **Returns:** (completion, error)

## Example Usage

```js
import llm

messages = [
    {"role": "system", "content": "You are a helpful assistant."},
    {"role": "user", "content": "Tell me a joke."}
]
response, err = llm.chat(messages)
if err != null {
    print("Chat error: ", err)
} else {
    print("Chat response: ", response)
}

completion, err = llm.completion("Write a poem about the stars.")
if err != null {
    print("Completion error: ", err)
} else {
    print("Completion: ", completion)
}
```

## Notes
- Requires an internet connection.
- Make sure your API key is kept secret.
- See OpenAI docs for more on models and parameters. 