package enum

type PayStatus int

const (
	UnPay PayStatus = 0
	Payed PayStatus = 1
)

func (p PayStatus) String() string {
	switch p {
	case UnPay:
		return "pay not complete"
	case Payed:
		return "pay complete"
	default:
		return "unknown"
	}
}
