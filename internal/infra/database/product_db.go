package database

import (
	"github.com/raul-franca/go-fc-apis/internal/entity"
	"gorm.io/gorm"
)

//	Create(product *entity.Product) error
//	FindAll(page, limit int, sort string) ([]entity.Product, error)
//	FindByID(id string) (*entity.Product, error)
//	Update(product *entity.Product) error
//	Delete(id string) error

type Product struct {
	DB *gorm.DB
}

func NewProduct(db *gorm.DB) *Product {
	return &Product{DB: db}
}

func (p Product) Create(product *entity.Product) error {
	return p.DB.Create(&product).Error
}

func (p Product) FindAll(page, limit int, sort string) ([]entity.Product, error) {
	var products []entity.Product
	var err error

	//regra de ordenação
	if sort != "" && sort != "asc" && sort != "desc" {
		sort = "asc"
	}
	if page != 0 && limit != 0 {
		err = p.DB.Limit(limit).Offset((page - 1) * limit).Order("created_at " + sort).Find(&products).Error
	} else {
		err = p.DB.Order("created_at " + sort).Find(&products).Error
	}
	return products, err
}

func (p Product) FindByID(id string) (*entity.Product, error) {
	var product = entity.Product{}
	err := p.DB.Where("id = ?", id).First(&product).Error
	return &product, err
}