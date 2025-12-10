package server

import (
	"fmt"
	container "go_service/core/server/injections"
	"go_service/core/server/shared"

	"go_service/docs"
	"os"

	"github.com/MarceloPetrucio/go-scalar-api-reference"
	"github.com/gin-gonic/gin"

	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

var engine *gin.Engine

func MainServer() {

	StartEngine()

	serverContainer := container.NewContainer()

	serverContainer.Init()
	serverContainer.RegisterRoutes(engine)

	HandleDocs()

	engine.Run(":3002")
}

func StartEngine() {
	RegisterValidations()

	err := godotenv.Load()
	if err != nil {
		panic(err.Error())
	}

	engine = gin.New()

	engine.Use(shared.Cors)

	engine.Use(gin.Recovery())

}

func HandleDocs() {

	docs.SwaggerInfo.Host = os.Getenv("BACKEND_HOST")

	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	engine.GET("/openapi.json", func(ctx *gin.Context) {
		ctx.File("./docs/swagger.json")
	})
	engine.GET("/docs", func(ctx *gin.Context) {

		htmlContent, err := scalar.ApiReferenceHTML(&scalar.Options{
			SpecURL: "./docs/swagger.json",

			CustomOptions: scalar.CustomOptions{
				PageTitle: "Escritorio - DOCUMENTAÇÃO",
			},
			DarkMode: true,
			Layout:   scalar.LayoutModern,
			Theme:    scalar.ThemePurple,

			Authentication: "jwt",
		})

		if err != nil {
			fmt.Printf("%v", err)
		}

		fmt.Fprintln(ctx.Writer, htmlContent)
	})

}
