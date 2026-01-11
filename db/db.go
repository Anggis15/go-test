package db

import (
	"database/sql"

	_ "github.com/microsoft/go-mssqldb"
)

func NewDB() (*sql.DB, error) {
	// jdbc:sqlserver://;serverName=localhost;port=14330;databaseName=master
	conn := "server=localhost,14330;user id=sa;password=KuatPass123!@#;database=quizdb;encrypt=true;trustservercertificate=true"
	db, err := sql.Open("sqlserver", conn)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(10)
	return db, nil
}
