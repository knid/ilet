package database

import (
	"database/sql"
	"log"

	"github.com/knid/ilet/internal/models"
)

func (db *PostgresDB) GetLinkById(id int) (models.Link, error) {
	var link models.Link
	row := db.db.QueryRow("SELECT * FROM links WHERE id = $1", link.ID)
	if err := scanIntoLink(row, &link); err != nil {
		return link, err
	}

	return link, nil
}

func (db *PostgresDB) GetLinkByShortURL(short string) (models.Link, error) {
	var link models.Link
	row := db.db.QueryRow("SELECT * FROM links WHERE short = $1", short)
	if err := scanIntoLink(row, &link); err != nil {
		return link, err
	}

	return link, nil
}

func (db *PostgresDB) GetLinksByUser(user models.User) ([]models.Link, error) {
	rows, err := db.db.Query("SELECT * FROM links WHERE user_id = $1", user.ID)
	if err != nil {
		return []models.Link{}, err
	}
	defer rows.Close()

	var links []models.Link

	for rows.Next() {
		var link models.Link
		if err := rows.Scan(&link.ID, &link.User.ID, &link.Short, &link.Long, &link.Active, &link.Visited, &link.CreatedAt, &link.UpdatedAt); err != nil {
			return links, err
		}

		links = append(links, link)
	}

	if err := rows.Err(); err != nil {
		return links, err
	}

	return links, nil
}

func (db *PostgresDB) CreateLink(link models.Link) (models.Link, error) {
	var linkId int
	if err := db.db.QueryRow("INSERT INTO links (user_id, short, long_url, active, visited) VALUES ($1, $2, $3, $4, $5) RETURNING id",
		link.User.ID, link.Short, link.Long, link.Active, link.Visited).Scan(&linkId); err != nil {
		return models.Link{}, err
	}

	log.Println(linkId)

	return db.GetLinkById(linkId)
}

func (db *PostgresDB) UpdateLink(link models.Link) (models.Link, error) {
	_, err := db.db.Exec("UPDATE links SET long_url = $1, active = $2, visited = $3, updated_at = NOW() WHERE id = $4", link.Long, link.Active, link.Visited, link.ID)
	if err != nil {
		return models.Link{}, err
	}

	return link, nil

}

func (db *PostgresDB) DeleteLink(link models.Link) error {
	_, err := db.db.Exec("DELETE FROM links WHERE id = $1", link.ID)
	if err != nil {
		return err
	}

	return nil

}

func scanIntoLink(row *sql.Row, link *models.Link) error {
	return row.Scan(
		&link.ID,
		&link.User.ID,
		&link.Short,
		&link.Long,
		&link.Active,
		&link.Visited,
		&link.CreatedAt,
		&link.UpdatedAt)
}
