package users

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/Bronsun/StatusChecker/helpers"
	"github.com/Bronsun/StatusChecker/interfaces"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

/// JWT Tokens //////
func prepareToken(user *interfaces.User) string {
	tokenContent := jwt.MapClaims{
		"user_id": user.ID,
		"expiry":  time.Now().Add(time.Minute ^ 60).Unix(),
	}
	jwtToken := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tokenContent)
	token, err := jwtToken.SignedString([]byte("TokenPassword"))
	helpers.HandleErr(err)

	return token

}

//// JWT Tokens and response /////
func prepareResponse(user *interfaces.User, userLinks []interfaces.ResponseLink, withToken bool) map[string]interface{} {
	responseUser := &interfaces.ResponseUser{
		ID:    user.ID,
		Login: user.Login,
		Email: user.Email,
		Links: userLinks,
	}
	//var token = prepareToken(user)
	var response = map[string]interface{}{"message": "ok"}
	if withToken {
		var token = prepareToken(user)
		response["jwt"] = token
	}
	response["data"] = responseUser
	return response
}

//// Get Status of the website ////
func getStatus(link string) string {
	resp, err := http.Get(link)

	if err != nil {
		log.Fatal(err)
	}
	status := strconv.Itoa(resp.StatusCode)
	return status

}

//// Login function /////

func Login(login string, pass string) map[string]interface{} {

	valid := helpers.Validation(
		[]interfaces.Validation{
			{Value: login, Valid: "login"},
			{Value: pass, Valid: "password"},
		})
	if valid {
		db := helpers.ConnectDB()
		user := &interfaces.User{}
		if db.Where("login = ?", login).First(&user).RecordNotFound() {
			return map[string]interface{}{"message": "User not found"}
		}

		passErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(pass))

		if passErr == bcrypt.ErrMismatchedHashAndPassword && passErr != nil {
			return map[string]interface{}{"message": "Wrong password"}
		}

		userLinks := []interfaces.ResponseLink{}
		db.Table("user_links").Select("id,link,status,time").Where("user_id=?", user.ID).Scan(&userLinks)

		defer db.Close()

		var response = prepareResponse(user, userLinks, true)
		return response
	} else {
		return map[string]interface{}{"message": "not valid values"}
	}

}

/// Register function //////

func Register(login string, email string, pass string) map[string]interface{} {
	valid := helpers.Validation(
		[]interfaces.Validation{
			{Value: login, Valid: "login"},
			{Value: email, Valid: "email"},
			{Value: pass, Valid: "password"},
		})
	if valid {
		db := helpers.ConnectDB()
		generatedPassword := helpers.HashAndSalt([]byte(pass))
		user := &interfaces.User{Login: login, Email: email, Password: generatedPassword}
		db.Create(&user)

		now := time.Now()
		status := getStatus("https://google.pl")
		userlink := &interfaces.UserLink{Link: "https://google.pl", Status: status, Time: now, UserId: user.ID}
		db.Create(&userlink)

		defer db.Close()

		userlinks := []interfaces.ResponseLink{}
		respLink := interfaces.ResponseLink{ID: userlink.ID, Link: userlink.Link, Status: userlink.Status, Time: userlink.Time}
		userlinks = append(userlinks, respLink)
		var response = prepareResponse(user, userlinks, true)

		return response

	} else {
		return map[string]interface{}{"message": "not valid values"}
	}
}

/// Get User data ////
func GetUser(id string, jwt string) map[string]interface{} {
	isValid := helpers.ValidateToken(id, jwt)

	if isValid {
		db := helpers.ConnectDB()

		user := &interfaces.User{}
		if db.Where("id = ?", id).First(&user).RecordNotFound() {
			return map[string]interface{}{"message": "User not found"}
		}

		userLinks := []interfaces.ResponseLink{}
		db.Table("user_links").Select("id,link,status,time").Where("user_id=?", user.ID).Scan(&userLinks)

		defer db.Close()
		var response = prepareResponse(user, userLinks, false)
		return response

	} else {
		return map[string]interface{}{"message": "Not valid token"}
	}

}

//// Add link to database //// - it will be change to get JWT token
func AddLink(login string, pass string, link string) map[string]interface{} {
	valid := helpers.Validation(
		[]interfaces.Validation{
			{Value: login, Valid: "login"},
			{Value: pass, Valid: "password"},
		})
	if valid {
		db := helpers.ConnectDB()
		user := &interfaces.User{}
		if db.Where("login = ?", login).First(&user).RecordNotFound() {
			return map[string]interface{}{"message": "User not found"}
		}

		passErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(pass))

		if passErr == bcrypt.ErrMismatchedHashAndPassword && passErr != nil {
			return map[string]interface{}{"message": "Wrong password"}
		}
		userLinks := &interfaces.UserLink{}
		if db.Where("link = ?", link).First(&userLinks).RecordNotFound() {
			now := time.Now()
			status := getStatus(link)
			userlink := &interfaces.UserLink{Link: link, Status: status, Time: now, UserId: user.ID}
			db.Create(&userlink)

			defer db.Close()

			userlinks := []interfaces.ResponseLink{}
			respLink := interfaces.ResponseLink{ID: userlink.ID, Link: userlink.Link, Status: userlink.Status, Time: userlink.Time}
			userlinks = append(userlinks, respLink)
			var response = prepareResponse(user, userlinks, true)

			return response

		} else {
			return map[string]interface{}{"message": "URL exists - check your status using /checkStatus"}
		}

	} else {
		return map[string]interface{}{"message": "not valid values"}
	}

}

/*
func CheckStatus() {
	db := helpers.ConnectDB()
	userlink := &interfaces.UserLink{}

	links := db.Select("link").Find(&userlink)

}
*/
