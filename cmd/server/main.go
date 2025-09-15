package main

import (
	"gin-boiler/internal/config"
	"gin-boiler/internal/database"
	"gin-boiler/internal/router"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	env, err := config.LoadEnv()
	if err != nil {
		panic(err)
	}
	log.Println(env.DbConfig)

	_, err = database.Connect(env)
	if err != nil {
		panic(err)
	}

	router.SetupRoutes(r)

	_ = r.Run("0.0.0.0:8080") // listen and serve on 0.0.0.0:8080
}
