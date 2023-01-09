package controllers

import (
	"net/http"
	"time"

	"github.com/cnugroho11/url_shortener/models"
	"github.com/cnugroho11/url_shortener/utils"
	"github.com/gin-gonic/gin"
	"github.com/lithammer/shortuuid/v4"
	"gorm.io/gorm"
)

type UrlController struct {
	DB *gorm.DB
}

func NewUrlController(DB *gorm.DB) UrlController {
	return UrlController{DB}
}

func (uc *UrlController) FetchAllUrl(ctx *gin.Context) {
	var urls []models.Url

	pagination := utils.Pagination(ctx)
	offset := (pagination.Page - 1) * pagination.Limit

	getUrls := uc.DB.Limit(pagination.Limit).Offset(offset).Order(pagination.Sort).Find(&urls)
	if getUrls.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": "Failed to get movies",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"pagination": gin.H{
			"page":  pagination.Page,
			"limit": pagination.Limit,
			"sort":  pagination.Sort,
		},
		"data": gin.H{
			"urls": urls,
		},
	})
}

func (uc *UrlController) InsertUrl(ctx *gin.Context) {
	var payload *models.UrlInput
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": err.Error(),
		})
		return
	}

	generatedUrl := shortuuid.New()
	newUrl := models.Url{
		RealUrl:      payload.RealUrl,
		ShortenedUrl: generatedUrl,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	insertDB := uc.DB.Create(&newUrl)
	if insertDB.Error != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{
			"status":  "fail",
			"message": "error insert data",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Short url generated",
	})
}

func (uc *UrlController) RedirectUrl(ctx *gin.Context) {
	shorten := ctx.Param("url")

	var url models.Url
	getUrl := uc.DB.Where("shortened_url = ?", shorten).First(&url)
	if getUrl.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"status":  "fail",
			"message": "url not found",
		})
		return
	}

	ctx.Redirect(http.StatusPermanentRedirect, url.RealUrl)
}
