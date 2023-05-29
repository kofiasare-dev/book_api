package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/kofiasare/book-challenge/controllers"
	"github.com/kofiasare/book-challenge/db"
	"github.com/kofiasare/book-challenge/models"
)

func main() {

	r := gin.Default()

	pg := db.GetPgClient()

	models := []interface{}{&models.User{}, &models.Author{}, &models.Book{}}

	// pg.Migrator().DropTable(models...)

	// migrate models
	pg.AutoMigrate(models...)

	r.GET("/", controllers.Root)
	r.POST("/api/authors", controllers.CreateAuthor)
	r.GET("/api/books", controllers.IndexBooks)
	r.GET("/api/books/:id", controllers.ShowBook)
	r.POST("/api/books", controllers.CreateBook)

	defer pg.Close()

	r.Run(":" + os.Getenv("PORT"))
}
