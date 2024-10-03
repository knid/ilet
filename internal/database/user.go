package database

import (
	"database/sql"

	"github.com/knid/ilet/internal/models"
)

func (db *PostgresDB) GetUserById(id int) (models.User, error) {
	var user models.User
	row := db.db.QueryRow("SELECT * FROM users WHERE id = $1", id)
    if err := scanIntoUser(row, &user); err != nil {
        return user, err
    }

	return user, nil
}

func (db *PostgresDB) GetUserByCredentials(username, password string) (models.User, error) {
	var user models.User
	row := db.db.QueryRow("SELECT * FROM users WHERE username = $1 AND password = $2", username, password)
    if err := scanIntoUser(row, &user); err != nil {
        return user, err
    }

	return user, nil
}

func (db *PostgresDB) CreateUser(user models.User) (models.User, error) {
	var userId int
	if err := db.db.QueryRow("INSERT INTO users (username, password) VALUES ($1, $2) RETURNING id", user.Username, user.Password).Scan(&userId); err != nil {
		return models.User{}, err
	}

	return db.GetUserById(userId)
}

func (db *PostgresDB) UpdateUser(user models.User) (models.User, error) {
	_, err := db.db.Exec("UPDATE users SET username = $1, password = $2, updated_at = NOW() WHERE id = $3", user.Username, user.Password, user.ID)
	if err != nil {
		return models.User{}, err
	}

	return user, nil

}

func (db *PostgresDB) DeleteUser(user models.User) error {
	_, err := db.db.Exec("DELETE FROM users WHERE id = $1", user.ID)
	if err != nil {
		return err
	}

	return nil

}

func scanIntoUser(row *sql.Row, user *models.User) error {
    return row.Scan(
        &user.ID,
        &user.Username,
        &user.Password,
        &user.CreatedAt,
        &user.UpdatedAt,
    )
} 
