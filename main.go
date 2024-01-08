package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"hackfest/controller"
	"hackfest/database"
	"hackfest/middleware"
	"hackfest/utils"
	"log"
	"net/http"
	"os"
)

func main() {
	fmt.Println("Hello, World!")

	if err := godotenv.Load(); err != nil {
		log.Fatal(err.Error())
	}

	databaseConfig, err := database.NewDatabase()
	if err != nil {
		log.Fatal(err.Error())
	}

	db, err := database.MakeDatabaseConnection(databaseConfig)
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Printf("Database Connected: %s\n", db)

	r := gin.Default()

	r.Use(middleware.CORS())

	r.GET("/", func(c *gin.Context) {
		utils.HttpRespSuccess(c, http.StatusOK, os.Getenv("ENV"), nil)
		return
	})

	controller.Auth(db, r)
	controller.Profile(db, r)

	if err := r.Run(fmt.Sprintf(":%s", os.Getenv("PORT"))); err != nil {
		panic(err.Error())
	}
}
