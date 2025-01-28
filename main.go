package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	r.POST("/renda", handlers.AdicionarRenda)
	r.POST("/gastos-fixos", handlers.AdicionarGastoFixo)
	r.POST("/gastos-variaveis", handlers.AdicionarGastoVariavel)
	r.GET("/resumo", handlers.ObterResumo)

	r.Run(":8080")
}