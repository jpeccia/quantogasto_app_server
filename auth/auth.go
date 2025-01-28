package auth

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte(os.Getenv("SECRETKEY")) // Substitua por uma chave segura

// Claims representa os dados armazenados no token JWT
type Claims struct {
    UsuarioID int `json:"usuario_id"`
    jwt.RegisteredClaims
}

// GerarToken gera um token JWT para o usuário
func GerarToken(usuarioID int) (string, error) {
    expirationTime := time.Now().Add(24 * time.Hour * 7) // Token expira em 7 dias

    claims := &Claims{
        UsuarioID: usuarioID,
        RegisteredClaims: jwt.RegisteredClaims{
            ExpiresAt: jwt.NewNumericDate(expirationTime),
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(jwtKey)
}

// ValidarToken valida um token JWT e retorna os claims
func ValidarToken(tokenString string) (*Claims, error) {
    claims := &Claims{}

    token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
        return jwtKey, nil
    })
    if err != nil {
        return nil, err
    }

    if !token.Valid {
        return nil, jwt.ErrSignatureInvalid
    }

    return claims, nil
}