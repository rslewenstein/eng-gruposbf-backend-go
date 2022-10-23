package repository

import (
	"database/sql"

	"go.mod/src/models"
)

// Representa um repositório de conversões
type Connection struct {
	db *sql.DB
}

// Cria um repositório
func NewRepository(db *sql.DB) *Connection {
	return &Connection{db}
}

// Insere um país e a sigla da moeda no banco de dados
func (repository Connection) CreateSymbol(symbol models.Symbol) (uint64, error) {
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

// Retorna todos as moedas cadastradas no banco de dados
func (repository Connection) GetCurrencies() ([]models.Symbol, error) {
	ret, err := repository.db.Query(
		"SELECT id, symbol, country FROM currencies",
	)

	if err != nil {
		return nil, err
	}
	defer ret.Close()

	var symbols []models.Symbol

	for ret.Next() {
		var symbol models.Symbol

		if err = ret.Scan(
			&symbol.ID,
			&symbol.Symbol,
			&symbol.Country,
		); err != nil {
			return nil, err
		}

		symbols = append(symbols, symbol)
	}

	return symbols, nil
}

func (repository Connection) CreateLog(typelog, method, logdescription string) error {
	statement, err := repository.db.Prepare(
		"INSERT INTO currencylogs (typelog, method, logdescription) VALUES (?, ?, ?)",
	)
	if err != nil {
		return err
	}
	defer statement.Close()
	TypeLog := typelog
	Method := method
	LogDescription := logdescription

	result, err := statement.Exec(TypeLog, Method, LogDescription)
	if err != nil {
		return err
	}

	lastIDinserted, err := result.LastInsertId()
	if err != nil {
		return err
	}

	if lastIDinserted > 0 {
		return nil
	}

	return err
}
