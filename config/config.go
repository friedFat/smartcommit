package config

import (
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type Config struct {
    Provider     string            `yaml:"provider"`
    Model        string            `yaml:"model"`
    APIKey       string            `yaml:"api_key"`
    BaseURL      string            `yaml:"base_url"`
    Headers      map[string]string `yaml:"headers"`
    SystemPrompt string            `yaml:"system_prompt"`
}

func defaultConfig() *Config {
    return &Config{
        Provider:     "ollama",
        Model:        "llama3",
        BaseURL:      "http://localhost:11434/api/generate",
        Headers:      map[string]string{},
        SystemPrompt: "You are an expert at writing concise commit messages. Only return the commit text.",
    }
}

func configPath() string {
    dir, _ := os.UserConfigDir()
    return filepath.Join(dir, "smartcommit", "config.yaml")
}

func LoadOrDefault() *Config {
    path := configPath()
    data, err := os.ReadFile(path)
    if err != nil {
        return defaultConfig()
    }
    var cfg Config
    if err := yaml.Unmarshal(data, &cfg); err != nil {
        return defaultConfig()
    }
    if cfg.Headers == nil {
        cfg.Headers = map[string]string{}
    }
    return &cfg
}

func Save(cfg *Config) error {
    path := configPath()
    os.MkdirAll(filepath.Dir(path), 0755)
    out, err := yaml.Marshal(cfg)
    if err != nil {
        return err
    }
    return os.WriteFile(path, out, 0644)
}
