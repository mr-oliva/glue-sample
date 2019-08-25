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
	logger     *di.Logger
}

func NewUser(u *controllers.User, l *di.Logger) *User {
	return &User{controller: u, logger: l}
}

func (u *User) HandleTaskGetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := u.controller.GetList()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		u.logger.Errorf("can't get user list: %v", err.Error())
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
		u.logger.Errorf("can't convert from string to int: %v", err.Error())
		return
	}
	userName, err := u.controller.GetNameById(searchID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		u.logger.Errorf("can't get username by id: %v", err.Error())
		return
	}
	u.logger.Printf("id : %d, name: %s", searchID, userName)
	fmt.Fprintf(w, userName)
}
