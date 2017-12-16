package repositories

import (
	"database/sql"

	"github.com/rymccue/golang-gin-todo-list-api/models"
)

func GetItems(db *sql.DB, all bool) ([]*models.Item, error) {
	var rows *sql.Rows
	var err error
	query := `
		select
			id,
			title,
			description,
			completed
		from
			items
	`
	if !all {
		query += "where completed = $1"
		rows, err = db.Query(query, all)
	} else {
		rows, err = db.Query(query)
	}

	if err != nil {
		return nil, err
	}
	items := make([]*models.Item, 0)

	for rows.Next() {
		var item models.Item
		err = rows.Scan(&item.ID, &item.Title, &item.Description, &item.Completed)
		if err != nil {
			return nil, err
		}
		items = append(items, &item)
	}
	return items, err
}

func CreateItem(db *sql.DB, title, description string) (int, error) {
	const query = `
		insert into items (
			title,
			description
		) values (
			$1,
			$2
		) returning id
	`
	var id int
	err := db.QueryRow(query, title, description).Scan(&id)
	return id, err
}

func UpdateItem(db *sql.DB, id int, title, description string, completed bool) error {
	const query = `
		update items set
			title = $1,
			description = $2,
			completed = $3
		where id = $4
	`
	_, err := db.Exec(query, title, description, completed, id)
	return err
}

func GetItem(db *sql.DB, id int) (*models.Item, error) {
	const query = `
		select
			id,
			title,
			description,
			completed
		from
			items
		where
			id = $1
	`
	var item models.Item
	err := db.QueryRow(query, id).Scan(&item.ID, &item.Title, &item.Description, &item.Completed)
	return &item, err
}

func DeleteItem(db *sql.DB, id int) error {
	const query = `delete from items where id = $1`
	_, err := db.Exec(query, id)
	return err
}
