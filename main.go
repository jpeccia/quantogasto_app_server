package main

import (
	"log"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/jpeccia/quantogasto_app_server/database"
	"github.com/jpeccia/quantogasto_app_server/handlers"
	middleware "github.com/jpeccia/quantogasto_app_server/middlewares"
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

	// Inicializa o roteador do Gin
	r := gin.Default()

	// Configura o middleware CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:8081"},         // URL do seu app Expo
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},  // Métodos permitidos
		AllowHeaders:     []string{"Content-Type", "Authorization"}, // Cabeçalhos permitidos
		AllowCredentials: true,                                      // Se você estiver utilizando cookies ou autenticação
	}))

	// Middleware global (opcional)
	r.Use(gin.Logger())   // Log de todas as requisições
	r.Use(gin.Recovery()) // Recupera de panics

	// Rotas de usuários
	usuarios := r.Group("/usuarios")
	{
		usuarios.POST("", handlers.RegistrarUsuario) // Cadastra um novo usuário
	}

	// Rotas protegidas por autenticação
	auth := r.Group("/")
	auth.Use(middleware.Autenticar()) // Middleware de autenticação aplicado
	{
		auth.PUT("/gastos-fixos/:id", handlers.EditarGastoFixo)             // Edita um gasto fixo
		auth.PUT("/gastos-variaveis/:id", handlers.EditarGastoVariavel)     // Edita um gasto variável
		auth.DELETE("/gastos-fixos/:id", handlers.RemoverGastoFixo)         // Remove um gasto fixo
		auth.DELETE("/gastos-variaveis/:id", handlers.RemoverGastoVariavel) // Remove um gasto variável
		auth.PUT("/renda", handlers.AdicionarRenda)                         // Adiciona renda
		auth.POST("/gastos-fixos", handlers.AdicionarGastoFixo)             // Adiciona gasto fixo
		auth.POST("/gastos-variaveis", handlers.AdicionarGastoVariavel)     // Adiciona gasto variável
		auth.POST("/usuarios/foto", handlers.UploadFotoPerfil)              // Rota para upload de foto de perfil
		auth.GET("/resumo", handlers.ObterResumo)                           // Obtém resumo financeiro
		auth.GET("/usuarios/:id", handlers.ObterUsuario)                    // Obtém dados de um usuário
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
