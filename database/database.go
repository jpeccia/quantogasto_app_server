package database

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool // Usando pgxpool.Pool para gerenciar o pool de conexões

func Connect() error {
	// Obtém as variáveis de ambiente
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	// Monta a string de conexão
	connString := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		dbUser, dbPassword, dbHost, dbPort, dbName,
	)

	// Configura o pool de conexões
	config, err := pgxpool.ParseConfig(connString)
	if err != nil {
		return fmt.Errorf("erro ao parsear a string de conexão: %w", err)
	}

	// Cria o pool de conexões
	DB, err = pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		return fmt.Errorf("erro ao conectar ao banco de dados: %w", err)
	}

	// Testa a conexão
	err = DB.Ping(context.Background())
	if err != nil {
		return fmt.Errorf("erro ao testar a conexão com o banco de dados: %w", err)
	}

	fmt.Println("Conectado ao banco de dados!")
	return nil
}