package main

import (
	_ "context"
	"fmt"
	"log"

	"github.com/bsanzhiev/tsurhai/controllers"
	"github.com/bsanzhiev/tsurhai/database"
	"github.com/bsanzhiev/tsurhai/firebaseapp"
	"github.com/bsanzhiev/tsurhai/middlewares"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile(".env")
	viperErr := viper.ReadInConfig()
	if viperErr != nil {
		panic(viperErr)
	}
	gin.SetMode(viper.GetString("APP_MODE"))

	database.Connect()
	database.Migrate()

	// Firebase app init
	if err := firebaseapp.InitFirebaseApp(); err != nil {
		log.Fatalf("Error while initializing Firebase: %v", err)
	}

	app := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowMethods = []string{"*"}
	config.AllowCredentials = true
	config.AllowHeaders = []string{"*"}
	app.Use(cors.New(config))

	app.GET("/ping", Ping)
	initRouter(app)

	log.Fatalln(
		app.Run(
			fmt.Sprintf("127.0.0.0:%d", viper.GetInt("APP_PORT")),
		),
	)
}

func Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func initRouter(app *gin.Engine) {

	api := app.Group("/auth")
	{
		api.POST("/test-token", controllers.GenerateToken)
		api.POST("/register", controllers.RegisterUser)
		api.POST("/login", controllers.LoginUser)
		api.POST("/verify-token", controllers.VerifyToken)

	}
	api = app.Group("/api/v1")
	api.Use(middlewares.Auth())
	{
		api.GET("/pong", controllers.Pong)
		api.GET("/profile", controllers.ProfileUser)
	}
}
