package handlers

import (
	"fmt"
	"net/http"

	"github.com/bookun/glue-sample/di"
	"github.com/bookun/glue-sample/services/serviceA/controllers"
)

type IP struct {
	controller *controllers.IPGateway
	server     *di.Server
}

func NewIP(i *controllers.IPGateway, s *di.Server) *IP {
	return &IP{controller: i, server: s}
}

func (i *IP) HandleTaskGetMyGIP(w http.ResponseWriter, r *http.Request) {
	ip, err := i.controller.GetMYGIP()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		i.server.Logger.Errorf("can't get my GIP: %v", err.Error())
		return
	}
	fmt.Fprintf(w, ip)
}
