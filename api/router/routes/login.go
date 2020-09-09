package routes

import (
	"encoding/json"
	"net/http"

	Jwt "github.com/dennischiu/govuekuber/api/jwt"
	Passwords "github.com/dennischiu/govuekuber/api/models/passwords"
	Users "github.com/dennischiu/govuekuber/api/models/users"
	"xorm.io/xorm"
)

// Login - login route
func Login(db *xorm.Engine) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()

		email := r.FormValue("email")
		password := r.FormValue("password")

		if len(email) == 0 || len(password) == 0 {
			http.Error(w, "An email and password are required", http.StatusBadRequest)
			return
		}

		users, err := Users.Index(db, &Users.User{Email: email})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if len(users) == 0 {
			http.Error(w, "User or password combination not found", http.StatusUnauthorized)
			return
		}

		user := users[0]

		if !Passwords.Compare(user.Password, password) {
			http.Error(w, "User or password combination not found", http.StatusUnauthorized)
			return
		}

		user.Password = ""

		token, err := Jwt.CreateToken(user.ID)

		if err != nil {
			http.Error(w, "Unable to create json token", http.StatusInternalServerError)
			return
		}

		packet, err := json.Marshal(token)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(packet)
	}
}
