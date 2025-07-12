package prompt

import "fmt"

func Build(diff string, systemPrompt string) string {
	return fmt.Sprintf(`%s

Here is the Git diff:
---
%s
---
Write a commit message:`, systemPrompt, diff)
}
