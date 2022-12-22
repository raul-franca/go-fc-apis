package handlers

import (
	"encoding/json"
	"github.com/raul-franca/go-fc-apis/internal/dto"
	"github.com/raul-franca/go-fc-apis/internal/entity"
	_ "github.com/raul-franca/go-fc-apis/internal/entity"
	"github.com/raul-franca/go-fc-apis/internal/infra/database"
	_ "gorm.io/gorm"
	"net/http"
)

type UserHandler struct {
	UserDB database.UserInterface
}

func NewUserHadler(userDB database.UserInterface) *UserHandler {
	return &UserHandler{UserDB: userDB}
}

type Error struct {
	Message string `json:"message"`
}

func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	var user dto.CreateUserInput
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	u, err := entity.NewUser(user.Name, user.Email, user.Password)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}
	err = h.UserDB.Create(u)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
