package pojo

type Product struct {
	// --json  gorm -- go orm -- mysql
	ProductId            string `json:"productId" gorm:"column:product_id"`
	ProductName          string `json:"productName" gorm:"column:product_name"`
	ProductIntro         string `json:"productIntro" gorm:"column:product_intro"`
	CategoryId           string `json:"categoryId" gorm:"column:category_id"`
	ProductCoverImg      string `json:"productCoverImg" gorm:"column:product_cover_img"`
	ProductBanner        string `json:"productBanner" gorm:"column:product_banner"`
	OriginalPrice        int    `json:"originalPrice" gorm:"column:original_price"`
	SallingPrice         int    `json:"sallingPrice" gorm:"column:salling_price"`
	StockNum             int    `json:"stockNum" gorm:"column:stock_num"`
	Tag                  string `json:"tag" gorm:"column:tag"`
	SellStatus           int    `json:"sellStatus" gorm:"column:sell_status"`
	CreateUser           string `json:"createUser" gorm:"column:create_user"`
	UpdateUser           string `json:"updateUser" gorm:"column:update_user"`
	ProductDetailContent string `json:"productDetailContent" gorm:"column:produc_detail_content"`
	IsDeleted            bool   `json:"isDeleted" gorm:"column:isDeleted"`
}
