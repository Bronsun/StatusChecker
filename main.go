package main

import (
	"github.com/Bronsun/StatusChecker/migrations"
	"github.com/Bronsun/StatusChecker/server"
)

func main() {
	migrations.Migrate()
	/*
		db := helpers.ConnectDB()
		userlink := []interfaces.UserLink{}

		db.Select("link").Find(&userlink)
		for i := 0; i < len(userlink); i++ {
			t := interfaces.UserLink{Link: userlink}
			log.Println(t)
		}
	*/
	server.Start()

}
