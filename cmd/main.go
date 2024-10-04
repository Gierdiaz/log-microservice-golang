package main

import (
	"fmt"
	"log"
	"time"

	"github.com/Gierdiaz/Log-Service/pkg/database"
	"github.com/Gierdiaz/Log-Service/pkg/logs"
	"github.com/Gierdiaz/Log-Service/pkg/scheduler"
)

func main() {
	fmt.Println("Iniciando o microserviço de leitura de logs...")

	dbConn, err := database.Connect()
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
	}
	defer dbConn.Close()

	migrationDir := "./pkg/database/migrations"

	if err := database.RunMigrations(dbConn, migrationDir); err != nil {
		log.Fatalf("Erro ao rodar as migrações: %v", err)
	}

	scheduler.Every(5*time.Minute, func() {
		err := logs.ProcessLogFile("logs/access.log", dbConn)
		if err != nil {
			log.Printf("Erro ao processar o arquivo de log: %v", err)
		} else {
			log.Println("Processamento de logs concluído com sucesso.")
		}
	})

	// Mantém o serviço rodando
	select {}
}
