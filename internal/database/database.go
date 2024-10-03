package database

import "github.com/knid/ilet/internal/models"

type Database interface {
	Connect() error
	CheckConnection() error
	MakeMigrations() error

    GetTokenByUser(models.User) (models.Token, error)
	GetTokenById(int) (models.Token, error)
	GetTokenByToken(string) (models.Token, error)
	CreateToken(models.Token) (models.Token, error)
	DeleteToken(models.Token) error

	GetUserByCredentials(string, string) (models.User, error)
	GetUserById(int) (models.User, error)
	CreateUser(models.User) (models.User, error)
	UpdateUser(models.User) (models.User, error)
	DeleteUser(models.User) error

	GetLinkById(int) (models.Link, error)
	GetLinksByUser(models.User) ([]models.Link, error)
	GetLinkByShortURL(string) (models.Link, error)
	CreateLink(models.Link) (models.Link, error)
	UpdateLink(models.Link) (models.Link, error)
	DeleteLink(models.Link) error
}
