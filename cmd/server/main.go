package main

import (
	"fmt"
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
	_ = r.Run(fmt.Sprintf("0.0.0.0:%s", env.ApiConfig.Port)) // listen and serve on 0.0.0.0:8080
}
