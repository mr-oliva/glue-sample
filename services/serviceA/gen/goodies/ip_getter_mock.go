package goodies

type IPGetterMock struct {
}

func (i *IPGetterMock) GetMyGIP() (string, error) {
	return "0.0.0.0", nil
}
