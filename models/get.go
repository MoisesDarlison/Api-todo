package models

import "postgres/api/db"

func Get(id int64) (todo Todo, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	sql := `SELECT * FROM todos WHERE id=$1`

	row := conn.QueryRow(sql, id)
	err = row.Scan(&todo.id, &todo.Title, &todo.Description, &todo.Done)
	return
}
