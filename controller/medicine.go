package controller

import (
	"fmt"
	"net/http"
	"toko_obat/repo/request"
	"toko_obat/repo/response"
	"toko_obat/repository"

	"github.com/gin-gonic/gin"
)

type MedicineController struct {
	medicineRepo *repository.MedicineRepository
}

func NewMedicineController(medicineRepo *repository.MedicineRepository) *MedicineController {
	return &MedicineController{
		medicineRepo: medicineRepo,
	}
}

func (m *MedicineController) GetAllMedicine(c *gin.Context) {
	var filter request.GetFilter
	var recordtotal int64
	var recordtotalfiltered int64

	c.Bind(&filter)

	res, recordtotal, recordtotalfiltered, err := m.medicineRepo.GetAllMedicine(filter)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusOK, response.ErrorResponseWithoutData(err))
		return
	}

	c.JSON(http.StatusOK, response.NewGetList(res, recordtotal, recordtotalfiltered))
}

func (m *MedicineController) GetMedicineByID(c *gin.Context) {
	idParam := c.Param("id")
	var id uint
	if _, err := fmt.Sscanf(idParam, "%d", &id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	medicine, err := m.medicineRepo.GetMedicineByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Medicine not found"})
		return
	}
	c.JSON(http.StatusOK, medicine)
}
