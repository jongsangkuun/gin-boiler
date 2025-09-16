package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"gin-boiler/docs"
	"gin-boiler/internal/config"
	"gin-boiler/internal/database"
	"gin-boiler/internal/router"
)

// @title           Gin Boiler API
// @version         1.0
// @description     Gin 보일러플레이트 API 서버입니다.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API 지원팀
// @contact.url    http://www.swagger.io/support
// @contact.email  support@example.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description "Bearer" 다음에 공백과 JWT 토큰을 입력하세요.
func main() {
	r := gin.Default()
	env, err := config.LoadEnv()
	if err != nil {
		panic(err)
	}

	_, err = database.Connect(env)
	if err != nil {
		panic(err)
	}

	// Swagger docs 초기화
	docs.SwaggerInfo.Title = "Gin Boiler API"
	docs.SwaggerInfo.Description = "Gin 보일러플레이트 API 서버입니다."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	router.SetupRoutes(r)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	_ = r.Run(fmt.Sprintf("0.0.0.0:%s", env.ApiConfig.Port)) // listen and serve on 0.0.0.0:8080
}
