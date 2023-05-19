package RegisterHandler

import (
	"github.com/Immerser01/InternAssignment/tree/main/Storage/Database"
	"github.com/Immerser01/InternAssignment/tree/main/Storage/TableStructs"
	"github.com/gin-gonic/gin"
	"regexp"
)

func RegisterHandler(storage *Database.Database) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Extract username and password from the request body

		var usr TableStructs.UserData
		if err := c.ShouldBindJSON(&usr); err != nil {
			c.JSON(400, gin.H{"error": "Invalid request"})
			return
		}

		if usr.Email == "" || usr.Name == "" || usr.DOB == "" || usr.Password == "" {
			c.JSON(500, gin.H{
				"ERROR": "None of the Credentials fields can be empty!"})
			return
		}

		emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
		if !emailRegex.MatchString(usr.Email) {
			c.JSON(501, gin.H{"error": "Invalid email format"})
			return
		}

		DOBRegex := regexp.MustCompile(`^\d{2}-\d{2}-\d{4}$`)
		if !DOBRegex.MatchString(usr.DOB) {
			c.JSON(502, gin.H{"error": "Invalid date of birth format"})
			return
		}

		//_, errornext1 := storage.Db.Exec("INSERT INTO Credentials (Email, Password) VALUES ($1, $2)", usr.Email, usr.Password)
		//if errornext1 != nil {
		//	c.JSON(503, gin.H{"error": "Failed to register Credentials"})
		//	return
		//}
		_, errornext2 := storage.Db.Exec("INSERT INTO UserData (Name, Email, DOB) VALUES ($1, $2, $3)", usr.Name, usr.Email, usr.DOB)
		if errornext2 != nil {
			c.JSON(504, gin.H{"error": "Failed to register User Details"})
			return
		}
		//_, err := storage.DB.Exec("INSERT INTO users (username, password) VALUES ($1, $2)", user.Username, user.Password)
		//if err != nil {
		//	c.JSON(500, gin.H{"error": "Failed to register user"})
		//	return
		//}

		c.JSON(200, gin.H{"message": "User registered successfully"})
	}
}

//curl -X POST -H "Content-Type: application/json"     -d '{"Name": "linuxize", "Email": "linuxize@example.com", "Password": "37294618713", "DOB": "07-08-2002"}' \localhost:8080/register
