package controllers

import (
	"encoding/json"
	"net/http"

	"ecommerce-backend/database"
	"ecommerce-backend/models"
	"ecommerce-backend/utils"

	"golang.org/x/crypto/bcrypt"
)

func Register(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	user.Password = string(hashedPassword)

	database.DB.Create(&user)
	json.NewEncoder(w).Encode(user)
}

func Login(w http.ResponseWriter, r *http.Request) {
	var user models.User
	var input models.User
	json.NewDecoder(r.Body).Decode(&input)

	database.DB.Where("email = ?", input.Email).First(&user)
	if user.ID == 0 || bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)) != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	token, _ := utils.GenerateToken(user.ID)
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}
