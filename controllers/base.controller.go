package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type BaseController struct{}

type BaseResponse struct {
	Status     string             `json:"status"`
	Message    string             `json:"message,omitempty"`
	Data       interface{}        `json:"data,omitempty"`
	Pagination *PaginationOptions `json:"pagination,omitempty"`
}

type PaginationOptions struct {
	Page   int   `json:"page"`
	Limit  int   `json:"limit"`
	Offset int   `json:"offset"`
	Total  int64 `json:"total"`
}

func (b *BaseController) errorResponse(ctx *gin.Context, status int, message string) {
	ctx.JSON(status, gin.H{"status": "error", "message": message})
}

func (b *BaseController) sendResponse(ctx *gin.Context, status int, response BaseResponse) {
	// b, _ = json.Marshal(response)
	ctx.JSON(status, response)
}

func (b *BaseController) preparePagination(ctx *gin.Context) PaginationOptions {
	page, err := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	if err != nil {
		page = 1
	}

	limit, err := strconv.Atoi(ctx.DefaultQuery("limit", "10"))
	if err != nil {
		limit = 10
	}

	intPage := page
	intLimit := limit
	offset := (intPage - 1) * intLimit

	return PaginationOptions{
		Page:   intPage,
		Limit:  intLimit,
		Offset: offset,
	}
}

func (b *BaseController) paginate(data interface{}, query *gorm.DB, options PaginationOptions) (error, BaseResponse) {
	result := query.Limit(options.Limit).Offset(options.Offset).Find(&data) // TODO: Use Scope
	if result.Error != nil {
		return result.Error, BaseResponse{}
	}

	query.Count(&options.Total)

	return nil, BaseResponse{
		Data:       data,
		Status:     "success",
		Pagination: &options,
	}
}
