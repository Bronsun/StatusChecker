package main

import (
	"github.com/Bronsun/StatusChecker/migrations"
	"github.com/Bronsun/StatusChecker/server"
)

func main() {
	migrations.Migrate()
	server.Start()
}
