package main

import (
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/kofiasare/book-challenge/controllers"
	"github.com/kofiasare/book-challenge/db"
	"github.com/kofiasare/book-challenge/models"
	"github.com/kofiasare/book-challenge/utils"
)

func main() {

	r := gin.Default()

	pg := db.GetPgClient()

	models := []interface{}{&models.User{}, &models.Author{}, &models.Book{}}

	// migrate models
	// pg.Migrator().DropTable(models...)
	pg.AutoMigrate(models...)

	r.GET("/", controllers.Root)
	r.POST("/register", controllers.RegisterUser)
	r.POST("/login", controllers.LoginUser)
	r.POST("/api/authors", utils.AuthMiddleware, controllers.CreateAuthor)
	r.GET("/api/books", controllers.IndexBooks)
	r.GET("/api/books/:id", utils.AuthMiddleware, controllers.ShowBook)
	r.POST("/api/books", utils.AuthMiddleware, controllers.CreateBook)
	r.PUT("/api/books/:id", utils.AuthMiddleware, controllers.UpdateBook)
	r.DELETE("/api/books/:id", utils.AuthMiddleware, controllers.DeleteBook)

	defer pg.Close()

	r.Use(cors.Default())
	r.Run(":" + os.Getenv("PORT"))
}
