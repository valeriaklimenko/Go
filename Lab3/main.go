package main

import (
	"database/sql"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, _ := sql.Open("sqlite3", "test.db")
	defer db.Close()

	db.Exec(`CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY, name TEXT)`)

	myApp := app.New()
	window := myApp.NewWindow("Работа с БД")
	window.Resize(fyne.NewSize(400, 200))

	nameEntry := widget.NewEntry()
	nameEntry.SetPlaceHolder("Введите имя")

	label := widget.NewLabel("Список пользователей:")

	loadUsers := func() {
		rows, _ := db.Query("SELECT name FROM users")
		defer rows.Close()

		names := "Список пользователей:\n"
		for rows.Next() {
			var name string
			rows.Scan(&name)
			names += "• " + name + "\n"
		}
		label.SetText(names)
	}

	addBtn := widget.NewButton("Добавить", func() {
		if nameEntry.Text != "" {
			db.Exec("INSERT INTO users (name) VALUES (?)", nameEntry.Text)
			nameEntry.SetText("")
			loadUsers()
		}
	})

	loadUsers()

	refreshBtn := widget.NewButton("Обновить", func() {
		loadUsers()
	})

	window.SetContent(container.NewVBox(
		widget.NewLabel("Добавить пользователя:"),
		nameEntry,
		container.NewHBox(addBtn, refreshBtn),
		widget.NewSeparator(),
		label,
	))

	window.ShowAndRun()
}

