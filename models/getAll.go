package models

import "postgres/api/db"

func GetAll() (todos []Todo, err error) {
	conn, err := db.OpenConnection()
	if err != nil {
		return
	}
	defer conn.Close()

	sql := `SELECT * FROM todos`

	rows, err := conn.Query(sql)
	if err != nil {
		return
	}

	for rows.Next() {
		var todo Todo
		err = rows.Scan(&todo.id, &todo.Title, &todo.Description, &todo.Done)
		if err != nil {
			continue
		}
		todos = append(todos, todo)
	}
	return
}
