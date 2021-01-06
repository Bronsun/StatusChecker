package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/Bronsun/StatusChecker/helpers"
	"github.com/Bronsun/StatusChecker/interfaces"
	"github.com/Bronsun/StatusChecker/users"
	"github.com/gorilla/mux"
)

type Login struct {
	Login    string
	Password string
}

type Register struct {
	Login    string
	Email    string
	Password string
}

type URL struct {
	Login    string
	Password string
	Link     string
}

func readBody(r *http.Request) []byte {
	body, err := ioutil.ReadAll(r.Body)
	helpers.HandleErr(err)

	return body
}

func apiResponse(call map[string]interface{}, w http.ResponseWriter) {
	if call["message"] == "ok" {
		resp := call
		json.NewEncoder(w).Encode(resp)
	} else {

		resp := interfaces.ErrResponse{Message: "Wrong username or password"}
		json.NewEncoder(w).Encode(resp)
	}
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {

	body := readBody(r)

	var loginData Login
	err := json.Unmarshal(body, &loginData)
	helpers.HandleErr(err)
	login := users.Login(loginData.Login, loginData.Password)

	apiResponse(login, w)
}
func RegisterHandler(w http.ResponseWriter, r *http.Request) {

	body := readBody(r)

	var registerData Register
	err := json.Unmarshal(body, &registerData)
	helpers.HandleErr(err)
	register := users.Register(registerData.Login, registerData.Email, registerData.Password)

	apiResponse(register, w)
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userId := vars["id"]
	auth := r.Header.Get("Authorization")

	user := users.GetUser(userId, auth)
	apiResponse(user, w)
}

func AddLinkHandler(w http.ResponseWriter, r *http.Request) {
	body := readBody(r)

	var addLink URL
	err := json.Unmarshal(body, &addLink)
	helpers.HandleErr(err)
	addlink := users.AddLink(addLink.Login, addLink.Password, addLink.Link)
	apiResponse(addlink, w)
}
