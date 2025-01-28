package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/jpeccia/quantogasto_app_server/database"
	"github.com/jpeccia/quantogasto_app_server/handlers"
	"github.com/jpeccia/quantogasto_app_server/middlewares" // Importa o middleware de autenticação
)

func main() {
	// Carrega as variáveis de ambiente
	if err := godotenv.Load(); err != nil {
		log.Println("Aviso: Arquivo .env não encontrado, usando variáveis do sistema.")
	}

	// Verifica se as variáveis de ambiente obrigatórias estão configuradas
	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("Erro: DB_URL não configurado.")
	}

	// Conecta ao banco de dados
	if err := database.Connect(); err != nil {
		log.Fatal("Erro ao conectar ao banco de dados: ", err)
	}

	// Inicializa o roteador
	r := gin.Default()

	// Middleware global (opcional)
	r.Use(gin.Logger())   // Log de todas as requisições
	r.Use(gin.Recovery()) // Recupera de panics

	// Rotas de usuários
	usuarios := r.Group("/usuarios")
	{
		usuarios.POST("/", handlers.RegistrarUsuario) // Cadastra um novo usuário
		usuarios.GET("/:id", handlers.ObterUsuario)   // Obtém dados de um usuário
	}

	// Rotas protegidas por autenticação
	auth := r.Group("/")
	auth.Use(middleware.Autenticar()) // Middleware de autenticação aplicado
	{
		auth.POST("/renda", handlers.AdicionarRenda)
		auth.POST("/gastos-fixos", handlers.AdicionarGastoFixo)
		auth.POST("/gastos-variaveis", handlers.AdicionarGastoVariavel)
		auth.GET("/resumo", handlers.ObterResumo)
	}

	// Inicia o servidor
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Porta padrão
	}
	log.Printf("Servidor iniciado na porta %s...", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatal("Erro ao iniciar o servidor: ", err)
	}
}
