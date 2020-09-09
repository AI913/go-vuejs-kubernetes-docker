package routes

import (
	"encoding/json"
	"net/http"

	"github.com/dennischiu/govuekuber/api/jwt"
	"xorm.io/xorm"
)

// Check - this route verifies authentication
func Check(db *xorm.Engine) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tokenVal := r.Header.Get("X-App-Token")

		user, err := jwt.GetUserFromToken(db, tokenVal)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		packet, err := json.Marshal(user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(packet)
	}
}
