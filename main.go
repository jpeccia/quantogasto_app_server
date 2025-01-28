package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/jpeccia/quantogasto_app_server/database"
	"github.com/jpeccia/quantogasto_app_server/handlers"
)

func main() {
	// Carrega as variáveis de ambiente do arquivo .env
	if err := godotenv.Load(); err != nil {
        log.Fatal("Erro ao carregar o arquivo .env")
    }

    // Conecta ao banco de dados
    if err := database.Connect(); err != nil {
        log.Fatal("Erro ao conectar ao banco de dados: ", err)
    }
	
	r := gin.Default()

	r.POST("/usuarios", handlers.RegistrarUsuario)       // Cadastra um novo usuário
    r.GET("/usuarios/:id", handlers.ObterUsuario)       // Obtém dados de um usuário
	r.POST("/renda", handlers.AdicionarRenda)
	r.POST("/gastos-fixos", handlers.AdicionarGastoFixo)
	r.POST("/gastos-variaveis", handlers.AdicionarGastoVariavel)
	r.GET("/resumo", handlers.ObterResumo)

	r.Run(":8080")
}