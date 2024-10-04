package logs

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"strings"
)

// ProcessLogFile lê o arquivo de logs e processa as linhas
func ProcessLogFile(logFilePath string, db *sql.DB) error {
	file, err := os.Open(logFilePath)
	if err != nil {
		return fmt.Errorf("erro ao abrir o arquivo de log: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if shouldSaveLog(line) {
			if err := saveLogToDB(line, db); err != nil {
				return fmt.Errorf("erro ao salvar log no banco de dados: %w", err)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("erro ao ler o arquivo de log: %w", err)
	}

	return nil
}

// shouldSaveLog verifica se a linha contém GET ou POST
func shouldSaveLog(line string) bool {
	return strings.Contains(line, "GET") || strings.Contains(line, "POST")
}

// saveLogToDB salva a linha no banco de dados
func saveLogToDB(line string, db *sql.DB) error {
	query := "INSERT INTO logs (content) VALUES ($1)"
	_, err := db.Exec(query, line)
	if err != nil {
		return fmt.Errorf("erro ao inserir log no banco de dados: %w", err)
	}
	return nil
}
