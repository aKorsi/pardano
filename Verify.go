package pardano

import (
	"github.com/tiaguinho/gosoap"
	"strings"
)

type VerifyPayment struct {
	Api       string
	Amount    int
	Authority int
}

func (s *VerifyPayment) Verify() error {
	wsdl := "http://pardano.com/p/webservice/?wsdl"
	if strings.EqualFold(s.Api, "test") {
		wsdl = "http://pardano.com/p/webservice-test/?wsdl"
	}
	client, err := gosoap.SoapClient(wsdl)
	if err != nil {
		return err
	}
	err = client.Call("verification", gosoap.Params{
		"api":       s.Api,
		"amount":    s.Amount,
		"authority": s.Authority,
	})
	if err != nil {
		return err
	}

	res := new(struct {
		Return int `xml:"return"`
	})
	err = client.Unmarshal(res)
	if err != nil {
		return err
	}

	return errorResolver(res.Return)
}
