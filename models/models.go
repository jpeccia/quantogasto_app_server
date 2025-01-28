package models

type Usuario struct {
	ID	int	`json:"id"`
	Nome	string	`json:"nome"`
	Renda	float64	`json:"renda"`
}

type GastoFixo struct {
	ID	int	`json:"id"`
	Nome	string	`json:"nome"`
	Valor	float64	`json:"valor"`
}

type GastoVariavel struct {
	ID	int	`json:"id"`
	Nome	string	`json:"nome"`
	Valor	float64	`json:"valor"`
	Data	string	`json:"data"`
}