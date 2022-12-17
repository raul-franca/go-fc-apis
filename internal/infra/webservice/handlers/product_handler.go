package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/raul-franca/go-fc-apis/internal/dto"
	"github.com/raul-franca/go-fc-apis/internal/entity"
	"github.com/raul-franca/go-fc-apis/internal/infra/database"
	"net/http"
)

type ProductHandler struct {
	ProductDB database.ProductInterface
}

func NewProductHandler(db database.ProductInterface) *ProductHandler {
	return &ProductHandler{ProductDB: db}
}

func (h ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product dto.CreateProductInput
	err := json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		fmt.Print(err)
		msg := struct {
			Message string `json:"message1"`
		}{
			Message: err.Error(),
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(msg)
		return
	}
	p, err := entity.NewProduct(product.Name, product.Price)
	if err != nil {
		msg := struct {
			Message string `json:"message2"`
		}{
			Message: err.Error(),
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(msg)
		return
	}
	err = h.ProductDB.Create(p)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
