package user

import (
	"encoding/json"
	"fmt"
	"go-practice-api/utilities"
	"net/http"
)

type ReqLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var req ReqLogin
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {
		fmt.Println(err)
		utilities.SendError(w, http.StatusBadRequest, "Invalid request body")
		return
	}

	usr, err := h.userRepo.Find(req.Email, req.Password)
	if err != nil {
		utilities.SendError(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	accessToken, err := utilities.CreateJwt(h.cnf.JwtSecretKey, utilities.Payload{
		Sub:       usr.ID,
		FirstName: usr.FirstName,
		LastName:  usr.LastName,
		Email:     usr.Email,
	})
	if err != nil {
		http.Error(w, "Failed to generate token", http.StatusInternalServerError)
		return
	}

	utilities.SendData(w, http.StatusOK, accessToken)
}
