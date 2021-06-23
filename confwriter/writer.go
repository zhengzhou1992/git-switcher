package confwriter

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/pkg/errors"
)

func writeGitConfigItem(gitConfPath, key, value string) error {

	gitEnv := fmt.Sprintf("GIT_CONFIG=%s", gitConfPath)

	cmd := exec.Command("git", "config", key, value)
	cmd.Env = append(os.Environ(), gitEnv)
	if err := cmd.Run(); err != nil {
		return errors.Wrapf(err, "setting key %v to %v", key, value)
	}
	return nil
}

func WriteGitConfig(gitConfPath, user, email string) error {
	err := writeGitConfigItem(gitConfPath, "user.name", user)
	if err != nil {
		return err
	}
	err = writeGitConfigItem(gitConfPath, "user.email", email)
	fmt.Printf("✔ configuration has taken effect!\n")
	return err
}
