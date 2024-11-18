package main

import (
	"log"

	"github.com/mfcbentes/go-oracle/internal/adapter/db"
	"github.com/mfcbentes/go-oracle/internal/app"
	"github.com/mfcbentes/go-oracle/internal/config"
)

func main() {
	// Carregar configurações
	cfg := config.LoadConfig()

	// Inicializar banco de dados
	oracleDB := &db.OracleDB{Config: cfg}

	// Inicializar aplicação
	application := &app.App{Database: oracleDB}

	// Rodar aplicação
	application.Start()
	log.Println("Encerrando aplicação...")
}
