package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/jpeccia/quantogasto_app_server/models"
)

var renda float64
var gastosFixos []models.GastoFixo
var gastosVariaveis []models.GastoVariavel

func AdicionarRenda(c *gin.Context) {
	var input struct {
		Renda	float64 `json:"renda"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	renda = input.Renda
	c.JSON(http.StatusOK, gin.H{"message": "Renda adicionada com sucesso!"})
}

func AdicionarGastoFixo(c *gin.Context) {
	var gasto models.GastoFixo
	if err := c.ShouldBindJSON(&gasto); err != nil {
		c.JSON(http.StatusBadRequest, 
		gin.H{"error": err.Error()})
		return
	}
	gastosFixos = append(gastosFixos, gasto)
	c.JSON(http.StatusOK, gin.H{"message:": "Gasto fixo adicionado com sucesso!"})
}

func AdicionarGastoVariavel(c *gin.Context) {
	var gasto models.GastoVariavel
	if err := c.ShouldBindJSON(&gasto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	gastosVariaveis = append(gastosVariaveis, gasto)
	c.JSON(http.StatusOK, gin.H{"message": "Gasto variável adicionado com sucesso!"})
}

func ObterResumo(c *gin.Context) {
	totalGastosFixos := 0.0
	for _, gasto := range gastosFixos {
		totalGastosFixos += gasto.Valor
	}

	totalGastosVariaveis := 0.0
	for _, gasto := range gastosVariaveis {
		totalGastosVariaveis += gasto.Valor
	}

	saldo := renda - totalGastosFixos - totalGastosVariaveis

	c.JSON(http.StatusOK, gin.H{
		"renda":	renda,
		"total_gastos_fixos":	totalGastosFixos,
		"total_gastos_variaveis": 	totalGastosVariaveis,
		"saldo":	saldo,
	})
}