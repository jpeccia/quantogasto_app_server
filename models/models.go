package models

import "time"

type Usuario struct {
	ID         int       `json:"id"`
	Nome       string    `json:"nome"`
	FotoPerfil string    `json:"foto_perfil"` // URL ou caminho da foto
	Cargo      string    `json:"cargo"`
	Renda      float64   `json:"renda"`
	CreatedAt  time.Time `json:"created_at"`
}

type GastoFixo struct {
	ID    int     `json:"id"`
	Nome  string  `json:"nome"`
	Valor float64 `json:"valor"`
}

type GastoVariavel struct {
	ID    int     `json:"id"`
	Nome  string  `json:"nome"`
	Valor float64 `json:"valor"`
	Data  string  `json:"data"`
}