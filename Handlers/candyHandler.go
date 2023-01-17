package handlers

import (
	services "delosTest/Services"
	"net/http"

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
	student := c.Param("student")
	candies := c.Param("candies")
	firstStudent := c.Param("firstStudent")

	whoGetSourCandy, err := h.candyService.WhoGetSourCandy(student, candies, firstStudent)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}
	c.JSON(http.StatusOK, whoGetSourCandy)

}
