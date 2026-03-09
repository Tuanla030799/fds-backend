package auth

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	Subject, Email, Name, Role string
	Expiry                     time.Time
}

type JWTManager struct{ secret string }

func NewJWTManager(secret string) *JWTManager { return &JWTManager{secret: secret} }
func (m *JWTManager) Sign(input Claims) (string, error) {
	now := time.Now()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{"sub": input.Subject, "email": input.Email, "name": input.Name, "role": input.Role, "iat": now.Unix(), "exp": input.Expiry.Unix()})
	return token.SignedString([]byte(m.secret))
}
func (m *JWTManager) Parse(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return []byte(m.secret), nil
	})
	if err != nil || !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("invalid token claims")
	}
	return claims, nil
}
