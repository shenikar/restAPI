package repository

import (
	"database/sql"
	"time"
)

type User struct {
	ID        int64
	Username  string
	Email     string
	Password  string
	Points    int
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserRepository interface {
	GetByID(id int64) (*User, error)
	GetByUsername(username string) (*User, error)
	GetLeaderboard() ([]*User, error)
	CompleteTask(userID int64, taskID int64) error
	SetReferrer(userID int64, referrerID int64) error
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) GetByUsername(username string) (*User, error) {
	var user User
	err := r.db.QueryRow(`
		SELECT id, username, email, password, points, created_at, updated_at
		FROM users
		WHERE username = $1
	`, username).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.Password,
		&user.Points,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) GetByID(id int64) (*User, error) {
	var user User
	err := r.db.QueryRow(`
		SELECT id, username, email, password, points, created_at, updated_at
		FROM users
		WHERE id = $1
	`, id).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.Password,
		&user.Points,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) GetLeaderboard() ([]*User, error) {
	rows, err := r.db.Query(`
		SELECT id, username, email, password, points, created_at, updated_at
		FROM users
		ORDER BY points DESC
		LIMIT 10
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*User
	for rows.Next() {
		var user User
		err := rows.Scan(
			&user.ID,
			&user.Username,
			&user.Email,
			&user.Password,
			&user.Points,
			&user.CreatedAt,
			&user.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}
	return users, nil
}

func (r *userRepository) CompleteTask(userID int64, taskID int64) error {
	_, err := r.db.Exec(`
		UPDATE users
		SET points = points + 100,
			updated_at = NOW()
		WHERE id = $1
	`, userID)
	return err
}

func (r *userRepository) SetReferrer(userID int64, referrerID int64) error {
	_, err := r.db.Exec(`
		UPDATE users
		SET referrer = $2,
			updated_at = NOW()
		WHERE id = $1
	`, userID, referrerID)
	return err
}
