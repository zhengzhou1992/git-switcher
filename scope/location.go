package scope

import (
	"os"
	"path/filepath"

	"github.com/pkg/errors"
)

func CurrentProject() (isInGitProject bool, projectGitPath string, err error) {
	var pwd, absPWD string
	pwd, err = os.Getwd()
	if err != nil {
		err = errors.Wrap(err, "get current path")
		return
	}
	absPWD, err = filepath.Abs(pwd)
	if err != nil {
		err = errors.Wrap(err, "get current abs path")
		return
	}
	for absPWD != "/" {
		gitPath := filepath.Join(absPWD, ".git")
		if _, e := os.Stat(gitPath); !os.IsNotExist(e) {
			isInGitProject = true
			projectGitPath = gitPath
			return
		}
		absPWD = filepath.Dir(absPWD)
	}
	return
}
