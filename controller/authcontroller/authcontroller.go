package authcontroller

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/riz-it/go-jwt-mux/config"
	"github.com/riz-it/go-jwt-mux/helper"
	"github.com/riz-it/go-jwt-mux/model"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Login(w http.ResponseWriter, r *http.Request) {
	var userRequest model.User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&userRequest); err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}
	defer r.Body.Close()

	var user model.User
	if err := model.DB.Where("email =?", userRequest.Email).First(&user).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			response := map[string]string{"message": "User not found"}
			helper.ResponseJSON(w, http.StatusUnauthorized, response)
			return
		default:
			response := map[string]string{"message": "Internal server error"}
			helper.ResponseJSON(w, http.StatusInternalServerError, response)
			return
		}
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userRequest.Password)); err != nil {
		response := map[string]string{"message": "Invalid credentials"}
		helper.ResponseJSON(w, http.StatusUnauthorized, response)
		return
	}

	expTime := time.Now().Add(time.Minute * 1)
	claims := &config.JWTClaim{
		Email: user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "go-jwt-mux",
			ExpiresAt: jwt.NewNumericDate(expTime),
		},
	}

	generatedToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := generatedToken.SignedString(config.JWT_KEY)
	if err != nil {
		response := map[string]string{"message": "Internal server error"}
		helper.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}

	// set to cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    token,
		Path:     "/",
		HttpOnly: true,
	})

	response := map[string]string{"message": "Login success"}
	helper.ResponseJSON(w, http.StatusOK, response)

}

func Register(w http.ResponseWriter, r *http.Request) {
	var userRequest model.User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&userRequest); err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusBadRequest, response)
		return
	}
	defer r.Body.Close()
	hashpassword, _ := bcrypt.GenerateFromPassword([]byte(userRequest.Password), bcrypt.DefaultCost)
	userRequest.Password = string(hashpassword)
	if err := model.DB.Create(&userRequest).Error; err != nil {
		response := map[string]string{"message": err.Error()}
		helper.ResponseJSON(w, http.StatusInternalServerError, response)
		return
	}
	response := map[string]string{"message": "success"}
	helper.ResponseJSON(w, http.StatusOK, response)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		MaxAge:   -1,
	})

	response := map[string]string{"message": "Logout success"}
	helper.ResponseJSON(w, http.StatusOK, response)
}
