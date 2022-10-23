package services

import (
	"go.mod/src/connection"
	"go.mod/src/repository"
)

// Cria Logs do sistema no banco de dados
func LoggerSystem(typelog, method, logdescription string) {
	if typelog != "" && method != "" && logdescription != "" {
		db, err := connection.Connect()
		if err != nil {
			println(err)
			return
		}
		defer db.Close()

		repository := repository.NewRepository(db)
		repository.CreateLog(typelog, method, logdescription)
	}
}
