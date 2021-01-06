package migrations

import (
	"github.com/Bronsun/StatusChecker/interfaces"

	"github.com/Bronsun/StatusChecker/helpers"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

/*
func createAccounts() {
	db := helpers.ConnectDB()

	users := &[2]interfaces.User{
		{Login: "Mateusz", Email: "mateusz.broncel1998@gmail.com"},
		{Login: "Zuza", Email: "zuza.twardowska@gmail.com"},
	}
	for i := 0; i < len(users); i++ {
		generatedPassword := helpers.HashAndSalt([]byte(users[i].Login))
		user := &interfaces.User{Login: users[i].Login, Email: users[i].Email, Password: generatedPassword}
		db.Create(&user)
	}
	defer db.Close()
}
*/
func Migrate() {
	User := &interfaces.User{}
	UserLink := &interfaces.UserLink{}
	db := helpers.ConnectDB()
	db.AutoMigrate(&User, &UserLink)
	defer db.Close()

	//createAccounts()
}
