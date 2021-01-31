package database

import (
	"github.com/alex-pumnea/go-voting-api/internal/config"
	_ "github.com/go-sql-driver/mysql" // mysql driver
	"github.com/jmoiron/sqlx"
)

func migrateUp(db *sqlx.DB) error {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS user (
			id varchar(40) NOT NULL,
			email varchar(250) NOT NULL,
			password varchar(250) NOT NULL,
			name varchar(200) NOT NULL,
			is_admin BOOL DEFAULT false NOT NULL,
			CONSTRAINT user_PK PRIMARY KEY (id),
			CONSTRAINT email_UNQ UNIQUE KEY (email)
		)
		ENGINE=InnoDB
		DEFAULT CHARSET=utf8mb4
		COLLATE=utf8mb4_0900_ai_ci;
	`)

	return err
}

// NewDB ...
func NewDB(config *config.Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("mysql", config.ConnectionStr)

	if err != nil {
		return nil, err
	}

	if config.Environment == "local" {
		if err = migrateUp(db); err != nil {
			return nil, err
		}
	}

	return db, err
}
