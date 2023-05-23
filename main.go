package main

import (
	"database/sql"
	_ "github.com/Immerser01/InternAssignment/tree/main/Handler/CredentialHandler"
	"github.com/Immerser01/InternAssignment/tree/main/Models"
	"github.com/Immerser01/InternAssignment/tree/main/Routes"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	// Database connection
	db, err := sql.Open("postgres", "postgres://sylvian-knight:Root2is1!414@localhost/InternAssignment?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			return
		}
	}(db)

	Models.CreateTables(db)

	r := gin.Default()
	Routes.StartRoutes(r, db)

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
