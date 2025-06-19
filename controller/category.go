package controller

import (
	"fmt"
	"net/http"
	"toko_obat/repo/request"
	"toko_obat/repo/response"
	"toko_obat/repository"

	"github.com/gin-gonic/gin"
)

type CategoryController struct {
	categoryRepo *repository.CategoryRepository
}

func NewCategoryController(categoryRepo *repository.CategoryRepository) *CategoryController {
	return &CategoryController{
		categoryRepo: categoryRepo,
	}
}

func (c *CategoryController) GetAllCategory(ctx *gin.Context) {
	var filter request.GetFilter
	var recordtotal int64
	var recordtotalfiltered int64

	ctx.Bind(&filter)

	res, recordtotal, recordtotalfiltered, err := c.categoryRepo.GetAllCategory(filter)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusOK, response.ErrorResponseWithoutData(err))
		return
	}

	ctx.JSON(http.StatusOK, response.NewGetList(res, recordtotal, recordtotalfiltered))
}

func (c *CategoryController) GetCategoryByID(ctx *gin.Context) {
	idParam := ctx.Param("id")
	var id uint
	if _, err := fmt.Sscanf(idParam, "%d", &id); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	category, err := c.categoryRepo.GetCategoryByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		return
	}
	ctx.JSON(http.StatusOK, category)
}

func (c *CategoryController) CreateCategory(ctx *gin.Context) {
	var req request.CreateCategoryRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, response.ErrorCustomRepsonse("Invalid request: "+err.Error()))
		return
	}
	uid := ctx.GetString("uid")
	if err := c.categoryRepo.CreateCategory(req, uid); err != nil {
		ctx.JSON(http.StatusInternalServerError, response.ErrorCustomRepsonse("Failed to create category: "+err.Error()))
		return
	}
	ctx.JSON(http.StatusOK, response.SuccessResponseWithMessage(nil, "Category created successfully"))
}
