package models

import "time"

// Usuario representa um usuário do sistema
type Usuario struct {
    ID         int       `json:"id"`
    Nome       string    `json:"nome"`
    FotoPerfil string    `json:"foto_perfil"` // URL ou caminho da foto (opcional)
    Cargo      string    `json:"cargo"`       // Cargo do usuário (opcional)
    Renda      float64   `json:"renda"`       // Renda mensal do usuário
    CreatedAt  time.Time `json:"created_at"`  // Data de criação
}

// Renda representa a renda mensal de um usuário
type Renda struct {
    ID        int       `json:"id"`
    UsuarioID int       `json:"usuario_id"` // ID do usuário associado
    Valor     float64   `json:"valor"`      // Valor da renda
    CreatedAt time.Time `json:"created_at"` // Data de criação
}

// GastoFixo representa um gasto fixo de um usuário
type GastoFixo struct {
    ID        int       `json:"id"`
    UsuarioID int       `json:"usuario_id"` // ID do usuário associado
    Nome      string    `json:"nome"`       // Nome do gasto fixo
    Valor     float64   `json:"valor"`      // Valor do gasto fixo
    CreatedAt time.Time `json:"created_at"` // Data de criação
}

// GastoVariavel representa um gasto variável de um usuário
type GastoVariavel struct {
    ID        int       `json:"id"`
    UsuarioID int       `json:"usuario_id"` // ID do usuário associado
    Nome      string    `json:"nome"`       // Nome do gasto variável
    Valor     float64   `json:"valor"`      // Valor do gasto variável
    Data      string    `json:"data"`       // Data do gasto (formato YYYY-MM-DD)
    CreatedAt time.Time `json:"created_at"` // Data de criação
}