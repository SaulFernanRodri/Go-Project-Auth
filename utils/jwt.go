package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var JwtKey = []byte("your_secret_key") // Define una clave secreta segura

// Claims personalizados para almacenar el username del usuario
type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// Función para generar un JWT
func GenerateJWT(username string) (string, error) {
	expirationTime := time.Now().Add(15 * time.Minute)

	claims := &Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(JwtKey)
}
