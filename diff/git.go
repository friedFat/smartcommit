package diff

import (
	"os/exec"
)

func GetStagedDiff() (string, error) {
	out, err := exec.Command("git", "diff", "--cached").Output()
	if err != nil {
		return "", err
	}
	return string(out), nil
}
