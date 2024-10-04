package database

import (
    "database/sql"
    "log"
    "github.com/pressly/goose/v3"
    _ "github.com/lib/pq"
)

// RunMigrations roda as migrações no banco de dados usando Goose
func RunMigrations(dbConn *sql.DB, migrationDir string) error {
    // Configura o dialeto do Goose para Postgres
    if err := goose.SetDialect("postgres"); err != nil {
        return err
    }

    // Aplica as migrações Up (somente as pendentes)
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

    // Roda as migrações Up novamente para recriar a tabela
    if err := RunMigrations(dbConn, migrationDir); err != nil {
        return err
    }

    log.Println("Tabela de logs resetada com sucesso.")
    return nil
}
