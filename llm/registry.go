package llm

import "smartcommit/config"

func GetProvider(cfg *config.Config) (Provider, error) {
	switch cfg.Provider {
	case "ollama":
		return &OllamaProvider{Model: cfg.Model}, nil
	default:
		return nil, ErrUnsupportedProvider(cfg.Provider)
	}
}

func ErrUnsupportedProvider(name string) error {
	return &UnsupportedProvider{name}
}

type UnsupportedProvider struct {
	Name string
}

func (e *UnsupportedProvider) Error() string {
	return "unsupported provider: " + e.Name
}
