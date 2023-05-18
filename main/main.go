package main

import (
	//"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	//"net/http"
	//"github.com/gorilla/mux"
)

func main() {
	router := gin.Default()
	log.Fatal(router.Run(":8080"))

}
