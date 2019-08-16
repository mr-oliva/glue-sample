package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/bookun/glue-sample/di"
	"github.com/bookun/glue-sample/services/serviceA/controllers"
)

type User struct {
	controller *controllers.User
	server     *di.Server
}

func NewUser(u *controllers.User, s *di.Server) *User {
	return &User{controller: u, server: s}
}

func (u *User) HandleTaskGetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := u.controller.GetList()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		u.server.Logger.Errorf("can't get user list: %v", err.Error())
		return
	}
	fmt.Fprintf(w, strings.Join(users, "\n"))
}

func (u *User) HandleTaskGetUserNameById(w http.ResponseWriter, r *http.Request) {
	requestPathParts := strings.Split(r.RequestURI, "/")
	searchIDStr := requestPathParts[len(requestPathParts)-1]
	searchID, err := strconv.Atoi(searchIDStr)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		u.server.Logger.Errorf("can't convert from string to int: %v", err.Error())
		return
	}
	userName, err := u.controller.GetNameById(searchID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		u.server.Logger.Errorf("can't get username by id: %v", err.Error())
		return
	}
	u.server.Logger.Printf("id : %d, name: %s", searchID, userName)
	fmt.Fprintf(w, userName)
}
