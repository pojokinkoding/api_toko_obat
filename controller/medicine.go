package controller

import (
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
