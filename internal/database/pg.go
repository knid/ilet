package database

import (
	"database/sql"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

type PostgresDB struct {
	Username string
	Password string
	DBName   string
	Address  string
	SSLMode  string
	db       *sql.DB
}

func (db *PostgresDB) Connect() error {
	connStr := "postgresql://" + db.Username + ":" + db.Password + "@" + db.Address + "/" + db.DBName + "?sslmode=" + db.SSLMode
	pgdb, err := sql.Open("postgres", connStr)
	if err != nil {
		return err
	}

	db.db = pgdb
	return nil
}

func (db *PostgresDB) CheckConnection() error {
	return db.db.Ping()
}

func (db *PostgresDB) MakeMigrations() error {
	driver, err := postgres.WithInstance(db.db, &postgres.Config{})
	m, err := migrate.NewWithDatabaseInstance(
		"file://db/migrations",
		db.DBName,
		driver,
	)
	m.Up()

	return err
}
