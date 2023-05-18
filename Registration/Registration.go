package RegisterUser

import (
	//"fmt"
	
	"net/http"
	"regexp"
	//"strconv"
	//"time"

	. "github.com/Immerser01/InternAssignment/tree/main/Creditionals"
	"github.com/gin-gonic/gin"
	//"gopkg.in/gomail.v2"
)

func RegisterUser(ginContext *gin.Context) {
	cred := new(Creditionals)
	cred.Name = ginContext.Query("Name")
	cred.EmailID = ginContext.Query("Email")
	cred.DOB = ginContext.Query("DOB")
	cred.Password = ginContext.Query("Password")

	if cred.EmailID == "" || cred.Name == "" || cred.DOB == "" || cred.Password == "" {
		ginContext.JSON(http.StatusBadRequest, gin.H{
			"ERROR": "None of the Creditional fields can be empty!",
		})
		ginContext.Abort()
		return
	}

	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(cred.EmailID) {
		ginContext.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email format"})
		return
	}

	DOBRegex := regexp.MustCompile(`^\d{4}-\d{2}-\d{2}$`)
	if !DOBRegex.MatchString(cred.DOB) {
		ginContext.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date of birth format"})
		return

	}
	ginContext.JSON(http.StatusOK, gin.H{"message": "User successfully registered"})
	ginContext.Set(cred.EmailID, "Email")
	ginContext.Set(cred.Name, "Name")
	ginContext.Set(cred.DOB, "DOB")
	ginContext.Set(cred.Password, "Password")

}
