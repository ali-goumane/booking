package sql

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// Informations de connexion à MySQL
var (
	username = "root"
	password = "root"
	hostname = "localhost"
	port     = 3306
	dbname   = "booking"
)

// Connexion à la base de données MySQL
func InitDB() *sql.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", username, password, hostname, port, dbname)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	// Vérifier si la connexion est établie
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	return db
}
