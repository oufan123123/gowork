package enum

type PayType int

const (
	bankPay PayType = 0
	wxPay   PayType = 1
	aliPay  PayType = 2
	tikPay  PayType = 3
)

func (p PayType) String() string {
	switch p {
	case bankPay:
		return "bank pay"
	case wxPay:
		return "wexin pay"
	case aliPay:
		return "alibaba pay"
	case tikPay:
		return "tiktok pay"
	default:
		return "unknown"
	}
}
