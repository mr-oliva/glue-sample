package handlers

import (
	"fmt"
	"net/http"

	"github.com/bookun/glue-sample/di"
	"github.com/bookun/glue-sample/services/serviceA/controllers"
)

type IP struct {
	controller *controllers.IPGateway
	logger     *di.Logger
}

func NewIP(i *controllers.IPGateway, s *di.Logger) *IP {
	return &IP{controller: i, logger: s}
}

func (i *IP) HandleTaskGetMyGIP(w http.ResponseWriter, r *http.Request) {
	ip, err := i.controller.GetMYGIP()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		i.logger.Errorf("can't get my GIP: %v", err.Error())
		return
	}
	fmt.Fprintf(w, ip)
}
