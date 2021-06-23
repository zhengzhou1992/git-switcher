package user

import (
	"net/mail"

	"github.com/manifoldco/promptui"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UserService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{db: db}
}

type User struct {
	Name  string `gorm:"primaryKey"`
	Email string
}

func (us *UserService) AllUsers() (users []User, err error) {
	result := us.db.Find(&users)
	return users, errors.Wrapf(result.Error, "fetch all users in sqlite db")
}

func (us *UserService) UpsertUsers(users []User) error {
	result := us.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "name"}},
		DoUpdates: clause.AssignmentColumns([]string{"email"}),
	}).Create(&users)
	return result.Error
}

func (us *UserService) RunSelecter() (User, error) {
	users, err := us.AllUsers()
	if err != nil {
		return User{}, errors.Wrapf(err, "run user selector")
	}
	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}?",
		Active:   "\U0001F64B {{ .Name | cyan }} ({{ .Email | blue }})",
		Inactive: "  {{ .Name | blue }} ({{ .Email | blue }})",
		Selected: "\U0001F64B selected user  {{ .Name | yellow }} {{ .Name | cyan }}",
		Details: `
--------- Selected User ----------
{{ "Name:" | faint }}	{{ .Name }}
{{ "Email:" | faint }}	{{ .Email }}`,
	}

	prompt := promptui.Select{
		Label:     "User",
		Items:     users,
		Templates: templates,
		Size:      5,
	}
	i, _, err := prompt.Run()

	if err != nil {
		return User{}, errors.Wrapf(err, "choose user")
	}

	return users[i], nil
}

func (us *UserService) AddUser() (err error) {
	validateName := func(input string) error {
		if len(input) < 4 {
			return errors.New("Name must have more than 4 characters")
		}
		return nil
	}
	promptName := promptui.Prompt{
		Label:    "Name",
		Validate: validateName,
	}
	var name string
	name, err = promptName.Run()
	if err != nil {
		return
	}

	validateEmail := func(email string) error {
		_, e := mail.ParseAddress(email)
		return e
	}
	promptEmail := promptui.Prompt{
		Label:    "Email",
		Validate: validateEmail,
	}
	var email string
	email, err = promptEmail.Run()
	if err != nil {
		return
	}

	err = us.UpsertUsers([]User{{
		Name:  name,
		Email: email,
	}})
	return errors.Wrapf(err, "add user %s / %s", name, email)
}
