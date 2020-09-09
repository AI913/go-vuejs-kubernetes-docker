package v1

import (
	"log"
	"net/http"

	Jwt "github.com/dennischiu/govuekuber/api/jwt"
	Routes "github.com/dennischiu/govuekuber/api/router/routes"

	"xorm.io/xorm"
)

// Middleware - for v1 routes
func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("X-App-Token")

		if _, err := Jwt.ParseToken(token); err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		log.Println("Athenticated")
		next.ServeHTTP(w, r)
	})
}

// GetRoutes - get v1(authenticated) routes
func GetRoutes(db *xorm.Engine) (SubRoute map[string]Routes.SubRoutePackage) {
	SubRoute = map[string]Routes.SubRoutePackage{
		"/v1": {
			Routes: Routes.Routes{
				Routes.Route{
					Name:        "V1HealthRoute",
					Method:      "GET",
					Pattern:     "/health",
					HandlerFunc: Health(db),
				},
			},
			Middleware: Middleware,
		},
	}
	return SubRoute
}
