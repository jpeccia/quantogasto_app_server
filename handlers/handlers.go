package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jpeccia/quantogasto_app_server/database"
	"github.com/jpeccia/quantogasto_app_server/models"
)

// AdicionarRenda adiciona a renda mensal do usuário
func AdicionarRenda(c *gin.Context) {
	usuarioID := c.GetInt("usuario_id") // Obtém o ID do usuário do contexto (middleware de autenticação)

	var input struct {
		Valor float64 `json:"valor" binding:"required"`
	}
	// Valida o JSON recebido
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "O campo 'valor' é obrigatório e deve ser um número válido"})
		return
	}

	// Verifica se o valor é positivo
	if input.Valor <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "O valor deve ser maior que zero"})
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

	var input struct {
		Nome  string  `json:"nome" binding:"required"`
		Valor float64 `json:"valor" binding:"required"`
	}

	// Valida o JSON recebido
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Os campos 'nome' e 'valor' são obrigatórios"})
		return
	}

	// Verifica se o valor é positivo
	if input.Valor <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "O valor deve ser maior que zero"})
		return
	}

	// Insere o gasto fixo no banco de dados
	query := `
        INSERT INTO gastos_fixos (usuario_id, nome, valor)
        VALUES ($1, $2, $3)
    `
	_, err := database.DB.Exec(context.Background(), query, usuarioID, input.Nome, input.Valor)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao adicionar gasto fixo"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Gasto fixo adicionado com sucesso!"})
}


// AdicionarGastoVariavel adiciona um gasto variável do usuário
func AdicionarGastoVariavel(c *gin.Context) {
	usuarioID := c.GetInt("usuario_id") // Obtém o ID do usuário do contexto

	var input struct {
		Nome  string  `json:"nome" binding:"required"`
		Valor float64 `json:"valor" binding:"required"`
		Data  string  `json:"data" binding:"required"`
	}

	// Valida o JSON recebido
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Os campos 'nome', 'valor' e 'data' são obrigatórios"})
		return
	}

	// Verifica se o valor é positivo
	if input.Valor <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "O valor deve ser maior que zero"})
		return
	}

	// Valida a data no formato esperado (YYYY-MM-DD)
	if _, err := time.Parse("2006-01-02", input.Data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "A data deve estar no formato YYYY-MM-DD"})
		return
	}

	// Insere o gasto variável no banco de dados
	query := `
        INSERT INTO gastos_variaveis (usuario_id, nome, valor, data)
        VALUES ($1, $2, $3, $4)
    `
	_, err := database.DB.Exec(context.Background(), query, usuarioID, input.Nome, input.Valor, input.Data)
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

// EditarGastoFixo atualiza um gasto fixo do usuário
func EditarGastoFixo(c *gin.Context) {
	usuarioID := c.GetInt("usuario_id") // Obtém o ID do usuário do contexto
	gastoID := c.Param("id")            // Obtém o ID do gasto da URL

	var input struct {
		Nome  string  `json:"nome" binding:"required"`
		Valor float64 `json:"valor" binding:"required"`
	}

	// Valida o JSON recebido
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Os campos 'nome' e 'valor' são obrigatórios"})
		return
	}

	// Verifica se o valor é positivo
	if input.Valor <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "O valor deve ser maior que zero"})
		return
	}

	// Atualiza o gasto fixo no banco de dados
	query := `
        UPDATE gastos_fixos
        SET nome = $1, valor = $2
        WHERE id = $3 AND usuario_id = $4
    `
	result, err := database.DB.Exec(context.Background(), query, input.Nome, input.Valor, gastoID, usuarioID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao atualizar o gasto fixo"})
		return
	}

	// Verifica se o gasto foi encontrado e atualizado
	if rowsAffected, _ := result.RowsAffected(); rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Gasto fixo não encontrado ou você não tem permissão para editá-lo"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Gasto fixo atualizado com sucesso!"})
}

// EditarGastoVariavel atualiza um gasto variável do usuário
func EditarGastoVariavel(c *gin.Context) {
	usuarioID := c.GetInt("usuario_id") // Obtém o ID do usuário do contexto
	gastoID := c.Param("id")            // Obtém o ID do gasto da URL

	var input struct {
		Nome  string  `json:"nome" binding:"required"`
		Valor float64 `json:"valor" binding:"required"`
		Data  string  `json:"data" binding:"required"`
	}

	// Valida o JSON recebido
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Os campos 'nome', 'valor' e 'data' são obrigatórios"})
		return
	}

	// Verifica se o valor é positivo
	if input.Valor <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "O valor deve ser maior que zero"})
		return
	}

	// Valida o formato da data
	if _, err := time.Parse("2006-01-02", input.Data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "A data deve estar no formato YYYY-MM-DD"})
		return
	}

	// Atualiza o gasto variável no banco de dados
	query := `
        UPDATE gastos_variaveis
        SET nome = $1, valor = $2, data = $3
        WHERE id = $4 AND usuario_id = $5
    `
	result, err := database.DB.Exec(context.Background(), query, input.Nome, input.Valor, input.Data, gastoID, usuarioID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao atualizar o gasto variável"})
		return
	}

	// Verifica se o gasto foi encontrado e atualizado
	if rowsAffected, _ := result.RowsAffected(); rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Gasto variável não encontrado ou você não tem permissão para editá-lo"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Gasto variável atualizado com sucesso!"})
}

// RemoverGastoFixo remove um gasto fixo do usuário
func RemoverGastoFixo(c *gin.Context) {
	usuarioID := c.GetInt("usuario_id") // Obtém o ID do usuário do contexto
	gastoID := c.Param("id")            // Obtém o ID do gasto da URL

	// Remove o gasto fixo do banco de dados
	query := `
        DELETE FROM gastos_fixos
        WHERE id = $1 AND usuario_id = $2
    `
	result, err := database.DB.Exec(context.Background(), query, gastoID, usuarioID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao remover o gasto fixo"})
		return
	}

	// Verifica se o gasto foi encontrado e removido
	if rowsAffected, _ := result.RowsAffected(); rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Gasto fixo não encontrado ou você não tem permissão para removê-lo"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Gasto fixo removido com sucesso!"})
}

// RemoverGastoVariavel remove um gasto variável do usuário
func RemoverGastoVariavel(c *gin.Context) {
	usuarioID := c.GetInt("usuario_id") // Obtém o ID do usuário do contexto
	gastoID := c.Param("id")            // Obtém o ID do gasto da URL

	// Remove o gasto variável do banco de dados
	query := `
        DELETE FROM gastos_variaveis
        WHERE id = $1 AND usuario_id = $2
    `
	result, err := database.DB.Exec(context.Background(), query, gastoID, usuarioID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao remover o gasto variável"})
		return
	}

	// Verifica se o gasto foi encontrado e removido
	if rowsAffected, _ := result.RowsAffected(); rowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Gasto variável não encontrado ou você não tem permissão para removê-lo"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Gasto variável removido com sucesso!"})
}


// Registrar Usuário registra o Nome do usuário
func RegistrarUsuario(c *gin.Context) {
    var usuario models.Usuario

    // Bind do JSON recebido para a struct Usuario
    if err := c.ShouldBindJSON(&usuario); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Insere o usuário no banco de dados
    query := `
        INSERT INTO usuarios (nome, foto_perfil, cargo, renda)
        VALUES ($1, $2, $3, $4)
        RETURNING id
    `
    var id int
    err := database.DB.QueryRow(context.Background(), query,
        usuario.Nome, usuario.FotoPerfil, usuario.Cargo, usuario.Renda,
    ).Scan(&id)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao registrar usuário"})
        return
    }

    // Retorna o ID do usuário criado
    c.JSON(http.StatusOK, gin.H{
        "message": "Usuário registrado com sucesso!",
        "id":      id,
    })
}

// Obter Usuario retorna os dados do usuário
func ObterUsuario(c *gin.Context) {
    id := c.Param("id") // Obtém o ID do usuário da URL

    var usuario models.Usuario
    query := `SELECT id, nome, foto_perfil, cargo, renda, created_at FROM usuarios WHERE id = $1`
    err := database.DB.QueryRow(context.Background(), query, id).Scan(
        &usuario.ID, &usuario.Nome, &usuario.FotoPerfil, &usuario.Cargo, &usuario.Renda, &usuario.CreatedAt,
    )
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Usuário não encontrado"})
        return
    }

    c.JSON(http.StatusOK, usuario)
}