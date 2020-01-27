package permissions

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"log"
	"strings"
)

func CheckPermissions(permission string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			user := c.Get("user").(*jwt.Token)
			if user == nil {
				log.Fatal("Error token is not parsed")
			}
			claims := user.Claims.(jwt.MapClaims)

			var permissions_scope_combined []string

			var keysInClaims []string
			for k, _ := range claims {
				keysInClaims = append(keysInClaims, k)
			}

			if contains(keysInClaims, "scope") {
				s := claims["scope"].(string)
				for _, p := range strings.Split(s, " ") {
					permissions_scope_combined = append(permissions_scope_combined, p)
				}
			}
			if contains(keysInClaims, "permissions") {
				for _, v := range claims["permissions"].([]interface{}) {
					s, ok := v.(string)
					if !ok {
						log.Println("Permission was coercible into string: %s", s)
					} else {
						permissions_scope_combined = append(permissions_scope_combined, s)
					}
				}
			}

			if !contains(permissions_scope_combined, permission) {
				return echo.ErrForbidden
			}
			return next(c)
		}
	}
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
