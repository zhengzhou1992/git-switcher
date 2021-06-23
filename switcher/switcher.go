package switcher

import (
	"github.com/pkg/errors"
	"github.com/zhengzhou1992/git-switcher/confwriter"
	"github.com/zhengzhou1992/git-switcher/user"
)

func Switch(gitConfPath string, user user.User) error {
	err := confwriter.WriteGitConfig(gitConfPath, user.Name, user.Email)
	return errors.Wrapf(err, "switch %v to user %v", gitConfPath, user)
}
