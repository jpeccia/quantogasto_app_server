package handlers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jpeccia/quantogasto_app_server/database"
	"github.com/jpeccia/quantogasto_app_server/models"
)

// AdicionarRenda adiciona a renda mensal do usuário
func AdicionarRenda(c *gin.Context) {
    usuarioID := c.GetInt("usuario_id") // Obtém o ID do usuário do contexto (middleware de autenticação)

    var input struct {
        Valor float64 `json:"valor"`
    }
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Insere a renda no banco de dados
    query := `
        INSERT INTO rendas (usuario_id, valor)
        VALUES ($1, $2)
    `
    _, err := database.DB.Exec(context.Background(), query, usuarioID, input.Valor)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao adicionar renda"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Renda adicionada com sucesso!"})
}

// AdicionarGastoFixo adiciona um gasto fixo do usuário
func AdicionarGastoFixo(c *gin.Context) {
    usuarioID := c.GetInt("usuario_id") // Obtém o ID do usuário do contexto

    var gasto models.GastoFixo
    if err := c.ShouldBindJSON(&gasto); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Insere o gasto fixo no banco de dados
    query := `
        INSERT INTO gastos_fixos (usuario_id, nome, valor)
        VALUES ($1, $2, $3)
    `
    _, err := database.DB.Exec(context.Background(), query, usuarioID, gasto.Nome, gasto.Valor)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao adicionar gasto fixo"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Gasto fixo adicionado com sucesso!"})
}

// AdicionarGastoVariavel adiciona um gasto variável do usuário
func AdicionarGastoVariavel(c *gin.Context) {
    usuarioID := c.GetInt("usuario_id") // Obtém o ID do usuário do contexto

    var gasto models.GastoVariavel
    if err := c.ShouldBindJSON(&gasto); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Insere o gasto variável no banco de dados
    query := `
        INSERT INTO gastos_variaveis (usuario_id, nome, valor, data)
        VALUES ($1, $2, $3, $4)
    `
    _, err := database.DB.Exec(context.Background(), query, usuarioID, gasto.Nome, gasto.Valor, gasto.Data)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao adicionar gasto variável"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Gasto variável adicionado com sucesso!"})
}

// ObterResumo retorna um resumo financeiro do usuário
func ObterResumo(c *gin.Context) {
    usuarioID := c.GetInt("usuario_id") // Obtém o ID do usuário do contexto

    // Obtém a renda total do usuário
    var rendaTotal float64
    queryRenda := `SELECT COALESCE(SUM(valor), 0) FROM rendas WHERE usuario_id = $1`
    err := database.DB.QueryRow(context.Background(), queryRenda, usuarioID).Scan(&rendaTotal)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar renda"})
        return
    }

    // Obtém o total de gastos fixos do usuário
    var gastosFixosTotal float64
    queryGastosFixos := `SELECT COALESCE(SUM(valor), 0) FROM gastos_fixos WHERE usuario_id = $1`
    err = database.DB.QueryRow(context.Background(), queryGastosFixos, usuarioID).Scan(&gastosFixosTotal)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar gastos fixos"})
        return
    }

    // Obtém o total de gastos variáveis do usuário
    var gastosVariaveisTotal float64
    queryGastosVariaveis := `SELECT COALESCE(SUM(valor), 0) FROM gastos_variaveis WHERE usuario_id = $1`
    err = database.DB.QueryRow(context.Background(), queryGastosVariaveis, usuarioID).Scan(&gastosVariaveisTotal)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar gastos variáveis"})
        return
    }

    // Calcula o saldo disponível
    saldo := rendaTotal - gastosFixosTotal - gastosVariaveisTotal

    // Retorna o resumo
    c.JSON(http.StatusOK, gin.H{
        "renda_total":           rendaTotal,
        "gastos_fixos_total":    gastosFixosTotal,
        "gastos_variaveis_total": gastosVariaveisTotal,
        "saldo_disponivel":      saldo,
    })
}