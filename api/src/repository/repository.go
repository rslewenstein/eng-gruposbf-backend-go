package repository

import (
	"database/sql"

	"go.mod/src/models"
)

// Representa um repositório de conversões
type Symbol struct {
	db *sql.DB
}

// Cria um repositório
func NewRepository(db *sql.DB) *Symbol {
	return &Symbol{db}
}

// Insere um país e a sigla da moeda no banco de dados
func (repository Symbol) CreateSymbol(symbol models.Symbol) (uint64, error) {
	statement, err := repository.db.Prepare(
		"INSERT INTO currencies (symbol, country) VALUES (?, ?)",
	)
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	result, err := statement.Exec(symbol.Symbol, symbol.Country)
	if err != nil {
		return 0, err
	}

	lastIDinserted, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(lastIDinserted), nil
}
