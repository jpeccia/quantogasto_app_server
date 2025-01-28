package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Autenticar() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Obtém o nome do usuário do header ou query string
		nomeUsuario := c.Query("nome")
		if nomeUsuario == "" {
			nomeUsuario = c.GetHeader("X-Nome-Usuario")
		}

		// Se o nome do usuário não for fornecido, retorna erro
		if nomeUsuario == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Nome do usuário é obrigatório"})
			c.Abort()
			return
		}

		// Armazena o nome do usuário no contexto
		c.Set("usuario_nome", nomeUsuario)

		// Passa para o próximo handler
		c.Next()
	}
}
