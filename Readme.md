<div dir="rtl">

# کتابخانه اتصال به سامانه پرداخت [پردانو](https://pardano.com/) 

روش کار در حالت تست:

<div dir="ltr">

```go

payment := pardano.RequestPayment{
							Api:         "test",
							OrderId:     orderId,
							Amount:      amount,
							CallBackUrl: "CallBackUrl",
							Description: "Description",
						}

paymentUrl, err := payment.Request()

// verify

payment := pardano.VerifyPayment{
						Api:       "test",
						Amount:    amount,
						Authority: authority,
					}

					err = payment.Verify()

```

</div>

در صورت خطا یا مشکلی حتما اعلام بفرمایید

<div dir="ltr">

`github.com/tiaguinho/gosoap`

**{SOAP package for Go}**

</div>

</div>

