//package Verification
//
//import (
//	//"fmt"
//	//"math/rand"
//	"net/http"
//	"regexp"
//	//"strconv"
//	//"time"
//
//	"github.com/gin-gonic/gin"
//	//"gopkg.in/gomail.v2"
//)
//
//func VerifyCreditionals(ginContext *gin.Context) {
//	cred := new(Creditionals)
//	cred.Name = ginContext.Query("Name")
//	cred.EmailID = ginContext.Query("Email")
//	cred.DOB = ginContext.Query("DOB")
//
//	if cred.EmailID == "" || cred.Name == "" || cred.DOB == "" {
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
//	ginContext.Set(cred.EmailID, "Email")
//	ginContext.Set(cred.Name, "Name")
//	ginContext.Set(cred.DOB, "DOB")
//
//	ginContext.Next()
//}

package VerifyCredentials

import (
	"database/sql"
	"errors"
)

type Database struct {
	db *sql.DB
}

func NewDatabase(connString string) (*Database, error) {
	db, err := sql.Open("postgres", connString)
	if err != nil {
		return nil, err
	}

	return &Database{db: db}, nil
}

func (d *Database) Close() error {
	return d.db.Close()
}

func (d *Database) VerifyCredentials(email, password string) error {
	var storedPassword string
	err := d.db.QueryRow("SELECT password FROM users WHERE email = $1", email).Scan(&storedPassword)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("invalid email or password")
		}
		return err
	}

	if password != storedPassword {
		return errors.New("invalid email or password")
	}

	return nil
}
