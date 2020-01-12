package pardano

import (
	"github.com/tiaguinho/gosoap"
	"strconv"
	"strings"
)

type RequestPayment struct {
	Api         string
	Amount      int
	CallBackUrl string
	OrderId     int64
	Description string
}

func (s *RequestPayment) Request() (string, error) {
	wsdl := "http://pardano.com/p/webservice/?wsdl"
	if strings.EqualFold(s.Api, "test") {
		wsdl = "http://pardano.com/p/webservice-test/?wsdl"
	}
	client, err := gosoap.SoapClient(wsdl)
	if err != nil {
		return "", err
	}
	err = client.Call("requestpayment", gosoap.Params{
		"api":         s.Api,
		"amount":      s.Amount,
		"callbackurl": s.CallBackUrl,
		"orderid":     s.OrderId,
		"description": s.Description,
	})
	if err != nil {
		return "", err
	}

	res := new(struct {
		Return int `xml:"return"`
	})
	err = client.Unmarshal(res)
	if err != nil {
		return "", err
	}

	err = errorResolver(res.Return)

	return "http://pardano.com/p/payment/" + strconv.Itoa(res.Return), err
}
