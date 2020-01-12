package pardano

import "errors"

func errorResolver(e int) error {
	switch e {
	case -1:
		return errors.New("api نامعتبر است")
	case -2:
		return errors.New("مبلغ از کف تعریف شده کمتر است")
	case -3:
		return errors.New("مبلغ از سقف تعریف شده بیشتر است")
	case -4:
		return errors.New("مبلغ نامعتبر است")
	case -6:
		return errors.New("درگاه غیرفعال است")
	case -7:
		return errors.New("آی پی شما مسدود است")
	case -9:
		return errors.New("آدرس کال بک خالی است")
	case -10:
		return errors.New("چنین تراکنشی یافت نشد")
	case -11:
		return errors.New("تراکنش انجام نشده")
	case -12:
		return errors.New("تراکنش انجام شده اما مبلغ نادرست است")
	case -13:
		return errors.New("این تراکنش قبلا تایید شده است")
	case -21:
		return errors.New("IP سرور با IP تعریف شده برای درگاه مغایر است")
	}
	return nil
}
