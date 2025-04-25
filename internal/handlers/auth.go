package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/user-management/internal/models"
	"github.com/user-management/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type AuthHandler struct {
	jwtSecret string
	userRepo  repository.UserRepository
}

func NewAuthHandler(jwtSecret string, userRepo repository.UserRepository) *AuthHandler {
	return &AuthHandler{
		jwtSecret: jwtSecret,
		userRepo:  userRepo,
	}
}

// @Summary Аутентификация пользователя
// @Description Вход в систему и получение JWT токена
// @Tags auth
// @Accept json
// @Produce json
// @Param credentials body models.LoginRequest true "Данные для входа"
// @Success 200 {object} models.LoginResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 401 {object} models.ErrorResponse
// @Router /auth/login [post]
func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req models.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Printf("Error decoding request body: %v", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	log.Printf("Login attempt for user: %s", req.Username)

	// Получаем пользователя из базы данных
	user, err := h.userRepo.GetByUsername(req.Username)
	if err != nil {
		log.Printf("Error getting user from database: %v", err)
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	log.Printf("Found user: %s with ID: %d", user.Username, user.ID)

	// Проверяем пароль
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		log.Printf("Password comparison failed: %v", err)
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	log.Printf("Password verified successfully for user: %s", user.Username)

	// Создаем JWT токен
	expirationTime := time.Now().Add(time.Hour * 24)
	claims := jwt.MapClaims{
		"user_id": user.ID,
		"exp":     expirationTime.Unix(),
		"iat":     time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	log.Printf("Created token with claims: %+v", claims)

	tokenString, err := token.SignedString([]byte(h.jwtSecret))
	if err != nil {
		log.Printf("Error generating token: %v", err)
		http.Error(w, "Error generating token", http.StatusInternalServerError)
		return
	}

	log.Printf("Token generated successfully for user: %s, token length: %d", user.Username, len(tokenString))

	response := models.LoginResponse{
		Token: tokenString,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
