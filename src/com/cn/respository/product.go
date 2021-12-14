package respository

import (
	"modules/pojo"
	"modules/query"
	"modules/utils"

	"github.com/jinzhu/gorm"
)

type ProductResposity struct {
	db *gorm.DB
}
type ProductRespoInterface interface {
	List(req *query.ListQuery) (products []pojo.Product, err error)
	GetTotal(req *query.ListQuery) (total int64, err error)
	Get(product pojo.Product) (*pojo.Product, error)
	Exist(product pojo.Product) *pojo.Product
	ExistByID(id string) *pojo.Product
	Add(product pojo.Product) (*pojo.Product, error)
	Edit(product pojo.Product) (bool, error)
	Delete(product pojo.Product) (bool, error)
}

func (repo *ProductResposity) List(req *query.ListQuery) (products []pojo.Product, err error) {
	db := repo.db
	if req.Where != "" {
		db = db.Where(req.Where)
	}
	limit, offset := utils.Page(int64(req.PageSize), int64(req.Page))

	err = db.Order("id desc").Limit(int(limit)).Offset(int(offset)).Find(&products).Error
	if err != nil {
		return products, err
	}
	return products, nil
}

func (repo *ProductResposity) GetTotal(req *query.ListQuery) (total int64, err error) {
	db := repo.db
	if req.Where != "" {
		db = db.Where(req.Where)
	}
	var products []pojo.Product
	if err := db.Find(&products).Count(&total).Error; err != nil {
		return total, err
	}
	return total, nil
}

func (repo *ProductResposity) Get(product pojo.Product) (*pojo.Product, error) {
	db := repo.db
	if err := db.Where(&product).First(&product).Error; err != nil {
		return &product, err
	}
	return &product, nil
}

func (repo *ProductResposity) Exist(product pojo.Product) *pojo.Product {
	db := repo.db
	if err := db.Where("category_id=?", product.CategoryId).First(&product).Error; err != nil {
		return &product
	}
	return &product
}

func (repo *ProductResposity) ExistByID(id string) *pojo.Product {
	db := repo.db
	var product pojo.Product
	if err := db.Where("category_id = ?", id).First(&product).Error; err != nil {
		return &product
	}
	return &product
}

func (repo *ProductResposity) Add(product pojo.Product) (*pojo.Product, error) {
	db := repo.db
	// check exist
	if ex := repo.Exist(product); ex != nil {
		return &product, nil
	}
	if err := db.Create(&product).Error; err != nil {
		return &product, err
	}
	return &product, nil

}

func (repo *ProductResposity) Edit(product pojo.Product) (bool, error) {
	db := repo.db
	err := db.Model(&product).Updates(map[string]interface{}{
		"product_Name": product.ProductName, "product_intro": product.ProductIntro,
		"SallingPrice": product.SallingPrice,
	}).Error
	if err != nil {
		return false, err
	}
	return true, err
}

func (repo *ProductResposity) Delete(product pojo.Product) (bool, error) {
	db := repo.db
	if err := db.Delete(&product).Error; err != nil {
		return false, err
	}
	return true, nil
}
