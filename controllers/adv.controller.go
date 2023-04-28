package controllers

import (
	"icedelect/avito-test-adv/database/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Adv struct {
	BaseController
	DB *gorm.DB
}

func NewAdv(db *gorm.DB) Adv {
	return Adv{
		DB: db,
	}
}

func (a *Adv) All(ctx *gin.Context) {
	var data []models.ShortAdv
	err, response := a.paginate(
		data,
		a.DB.Model(&models.Adv{}), // TODO: Add sort + filters
		a.preparePagination(ctx),
	)

	if err != nil {
		a.errorResponse(ctx, http.StatusBadGateway, err.Error())
		return
	}

	a.sendResponse(ctx, http.StatusOK, response)
}

func (a *Adv) Find(ctx *gin.Context) {
	var results *gorm.DB
	query := a.DB.Model(&models.Adv{})

	var adv models.ShortAdv
	results = query.First(&adv, ctx.Param("id"))

	if results.Error != nil {
		a.errorResponse(ctx, http.StatusBadGateway, results.Error.Error())
		return
	}

	a.sendResponse(ctx, http.StatusOK, BaseResponse{
		Status: "success",
		Data:   adv,
	})
}

func (a *Adv) Create(ctx *gin.Context) {
	var adv models.Adv
	if err := ctx.ShouldBindJSON(&adv); err != nil {
		a.errorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	results := a.DB.Create(&adv)
	if results.Error != nil {
		a.errorResponse(ctx, http.StatusBadGateway, results.Error.Error())
		return
	}

	a.sendResponse(ctx, http.StatusOK, BaseResponse{
		Status: "success",
		Data:   adv,
	})
}
