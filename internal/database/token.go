package database

import (
	"database/sql"

	"github.com/knid/ilet/internal/models"
)

func (db *PostgresDB) GetTokenById(id int) (models.Token, error) {
	var token models.Token
	row := db.db.QueryRow("SELECT * FROM tokens WHERE id = $1", id)
	if err := scanIntoToken(row, &token); err != nil {
		return token, err
	}

	return token, nil
}

func (db *PostgresDB) GetTokenByToken(t string) (models.Token, error) {
	var token models.Token
	row := db.db.QueryRow("SELECT * FROM tokens WHERE token = $1", t)
	if err := scanIntoToken(row, &token); err != nil {
		return token, err
	}

	return token, nil
}

func (db *PostgresDB) GetTokenByUser(user models.User) (models.Token, error) {
	var token models.Token
	row := db.db.QueryRow("SELECT * FROM tokens WHERE user_id = $1", user.ID)
	if err := scanIntoToken(row, &token); err != nil {
		return token, err
	}

	return token, nil
}

func (db *PostgresDB) CreateToken(token models.Token) (models.Token, error) {
	var tokenId int
	if err := db.db.QueryRow("INSERT INTO tokens (token, user_id) VALUES ($1, $2) RETURNING id", token.Token, token.User.ID).Scan(&tokenId); err != nil {
		return models.Token{}, err
	}

	return db.GetTokenById(tokenId)
}

func (db *PostgresDB) DeleteToken(token models.Token) error {
	_, err := db.db.Exec("DELETE FROM tokens WHERE id = $1", token.ID)
	if err != nil {
		return err
	}

	return nil

}

func scanIntoToken(row *sql.Row, token *models.Token) error {
	return row.Scan(
		&token.ID,
		&token.User.ID,
		&token.Token,
		&token.CreatedAt,
		&token.UpdatedAt,
	)
}
