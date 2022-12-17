package handlers

import "github.com/raul-franca/go-fc-apis/internal/infra/database"

type ProductHandler struct {
	ProductDB database.ProductInterface
}

func NewProductHandler(db database.ProductInterface) *ProductHandler {
	return &ProductHandler{ProductDB: db}
}
