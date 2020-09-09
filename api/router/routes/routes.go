package routes

import (
	"log"
	"net/http"

	"xorm.io/xorm"
)

// Middleware - main middleware for the app
func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Hit main middleware")
		next.ServeHTTP(w, r)
	})
}

// Get routes
func GetRoutes(db *xorm.Engine) Routes {
	return Routes{
		Route{
			Name:        "HealthCheck",
			Method:      "GET",
			Pattern:     "/health",
			HandlerFunc: Health(db),
		},
		Route{
			Name:        "Login",
			Method:      "POST",
			Pattern:     "/auth/login",
			HandlerFunc: Login(db),
		},
		Route{
			Name:        "Check",
			Method:      "POST",
			Pattern:     "/auth/check",
			HandlerFunc: Check(db),
		},
	}
}
