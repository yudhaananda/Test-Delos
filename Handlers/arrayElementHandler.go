package handlers

import (
	services "delosTest/Services"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type arrayElementHandler struct {
	arrayElementService services.ArrayElementService
}

func NewArrayElementHandler(arrayElementService services.ArrayElementService) *arrayElementHandler {
	return &arrayElementHandler{arrayElementService: arrayElementService}
}

// @BasePath /api/v1
// PingExample godoc
// @Summary
// @Schemes
// @Description
// @Tags ArrayElement
// @Param arrLen path int true "arrLen"
// @Param arr path []int true "arr"
// @Produce json
// @Success 200 {string} SameSumElement
// @Router /arrayelement/samesumelement/{arrLen}/{arr} [get]
func (h *arrayElementHandler) SameSumElement(c *gin.Context) {
	arrLen, err := strconv.Atoi(c.Param("arrLen"))
	if err != nil {
		c.JSON(http.StatusBadRequest, "arrLen harus berupa angka")
	}
	arr := strings.Split(c.Param("arr"), ",")
	sameSumElement, err := h.arrayElementService.SameSumElement(arrLen, arr)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}
	c.JSON(http.StatusOK, sameSumElement)
}
