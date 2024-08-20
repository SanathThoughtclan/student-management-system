package utils

import (
	"context"
	"errors"

	"student-management-system/config"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type contextKey string

const (
	UserIDKey   contextKey = "user_id"
	UsernameKey contextKey = "username"
)

var jwtSecret []byte

func init() {
	cfg := config.LoadConfig()
	jwtSecret = []byte(cfg.JWT.SecretKey)
}

func GenerateJWT(userID, username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":  userID,
		"username": username,
		"exp":      time.Now().Add(time.Hour * 72).Unix(),
	})

	return token.SignedString(jwtSecret)
}

func ParseJWT(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return jwtSecret, nil
	})
}

func NewContextWithUserID(ctx context.Context, userID string) context.Context {

	newCtx := context.WithValue(ctx, UserIDKey, userID)

	return newCtx

}

func GetUserIDFromContext(ctx context.Context) (string, bool) {
	userID, ok := ctx.Value(UserIDKey).(string)
	return userID, ok
}

func NewContextWithUserName(ctx context.Context, username string) context.Context {
	newCtx := context.WithValue(ctx, UsernameKey, username)

	return newCtx
}

func GetUsernameFromContext(ctx context.Context) (string, bool) {
	username, ok := ctx.Value(UsernameKey).(string)

	return username, ok
}
