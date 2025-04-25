package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/user-management/internal/models"
	"github.com/user-management/internal/repository"
)

type UserHandler struct {
	userRepo repository.UserRepository
}

func NewUserHandler(userRepo repository.UserRepository) *UserHandler {
	return &UserHandler{userRepo: userRepo}
}

// @Summary Получение статуса пользователя
// @Description Получение информации о пользователе и его балансе
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "ID пользователя"
// @Security Bearer
// @Success 200 {object} models.UserStatusResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 401 {object} models.ErrorResponse
// @Failure 404 {object} models.ErrorResponse
// @Router /users/{id}/status [get]
func (h *UserHandler) GetUserStatus(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	user, err := h.userRepo.GetByID(id)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	response := models.UserStatusResponse{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		Points:    user.Points,
		CreatedAt: user.CreatedAt.Format("2006-01-02T15:04:05.999999Z"),
		UpdatedAt: user.UpdatedAt.Format("2006-01-02T15:04:05.999999Z"),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// @Summary Получение таблицы лидеров
// @Description Получение списка пользователей, отсортированного по количеству поинтов
// @Tags users
// @Accept json
// @Produce json
// @Security Bearer
// @Success 200 {array} models.UserLeaderboardResponse
// @Failure 401 {object} models.ErrorResponse
// @Router /users/leaderboard [get]
func (h *UserHandler) GetLeaderboard(w http.ResponseWriter, r *http.Request) {
	users, err := h.userRepo.GetLeaderboard()
	if err != nil {
		http.Error(w, "Error getting leaderboard", http.StatusInternalServerError)
		return
	}

	var response []models.UserLeaderboardResponse
	for _, user := range users {
		response = append(response, models.UserLeaderboardResponse{
			ID:       user.ID,
			Username: user.Username,
			Points:   user.Points,
		})
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// @Summary Выполнение задания
// @Description Отметка задания как выполненного и начисление поинтов
// @Tags tasks
// @Accept json
// @Produce json
// @Param id path int true "ID пользователя"
// @Param task body models.CompleteTaskRequest true "Данные задания"
// @Security Bearer
// @Success 200 {object} models.CompleteTaskResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 401 {object} models.ErrorResponse
// @Failure 404 {object} models.ErrorResponse
// @Router /users/{id}/task/complete [post]
func (h *UserHandler) CompleteTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	var req models.CompleteTaskRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := h.userRepo.CompleteTask(id, req.TaskID); err != nil {
		http.Error(w, "Error completing task", http.StatusInternalServerError)
		return
	}

	response := models.CompleteTaskResponse{
		Success: true,
		Message: "Задание успешно выполнено",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// @Summary Установка реферера
// @Description Установка пользователя, который пригласил текущего пользователя
// @Tags users
// @Accept json
// @Produce json
// @Param id path int true "ID пользователя"
// @Param referrer body models.SetReferrerRequest true "Данные реферера"
// @Security Bearer
// @Success 200 {object} models.SetReferrerResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 401 {object} models.ErrorResponse
// @Failure 404 {object} models.ErrorResponse
// @Router /users/{id}/referrer [post]
func (h *UserHandler) SetReferrer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	var req models.SetReferrerRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := h.userRepo.SetReferrer(id, req.ReferrerID); err != nil {
		http.Error(w, "Error setting referrer", http.StatusInternalServerError)
		return
	}

	response := models.SetReferrerResponse{
		Success: true,
		Message: "Реферер успешно установлен",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
