package main

import (
	"log"

	"github.com/cnugroho11/url_shortener/controllers"
	"github.com/cnugroho11/url_shortener/routes"

	"github.com/cnugroho11/url_shortener/initializers"
	"github.com/gin-gonic/gin"
)

var (
	server *gin.Engine

	UrlController controllers.UrlController
	UrlRoute      routes.UrlRoute
)

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("Could not load env", err)
	}

	initializers.ConnectDB(&config)

	UrlController = controllers.NewUrlController(initializers.DB)
	UrlRoute = routes.NewUrlRoute(UrlController)

	server = gin.Default()
}

func main() {
	server.GET("/:url", UrlController.RedirectUrl)

	router := server.Group("/api")
	UrlRoute.UrlRoute(router)

	log.Fatal(server.Run(":" + "8000"))
}
