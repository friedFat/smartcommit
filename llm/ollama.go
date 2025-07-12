package llm

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
)

type OllamaProvider struct {
	Model string
}

func (o *OllamaProvider) Name() string {
	return "ollama"
}

func (o *OllamaProvider) Generate(prompt string) (string, error) {
	payload := map[string]interface{}{
		"model":  o.Model,
		"prompt": prompt,
		"stream": false,
	}

	body, _ := json.Marshal(payload)

	resp, err := http.Post("http://localhost:11434/api/generate", "application/json", bytes.NewBuffer(body))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var result struct {
		Response string `json:"response"`
	}

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return "", err
	}
	if result.Response == "" {
		return "", errors.New("no response from model")
	}

	return result.Response, nil
}
