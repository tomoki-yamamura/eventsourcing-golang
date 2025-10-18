package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type App struct {
	DB     *gorm.DB
	Router *gin.Engine
}

func NewApp() *App {
	app := &App{}
	app.initDB()
	app.initRouter()
	return app
}

func (app *App) initDB() {
	dsn := "root:password@tcp(localhost:3306)/eventsourcing?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	app.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
}

func (app *App) initRouter() {
	app.Router = gin.Default()
	
	app.Router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "OK",
		})
	})
}

func (app *App) Run(addr string) error {
	return app.Router.Run(addr)
}

func main() {
	app := NewApp()
	
	log.Println("Starting server on :8080")
	if err := app.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}