package handlers

import (
	_ "github.com/raul-franca/go-fc-apis/internal/entity"
	"github.com/raul-franca/go-fc-apis/internal/infra/database"
	_ "gorm.io/gorm"
)

type UserHandler struct {
	UserDB database.UserInterface
}

func NewUserHadler(userDB database.UserInterface) *UserHandler {
	return &UserHandler{UserDB: userDB}
}
