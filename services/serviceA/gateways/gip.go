package gateways

import (
	"encoding/json"
	"net/http"

	"github.com/bookun/glue-sample/services/serviceA/controllers"
)

type GIPResult struct {
	IP       string `json:"ip"`
	Hostname string `json:"hostname"`
	City     string `json:"city"`
	Region   string `json:"region"`
	Country  string `json:"country"`
	Loc      string `json:"loc"`
	Org      string `json:"org"`
	Postal   string `json:"postal"`
	Readme   string `json:"readme"`
}

type GIPGateway struct{}

func NewGIPGateway() controllers.IPGetter {
	return &GIPGateway{}
}

func (g *GIPGateway) GetMyGIP() (string, error) {
	res, err := http.Get("http://ipinfo.io")
	if err != nil {
		return "", err
	}
	var result GIPResult
	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return "", err
	}
	return result.IP, nil
}
