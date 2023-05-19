package main

import (
	"database/sql"
	"fmt"
	"github.com/Immerser01/InternAssignment/tree/main/Storage/Database"
	"log"

	"github.com/Immerser01/InternAssignment/tree/main/Handler/LoginHandler"
	"github.com/Immerser01/InternAssignment/tree/main/Handler/RegisterHandler"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

const (
	dbHost     = "localhost"
	dbPort     = 5432
	dbUser     = "sylvian-knight"
	dbPassword = "Root2is1!414"
	dbName     = "InternAssignment"
)

func main() {
	// Initialize the database connection
	dbInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)
	db, err := sql.Open("postgres", dbInfo)
	if err != nil {
		log.Fatal(err)
	}

	// Create a new Gin router
	router := gin.Default()

	// Create the storage instance
	storage := Database.NewStorage(db)

	// Routes
	router.POST("/register", RegisterHandler.RegisterHandler(storage))
	router.POST("/login", LoginHandler.LoginHandler(storage))

	// Start the server
	err = router.Run(":8080")
	if err != nil {
		return
	}
}

//import (
//	//"fmt"
//
//	"net/http"
//	"regexp"
//
//	"github.com/gin-gonic/gin"
//	//"gopkg.in/gomail.v2"
//)
//
//func RegisterUser(ginContext *gin.Context) {
//	cred := new(Creditionals)
//	cred.Name = ginContext.Query("Name")
//	cred.EmailID = ginContext.Query("Email")
//	cred.DOB = ginContext.Query("DOB")
//	cred.Password = ginContext.Query("Password")
//
//	if cred.EmailID == "" || cred.Name == "" || cred.DOB == "" || cred.Password == "" {
//		ginContext.JSON(http.StatusBadRequest, gin.H{
//			"ERROR": "None of the Creditional fields can be empty!",
//		})
//		ginContext.Abort()
//		return
//	}
//
//	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
//	if !emailRegex.MatchString(cred.EmailID) {
//		ginContext.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email format"})
//		return
//	}
//
//	DOBRegex := regexp.MustCompile(`^\d{4}-\d{2}-\d{2}$`)
//	if !DOBRegex.MatchString(cred.DOB) {
//		ginContext.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date of birth format"})
//		return
//
//	}
//
//	ginContext.JSON(http.StatusOK, gin.H{"message": "User successfully registered"})
//	//ginContext.Set(cred.EmailID, "Email")
//	//ginContext.Set(cred.Name, "Name")
//	//ginContext.Set(cred.DOB, "DOB")
//	//ginContext.Set(cred.Password, "Password")
//
//}
