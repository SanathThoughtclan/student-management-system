package middlewares

import (
	//"fmt"
	"net/http"
	"student-management-system/utils"

	"github.com/dgrijalva/jwt-go"
)

func JWTAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			http.Error(w, "Authorization header is required", http.StatusUnauthorized)
			return
		}

		token, err := utils.ParseJWT(tokenString)
		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			http.Error(w, "Invalid token claims", http.StatusUnauthorized)
			return
		}

		username := claims["username"].(string)
		ctx := utils.NewContextWithUserName(r.Context(), username)

		r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	})
}
