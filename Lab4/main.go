package main

import (
"database/sql"
"log"

"lab4/app"
"lab4/database"

"fyne.io/fyne/v2"
fyneapp "fyne.io/fyne/v2/app"
"fyne.io/fyne/v2/container"
"fyne.io/fyne/v2/widget"
_ "github.com/mattn/go-sqlite3"
)

func main() {
conn, err := sql.Open("sqlite3", "test.db")
if err != nil {
log.Fatal(err)
}
defer conn.Close()

db := database.New(conn)
appLogic := app.New(db)

myApp := fyneapp.New()
window := myApp.NewWindow("Работа с БД")
window.Resize(fyne.NewSize(500, 400))

nameEntry := widget.NewEntry()
nameEntry.SetPlaceHolder("Введите имя")

label := widget.NewLabel("Список пользователей:")

loadUsers := func() {
users, err := appLogic.Run("")
if err != nil {
label.SetText("Ошибка загрузки: " + err.Error())
return
}

if len(users) == 0 {
label.SetText("Список пользователей:\n(пусто)")
return
}

names := "Список пользователей:\n"
for _, name := range users {
names += "• " + name + "\n"
}
label.SetText(names)
}

addBtn := widget.NewButton("Добавить", func() {
if nameEntry.Text != "" {
_, err := appLogic.Run(nameEntry.Text)
if err != nil {
label.SetText("Ошибка: " + err.Error())
return
}
nameEntry.SetText("")
loadUsers()
}
})

refreshBtn := widget.NewButton("Обновить", func() {
loadUsers()
})

clearBtn := widget.NewButton("Очистить всё", func() {
if err := appLogic.Clear(); err != nil {
label.SetText("Ошибка очистки: " + err.Error())
return
}
loadUsers()
})

window.SetContent(container.NewVBox(
widget.NewLabel("Добавить пользователя:"),
nameEntry,
container.NewHBox(addBtn, refreshBtn, clearBtn),
widget.NewSeparator(),
label,
))

loadUsers()
window.ShowAndRun()
}
