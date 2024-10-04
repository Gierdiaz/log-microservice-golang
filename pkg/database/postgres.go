package database

import (
    "database/sql"
    "fmt"
    "log"
    "os"

    "github.com/joho/godotenv"
    _ "github.com/lib/pq"
)

func Connect() (*sql.DB, error) {
    err := godotenv.Load()
    if err != nil {
        log.Fatalf("Erro ao carregar o arquivo .env: %v", err)
    }

    // Recuperar as variáveis de ambiente
    Host := os.Getenv("POSTGRES_HOST")
    Port := os.Getenv("POSTGRES_PORT")
    User := os.Getenv("POSTGRES_USER")
    Password := os.Getenv("POSTGRES_PASSWORD")
    DBName := os.Getenv("POSTGRES_DBNAME")

    connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", Host, Port, User, Password, DBName)

    db, err := sql.Open("postgres", connStr)
    if err != nil {
        return nil, fmt.Errorf("erro ao abrir conexão com o banco: %w", err)
    }

    if err := db.Ping(); err != nil {
        return nil, fmt.Errorf("erro ao pingar o banco de dados: %w", err)
    }

    fmt.Println("Conexão com o banco de dados estabelecida!")
    return db, nil
}
