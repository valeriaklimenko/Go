package main

import (
"database/sql"
"encoding/json"
"log"
"net/http"

"lab5/app"
"lab5/database"

_ "github.com/mattn/go-sqlite3"
)

type AddUserRequest struct {
Name string `json:"name"`
}

func writeJSON(w http.ResponseWriter, status int, v any) {
w.Header().Set("Content-Type", "application/json")
w.WriteHeader(status)
json.NewEncoder(w).Encode(v)
}

func addUserHandler(appService *app.App) http.HandlerFunc {
return func(w http.ResponseWriter, r *http.Request) {
if r.Method != http.MethodPost {
http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
return
}

var req AddUserRequest
if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
http.Error(w, "bad json", http.StatusBadRequest)
return
}

if req.Name == "" {
http.Error(w, "name is empty", http.StatusBadRequest)
return
}

_, err := appService.Run(req.Name)
if err != nil {
http.Error(w, err.Error(), http.StatusInternalServerError)
return
}

writeJSON(w, http.StatusCreated, map[string]string{"message": "user created"})
}
}

func getAllUsersHandler(appService *app.App) http.HandlerFunc {
return func(w http.ResponseWriter, r *http.Request) {
if r.Method != http.MethodGet {
http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
return
}

users, err := appService.Run("")
if err != nil {
http.Error(w, "failed to get users", http.StatusInternalServerError)
return
}

writeJSON(w, http.StatusOK, users)
}
}

func deleteAllUsersHandler(appService *app.App) http.HandlerFunc {
return func(w http.ResponseWriter, r *http.Request) {
if r.Method != http.MethodDelete {
http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
return
}

if err := appService.Clear(); err != nil {
http.Error(w, "failed to delete users", http.StatusInternalServerError)
return
}

writeJSON(w, http.StatusOK, map[string]string{"message": "all users deleted"})
}
}

func main() {
conn, err := sql.Open("sqlite3", "test.db")
if err != nil {
log.Fatal(err)
}
defer conn.Close()

db := database.New(conn)
appService := app.New(db)

if err := db.CreateTable(); err != nil {
log.Fatal(err)
}

mux := http.NewServeMux()
mux.HandleFunc("/add", addUserHandler(appService))
mux.HandleFunc("/users", getAllUsersHandler(appService))
mux.HandleFunc("/clear", deleteAllUsersHandler(appService))

log.Println("HTTP сервер запущен на :8080")
log.Println("POST   /add   - добавить пользователя")
log.Println("GET    /users - получить всех пользователей")
log.Println("DELETE /clear - удалить всех пользователей")
log.Fatal(http.ListenAndServe(":8080", mux))
}
