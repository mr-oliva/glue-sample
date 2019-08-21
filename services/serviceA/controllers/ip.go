package controllers

type IPGetter interface {
	GetMyGIP() (string, error)
}

type IPGateway struct {
	getter IPGetter
}

func NewIPGateway(g IPGetter) *IPGateway {
	return &IPGateway{getter: g}
}

func (i *IPGateway) GetMYGIP() (string, error) {
	return i.getter.GetMyGIP()
}
