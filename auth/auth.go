package auth

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Claims representa os dados armazenados no token JWT
type Claims struct {
	UsuarioID int `json:"usuario_id"`
	jwt.RegisteredClaims
}

// InicializarJWT verifica se a chave secreta foi configurada corretamente
func InicializarJWT() error {
	jwtKey := []byte(os.Getenv("SECRETKEY"))
	if len(jwtKey) == 0 {
		return errors.New("variável de ambiente 'SECRETKEY' não configurada. Verifique seu arquivo .env ou variáveis do sistema.")
	}
	return nil
}

// GerarToken gera um token JWT para o usuário
func GerarToken(usuarioID int) (string, error) {
	// Verifica se a chave secreta está configurada
	if err := InicializarJWT(); err != nil {
		return "", err
	}

	// Define o tempo de expiração do token (7 dias)
	expirationTime := time.Now().Add(24 * time.Hour * 7)

	claims := &Claims{
		UsuarioID: usuarioID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	// Cria um novo token com os claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Assina o token com a chave secreta
	signedToken, err := token.SignedString(jwtKey)
	if err != nil {
		return "", fmt.Errorf("erro ao gerar o token: %v", err)
	}

	fmt.Println("Token gerado:", signedToken) // Exibe o token gerado para depuração
	return signedToken, nil
}

// ValidarToken valida um token JWT e retorna os claims
func ValidarToken(tokenString string) (*Claims, error) {
	claims := &Claims{}

	// Faz o parsing do token e valida os claims
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		// Verificando se o erro é devido ao token expirado
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, fmt.Errorf("token expirado. Por favor, faça login novamente.")
		}
		return nil, fmt.Errorf("erro ao validar o token: %v", err)
	}

	// Verifica se o token é válido (não corrompido ou alterado)
	if !token.Valid {
		return nil, fmt.Errorf("token inválido: assinatura não corresponde")
	}

	return claims, nil
}
