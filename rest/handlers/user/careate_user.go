package user

import (
	"encoding/json"
	"fmt"
	"go-practice-api/domain"
	"go-practice-api/utilities"
	"net/http"
)

type ReqCreateUser struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	IsShopOwner bool   `json:"is_shop_owner"`
}

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var req ReqCreateUser
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {
		fmt.Println(err)
		utilities.SendError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	usr, err := h.svc.Create(domain.User{
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		Email:       req.Email,
		Password:    req.Password,
		IsShopOwner: req.IsShopOwner,
	})
	if err != nil {
		utilities.SendData(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}
	utilities.SendData(w, http.StatusCreated, usr)
}
