package scope

import (
	"github.com/manifoldco/promptui"
	"github.com/pkg/errors"
)

func RunSelecter() (string, error) {
	scs, err := GitConfigScopes()
	if err != nil {
		return "", errors.Wrapf(err, "run scope selector")
	}
	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}?",
		Active:   "\U0001F695 {{ .ConfigScope | cyan }} ({{ .FileName | blue }})",
		Inactive: "  {{ .ConfigScope | blue }} ({{ .FileName | blue }})",
		Selected: "\U0001F695 selected git config {{ .ConfigScope | yellow }} {{ .FileName | cyan }}",
		Details: `
--------- Selected Scope ----------
{{ "Effective scope:" | faint }}	{{ .ConfigScope }}
{{ "Git config:" | faint }}	{{ .FileName }}`,
	}

	prompt := promptui.Select{
		Label:     "Git config scope",
		Items:     scs,
		Templates: templates,
		Size:      2,
	}

	i, _, err := prompt.Run()

	if err != nil {
		return "", errors.Wrapf(err, "choose scope")
	}

	return scs[i].FileName, nil
}
