package handlers

import (
	services "delosTest/Services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type candyHandler struct {
	candyService services.CandyService
}

func NewCandyHandler(candyService services.CandyService) *candyHandler {
	return &candyHandler{candyService: candyService}
}

// @BasePath /api/v1
// PingExample godoc
// @Summary
// @Schemes
// @Description
// @Tags Candy
// @Param student path int true "student"
// @Param candies path int true "candies"
// @Param firstStudent path int true "firstStudent"
// @Produce json
// @Success 200 {string} WhoGetSourCandy
// @Router /candy/whogetsourcandy/{student}/{candies}/{firstStudent} [get]
func (h *candyHandler) WhoGetSourCandy(c *gin.Context) {
	student, err := strconv.Atoi(c.Param("student"))
	if err != nil {
		c.JSON(http.StatusBadRequest, "student harus berupa angka")
	}
	candies, err := strconv.Atoi(c.Param("candies"))
	if err != nil {
		c.JSON(http.StatusBadRequest, "candies harus berupa angka")
	}
	firstStudent, err := strconv.Atoi(c.Param("firstStudent"))
	if err != nil {
		c.JSON(http.StatusBadRequest, "firstStudent harus berupa angka")
	}

	whoGetSourCandy := h.candyService.WhoGetSourCandy(student, candies, firstStudent)
	c.JSON(http.StatusOK, whoGetSourCandy)

}
