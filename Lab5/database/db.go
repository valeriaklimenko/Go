package database

import "database/sql"

type SQLiteDB struct {
	conn *sql.DB
}

func New(conn *sql.DB) *SQLiteDB {
	return &SQLiteDB{conn: conn}
}

func (d *SQLiteDB) CreateTable() error {
	_, err := d.conn.Exec(`CREATE TABLE IF NOT EXISTS users(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL
	)`)
	return err
}

func (d *SQLiteDB) Insert(name string) error {
	_, err := d.conn.Exec("INSERT INTO users (name) VALUES (?)", name)
	return err
}

func (d *SQLiteDB) GetAll() ([]string, error) {
	rows, err := d.conn.Query("SELECT name FROM users ORDER BY id DESC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var names []string
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			return nil, err
		}
		names = append(names, name)
	}
	return names, nil
}

func (d *SQLiteDB) DeleteAll() error {
	_, err := d.conn.Exec("DELETE FROM users")
	return err
}
