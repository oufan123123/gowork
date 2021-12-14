package pojo

type Order struct {
	OrderId     string `json:"orderId" gorm:"column:order_id"`
	UserId      string `json:"userId" gorm:"column:user_id"`
	UserName    string `json:"userName" gorm:"column:user_name"`
	Mobile      string `json:"mobile" gorm:"column:mobile"`
	TotalPrice  int64  `json:"totalPrice" gorm:"column:total_price"`
	PayStatus   int    `json:"payStatus" gorm:"column:pay_status"`
	PayType     int    `json:"payType" gorm:"column:pay_type"`
	PayTime     string `json:"payTime" gorm:"column:pay_time"`
	OrderStatus int    `json:"orderStatus" gorm:"column:order_status"`
	ExtraInfo   string `json:"extraInfo" gorm:"column:extra_info"`
	UserAddress string `json:"userAddress" gorm:"column:userAddress"`
	IsDeleted   bool   `json:"isDeleted" gorm:"column:isDeleted"`
}
