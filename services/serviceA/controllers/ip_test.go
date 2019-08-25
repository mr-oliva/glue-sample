package controllers_test

import (
	"testing"

	"github.com/bookun/glue-sample/services/serviceA/controllers"
	"github.com/bookun/glue-sample/services/serviceA/gen/goodies"
)

func TestGetMYGIP(t *testing.T) {
	mock := &goodies.IPGetterMock{}
	controller := controllers.NewIPGateway(mock)
	ip, err := controller.GetMYGIP()
	if err != nil {
		t.Errorf("occur error")
	}
	if ip != "0.0.0.0" {
		t.Errorf("want 0.0.0.0, got %s", ip)
	}
}
