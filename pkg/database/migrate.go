package database

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
	"log"
)

func RunMigrations(dbConn *sql.DB, migrationDir string) error {
	if err := goose.SetDialect("postgres"); err != nil {
		return err
	}
	if err := goose.Up(dbConn, migrationDir); err != nil {
		return err
	}
	log.Println("Migrações aplicadas com sucesso.")
	return nil
}

// ResetMigrations reseta a tabela ao iniciar o servidor
func ResetMigrations(dbConn *sql.DB, migrationDir string) error {
	log.Println("Resetando a tabela de logs...")

	// Aplica as migrações Down para resetar
	if err := goose.Down(dbConn, migrationDir); err != nil {
		return err
	}
	if err := RunMigrations(dbConn, migrationDir); err != nil {
		return err
	}
	log.Println("Tabela de logs resetada com sucesso.")
	return nil
}
