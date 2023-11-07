package mesStorage

import (
	cfg "TaskService/internal/app/api/config"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

func InitDb(cfg *cfg.Config) (*sql.DB, error) {
	const op = "memStorage.InitDb"

	dbCfg := cfg.DbConfig

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s",
		dbCfg.Host, dbCfg.Port, dbCfg.User, dbCfg.Password, dbCfg.DbName)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("error with connect to db: %w", err)
	}

	log.Printf("succesful conntct to: %s", cfg.DbName)
	return db, nil
}
