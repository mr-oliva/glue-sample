package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/bookun/glue-sample/services/serviceA/controllers"
)

type User struct {
	controller *controllers.User
}

func NewUser(u *controllers.User) *User {
	return &User{u}
}

func (u *User) HandleTaskGetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := u.controller.GetList()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, strings.Join(users, ","))
}

func (u *User) HandleTaskGetUserNameById(w http.ResponseWriter, r *http.Request) {
	requestPathParts := strings.Split(r.RequestURI, "/")
	searchIDStr := requestPathParts[len(requestPathParts)-1]
	searchID, err := strconv.Atoi(searchIDStr)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	userName, err := u.controller.GetNameById(searchID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, userName)
}
