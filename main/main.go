package main

import (
	"github.com/Immerser01/InternAssignment/tree/main/Registration"
	"github.com/Immerser01/InternAssignment/tree/main/Storage/Database"
	"github.com/Immerser01/InternAssignment/tree/main/Verification/LoginCredentials"
	"github.com/Immerser01/InternAssignment/tree/main/Verification/VerifyCredentials"

	//"github.com/Immerser01/InternAssignment/tree/main/Verification"
	"github.com/gin-gonic/gin"
	//"net/http"
	//"github.com/gorilla/mux"
	"log"
)

func main() {
	db, err := Database.NewDatabase("postgres://user:password@localhost/mydatabase?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer func(db *Database.Database) {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(db)

	vdb, err := VerifyCredentials.NewDatabase("postgres://user:password@localhost/mydatabase?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer func(vdb *VerifyCredentials.Database) {
		err := vdb.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(vdb)

	router := gin.Default()

	router.POST("/register", Registration.Registration(db))
	router.POST("/login", LoginCredentials.LoginCredentials(vdb))

	ERROR := router.Run(":8080")
	if ERROR != nil {
		return
	}
}
