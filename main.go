package main

import (
	"flag"
	"log"

	"github.com/zhengzhou1992/git-switcher/scope"
	"github.com/zhengzhou1992/git-switcher/switcher"
	"github.com/zhengzhou1992/git-switcher/user"
)

func main() {
	addUserPtr := flag.Bool("a", false, "add git user")
	flag.Parse()

	// init
	db, err := initDependency()
	if err != nil {
		log.Fatalln(err)
	}
	userSvc := user.NewUserService(db)
	err = initUser(userSvc)
	if err != nil {
		log.Fatalln(err)
	}

	// add user
	if *addUserPtr {
		err := userSvc.AddUser()
		log.Fatalln(err)
		return
	}

	// select scope
	gitConfigPath, err := scope.RunSelecter()
	if err != nil {
		panic(err)
	}

	// select or create user
	var u user.User
	u, err = userSvc.RunSelecter()
	if err != nil {
		panic(err)
	}

	// switch user
	err = switcher.Switch(gitConfigPath, u)
	if err != nil {
		panic(err)
	}
}
