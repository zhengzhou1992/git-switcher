package main

import (
	"os"
	"path/filepath"

	"github.com/pkg/errors"
	gitconfig "github.com/tcnksm/go-gitconfig"
	"github.com/zhengzhou1992/git-switcher/user"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func initUser(us *user.UserService) error {
	gName, _ := gitconfig.Global("user.name")
	gEmail, _ := gitconfig.Global("user.email")
	lName, _ := gitconfig.Local("user.name")
	lEmail, _ := gitconfig.Local("user.email")

	var users []user.User
	if gName != "" {
		users = append(users, user.User{
			Name:  gName,
			Email: gEmail,
		})
	}
	if lName != "" && lName != gName {
		users = append(users, user.User{
			Name:  lName,
			Email: lEmail,
		})
	}
	err := us.UpsertUsers(users)
	if err != nil {
		return errors.Wrap(err, "init user, upsert users from git configs")
	}
	return nil
}

func initDependency() (db *gorm.DB, err error) {
	var home string
	home, err = os.UserHomeDir()
	if err != nil {
		err = errors.Wrap(err, "init, get user home dir")
		return
	}

	projectDir := filepath.Join(home, ".git-switcher")
	if _, err = os.Stat(projectDir); os.IsNotExist(err) {
		err = os.Mkdir(projectDir, 0755)
		if err != nil {
			err = errors.Wrap(err, "mkdir project home")
			return
		}
	}

	db, err = gorm.Open(sqlite.Open(filepath.Join(home, "/.git-switcher/user.db")), &gorm.Config{})
	if err != nil {
		err = errors.Wrap(err, "failed to connect database")
		return
	}

	err = db.AutoMigrate(&user.User{})
	if err != nil {
		err = errors.Wrap(err, "failed to init user schema")
		return
	}
	return
}
