package routes

import (
	"net/http"

	"xorm.io/xorm"
)

// Health - health route
func Health(db *xorm.Engine) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Server is healthy!"))
	}
}
