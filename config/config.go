// File: config/config.go
package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Provider     string `yaml:"provider"`
	Model        string `yaml:"model"`
	APIKey       string `yaml:"api_key"`
	SystemPrompt string `yaml:"system_prompt"`
}

func defaultConfig() *Config {
	return &Config{
		Provider:     "ollama",
		Model:        "llama3",
		SystemPrompt: "You are a terse, developer-friendly AI who writes conventional commit messages. Prefer single-line messages like 'feat: add ...' or 'fix: resolve ...'. Avoid PR-style summaries, bullet points, also this is official commit message. JUST SHOW THE COMMIT MESSAGE, NO OTHER BS. DONT GIVE ANY OTHER SUGGESTIONS OR BRACKETS OR ANYTHING, THIS ARE LIKE LITERAL COMMITS! DONT MESS THEM UP JUST GIVE THE COMMIT!",
	}
}

func configPath() string {
	configDir, _ := os.UserConfigDir()
	full := filepath.Join(configDir, "smartcommit", "config.yaml")
	fmt.Println("🕵️ Config path being used:", full) // <- add this
	return full
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
	return &cfg
}

func Save(cfg *Config) error {
	path := configPath()
	os.MkdirAll(filepath.Dir(path), 0755)
	data, err := yaml.Marshal(cfg)
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0644)
}

func (cfg *Config) Set(key, value string) {
	switch strings.ToLower(key) {
	case "provider":
		cfg.Provider = value
	case "model":
		cfg.Model = value
	case "api_key":
		cfg.APIKey = value
	case "system_prompt":
		cfg.SystemPrompt = value
	}
}

func (cfg *Config) PrettyPrint() {
	fmt.Println("Current config:")
	fmt.Println("  provider:      ", cfg.Provider)
	fmt.Println("  model:         ", cfg.Model)
	if cfg.APIKey != "" {
		fmt.Println("  api_key:       ", mask(cfg.APIKey))
	}
	fmt.Println("  system_prompt: ", cfg.SystemPrompt)
}

func mask(s string) string {
	if len(s) <= 6 {
		return "******"
	}
	return s[:3] + "..." + s[len(s)-3:]
}
