package models

import "time"

type User struct {
	ID        int64     `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	Points    int       `json:"points"`
	Referrer  *int64    `json:"referrer,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Task struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Points      int       `json:"points"`
	CreatedAt   time.Time `json:"created_at"`
}

type CompletedTask struct {
	ID          int64     `json:"id"`
	UserID      int64     `json:"user_id"`
	TaskID      int64     `json:"task_id"`
	CompletedAt time.Time `json:"completed_at"`
}
