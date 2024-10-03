package utils

import (
	"net/http"

	"github.com/knid/ilet/internal/database"
	"github.com/knid/ilet/internal/models"
)

func GetUserFromRequest(r *http.Request, db database.Database) (models.User, error) {
	token, err := ExtractTokenFromHeader(r)
	if err != nil {
		return models.User{}, err
	}

	tokenModel, err := db.GetTokenByToken(token)

	if err != nil {
		return models.User{}, err
	}

	user, err := db.GetUserById(tokenModel.User.ID)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}
