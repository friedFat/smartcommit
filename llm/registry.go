package llm

import (
	"fmt"
	"github.com/manyfacedqod/smartcommit/config"
)

// Provider is anything that can Generate a commit message from a prompt.
type Provider interface {
    Generate(prompt string) (string, error)
}

// GetProvider returns an HTTPProvider (which implements Provider) for any cfg.Provider.
func GetProvider(cfg *config.Config) (Provider, error) {
    switch cfg.Provider {
    case "ollama", "openai", "http":
        return NewHTTPProvider(cfg)
    default:
        return nil, fmt.Errorf("unsupported provider: %s", cfg.Provider)
    }
}
