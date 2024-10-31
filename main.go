package main

import (
	"database/sql"
	"dish-dash/routes"
	"log"
   _ "github.com/jackc/pgx/v4/stdlib"
	"github.com/gin-gonic/gin"
)

var db *sql.DB

func main() {

	dsn := "postgres://postgres.<username>:<password>,<host>:<port>/postgres?sslmode=require"

	var err error
	db, err = sql.Open("pgx", dsn)
	if err != nil {
		log.Fatal("Error connecting to Database: ", err)
	}
	
	if err := db.Ping(); err != nil {
		log.Fatal("Error pinging Database: ", err)
	}

	defer db.Close()

    r := gin.Default()

	routes.RegisterRoutes(r, db)

	r.Run(":8080")
    
}