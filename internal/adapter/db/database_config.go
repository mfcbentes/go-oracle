package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/godror/godror"

	"github.com/mfcbentes/go-oracle/internal/domain"
)

type OracleDB struct {
	Config domain.Config
	DB     *sql.DB
}

func (o *OracleDB) Connect() error {
	dsn := fmt.Sprintf(
		`user="%s" password="%s" connectString="%s:%s/%s"`,
		o.Config.DBUser, o.Config.DBPassword, o.Config.DBHost, o.Config.DBPort, o.Config.DBService,
	)

	db, err := sql.Open("godror", dsn)
	if err != nil {
		return fmt.Errorf("erro ao conectar ao banco: %w", err)
	}

	o.DB = db
	if err := db.Ping(); err != nil {
		return fmt.Errorf("erro ao testar conexão: %w", err)
	}

	log.Println("Conexão com o banco estabelecida")
	return nil
}

func (o *OracleDB) Close() error {
	if o.DB != nil {
		return o.DB.Close()
	}
	return nil
}
