package routes

import (
	"github.com/cnugroho11/url_shortener/controllers"
	"github.com/gin-gonic/gin"
)

type UrlRoute struct {
	urlController controllers.UrlController
}

func NewUrlRoute(urlController controllers.UrlController) UrlRoute {
	return UrlRoute{urlController}
}

func (uc *UrlRoute) UrlRoute(rg *gin.RouterGroup) {
	router := rg.Group("/url")

	router.GET("/all", uc.urlController.FetchAllUrl)
	router.POST("/add", uc.urlController.InsertUrl)
}
