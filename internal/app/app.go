package app

import (
	"log"

	"github.com/mfcbentes/go-oracle/internal/domain"
)

type App struct {
	Database domain.Database
}

func (a *App) Start() {
	if err := a.Database.Connect(); err != nil {
		log.Fatalf("Erro ao iniciar o app: %v", err)
	}
	defer a.Database.Close()

	log.Println("Aplicação rodando!")
}
