package exercice2

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)
var DB *sql.DB

func ConnectDB() {
	var err error
	DB, err = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/bd_employee")
	if err != nil {
		log.Fatal("Erreur lors de la connexion à la base de données :", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("Erreur lors du ping de la base de données :", err)
	}

	fmt.Println("Connexion à la base de données réussie!")
}
