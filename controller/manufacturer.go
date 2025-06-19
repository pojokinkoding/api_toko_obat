package controller

import (
	"fmt"
	"net/http"
	"toko_obat/repo/request"
	"toko_obat/repo/response"
	"toko_obat/repository"

	"github.com/gin-gonic/gin"
)

type ManufacturerController struct {
	manufacturerRepo *repository.ManufacturerRepository
}

func NewManufacturerController(manufacturerRepo *repository.ManufacturerRepository) *ManufacturerController {
	return &ManufacturerController{
		manufacturerRepo: manufacturerRepo,
	}
}

func (c *ManufacturerController) GetAllManufacturer(ctx *gin.Context) {
	var filter request.GetFilter
	var recordtotal int64
	var recordtotalfiltered int64

	ctx.Bind(&filter)

	res, recordtotal, recordtotalfiltered, err := c.manufacturerRepo.GetAllManufacturer(filter)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusOK, response.ErrorResponseWithoutData(err))
		return
	}

	ctx.JSON(http.StatusOK, response.NewGetList(res, recordtotal, recordtotalfiltered))
}

func (c *ManufacturerController) GetManufacturerByID(ctx *gin.Context) {
	idParam := ctx.Param("id")
	var id uint
	if _, err := fmt.Sscanf(idParam, "%d", &id); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	manufacturer, err := c.manufacturerRepo.GetManufacturerByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Manufacturer not found"})
		return
	}
	ctx.JSON(http.StatusOK, manufacturer)
}
