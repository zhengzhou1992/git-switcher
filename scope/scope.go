package scope

import (
	"os"
	"path/filepath"

	"github.com/pkg/errors"
)

type GitConfigScope string

const (
	GitConfScopeGlobal  GitConfigScope = "global"
	GitConfScopeProject GitConfigScope = "project"
)

type Scope struct {
	ConfigScope GitConfigScope
	FileName    string
}

func GitConfigScopes() (scs []Scope, err error) {
	var home string
	home, err = os.UserHomeDir()
	if err != nil {
		return
	}
	homeConfigPath := filepath.Join(home, ".gitconfig")
	scs = append(scs, Scope{
		ConfigScope: GitConfScopeGlobal,
		FileName:    homeConfigPath,
	})

	var isInProject bool
	var projectGitPath string
	isInProject, projectGitPath, err = CurrentProject()
	if err != nil {
		err = errors.Wrapf(err, "get project config scope")
		return
	}
	if isInProject {
		scs = append(scs, Scope{
			ConfigScope: GitConfScopeProject,
			FileName:    filepath.Join(projectGitPath, "config"),
		})
	}
	return
}
