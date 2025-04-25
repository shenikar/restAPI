package models

// LoginRequest представляет запрос на аутентификацию
type LoginRequest struct {
	Username string `json:"username" example:"testuser"`
	Password string `json:"password" example:"password123"`
}

// LoginResponse представляет ответ с JWT токеном
type LoginResponse struct {
	Token string `json:"token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."`
}

// UserStatusResponse представляет информацию о пользователе
type UserStatusResponse struct {
	ID        int64  `json:"id" example:"1"`
	Username  string `json:"username" example:"testuser"`
	Email     string `json:"email" example:"test@example.com"`
	Points    int    `json:"points" example:"100"`
	CreatedAt string `json:"created_at" example:"2025-04-02T15:37:13.347084Z"`
	UpdatedAt string `json:"updated_at" example:"2025-04-02T15:37:13.347084Z"`
}

// UserLeaderboardResponse представляет запись в таблице лидеров
type UserLeaderboardResponse struct {
	ID       int64  `json:"id" example:"1"`
	Username string `json:"username" example:"testuser"`
	Points   int    `json:"points" example:"100"`
}

// CompleteTaskRequest представляет запрос на выполнение задания
type CompleteTaskRequest struct {
	TaskID int64 `json:"task_id" example:"1"`
}

// CompleteTaskResponse представляет ответ на выполнение задания
type CompleteTaskResponse struct {
	Success bool   `json:"success" example:"true"`
	Message string `json:"message" example:"Задание успешно выполнено"`
}

// SetReferrerRequest представляет запрос на установку реферера
type SetReferrerRequest struct {
	ReferrerID int64 `json:"referrer_id" example:"2"`
}

// SetReferrerResponse представляет ответ на установку реферера
type SetReferrerResponse struct {
	Success bool   `json:"success" example:"true"`
	Message string `json:"message" example:"Реферер успешно установлен"`
}

// ErrorResponse представляет сообщение об ошибке
type ErrorResponse struct {
	Error string `json:"error" example:"Произошла ошибка"`
}
