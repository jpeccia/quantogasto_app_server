package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jpeccia/quantogasto_app_server/auth"
)

func Autenticar() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Obtém o token do cabeçalho Authorization
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token não fornecido"})
			c.Abort()
			return
		}

		// Valida o token
		claims, err := auth.ValidarToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token inválido"})
			c.Abort()
			return
		}

		// Armazena o ID do usuário no contexto
		c.Set("usuario_id", claims.UsuarioID)

		// Passa para o próximo handler
		c.Next()
	}
}