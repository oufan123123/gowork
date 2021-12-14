package enum

type SellStatus int

const (
	Selling  SellStatus = 0
	StopSell SellStatus = 1
)

func (s SellStatus) String() string {
	switch s {
	case Selling:
		return "selling"
	case StopSell:
		return "stop sell"
	default:
		return "unknown"
	}
}
