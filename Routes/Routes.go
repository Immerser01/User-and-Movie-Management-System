package Routes

import (
	"database/sql"
	"github.com/Immerser01/InternAssignment/tree/main/Admin"
	"github.com/Immerser01/InternAssignment/tree/main/Handler/CredentialHandler"
	"github.com/Immerser01/InternAssignment/tree/main/Handler/Moviehandler"
	"github.com/Immerser01/InternAssignment/tree/main/Handler/UserHandler"
	"github.com/gin-gonic/gin"
)

func StartRoutes(r *gin.Engine, db *sql.DB) {
	userHandler := &UserHandler.UserHandler{
		DB: db,
	}
	movieHandler := &Moviehandler.MovieHandler{
		DB: db,
	}
	credentialHandler := &CredentialHandler.CredentialHandler{
		DB: db,
	}
	adminHandler := &Admin.AdminHandler{
		DB: db,
	}

	r.POST("/users", userHandler.CreateUser)
	r.POST("/credentials", credentialHandler.UpdateCredentials)
	r.GET("/users/:accessPassword", userHandler.ListUsers)
	r.GET("/AdminCredentialsPage/:mainPassword", adminHandler.ListCredentials)
	r.POST("/movies", movieHandler.AddMovie)
	r.DELETE("/movies/:id", movieHandler.DeleteMovie)
	r.GET("/movies/:id/:password", movieHandler.ListMoviesByUser)
	r.POST("/admin", adminHandler.PasswordManager)
	r.DELETE("admin/:password/:mainPassword", adminHandler.DeletePassword)
	r.GET("admin/:mainPassword", adminHandler.ListPassword)
}
