package llm

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	"smartcommit/config"
)

type HTTPProvider struct {
    cfg *config.Config
}

func NewHTTPProvider(cfg *config.Config) (*HTTPProvider, error) {
    if cfg.BaseURL == "" || cfg.Model == "" {
        return nil, errors.New("base_url and model required")
    }
    return &HTTPProvider{cfg}, nil
}

func (h *HTTPProvider) Generate(prompt string) (string, error) {
    payload := map[string]interface{}{
        "model":  h.cfg.Model,
        "prompt": prompt,
    }
    body, _ := json.Marshal(payload)

    req, err := http.NewRequest("POST", h.cfg.BaseURL, bytes.NewReader(body))
    if err != nil {
        return "", err
    }
    req.Header.Set("Content-Type", "application/json")
    if h.cfg.APIKey != "" {
        req.Header.Set("Authorization", "Bearer "+h.cfg.APIKey)
    }
    for k, v := range h.cfg.Headers {
        req.Header.Set(k, v)
    }

    resp, err := http.DefaultClient.Do(req)
    if err != nil {
        return "", err
    }
    defer resp.Body.Close()
    out, _ := io.ReadAll(resp.Body)

    var js map[string]interface{}
    if err := json.Unmarshal(out, &js); err != nil {
        return "", err
    }
    // Try common fields
    if v, ok := js["response"]; ok {
        return strings.TrimSpace(v.(string)), nil
    }
    if v, ok := js["content"]; ok {
        return strings.TrimSpace(v.(string)), nil
    }
    return "", fmt.Errorf("unexpected response: %s", string(out))
}
