package llm

type Provider interface {
	Name() string
	Generate(prompt string) (string, error)
}
