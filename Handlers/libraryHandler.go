package handlers

import (
	services "delosTest/Services"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type libraryHandler struct {
	libraryService services.LibraryService
}

func NewLibraryHandler(libraryService services.LibraryService) *libraryHandler {
	return &libraryHandler{libraryService: libraryService}
}

// @BasePath /api/v1
// PingExample godoc
// @Summary
// @Schemes
// @Description
// @Tags Library
// @Param dueDate path string true "dueDate Format DD-MM-YYYY"
// @Param submitDate path string true "submitDate Format DD-MM-YYYY"
// @Produce json
// @Success 200 {string} BookReturn
// @Router /library/bookreturn/{dueDate}/{submitDate} [get]
func (h *libraryHandler) BookReturn(c *gin.Context) {
	dueDate := c.Param("dueDate")
	submitDate := c.Param("submitDate")

	dueDateSplit := strings.Split(dueDate, "-")
	if len(dueDateSplit) != 3 {
		c.JSON(http.StatusBadRequest, "Pastikan Format dueDate adalah DD-MM-YYYY")
	}
	submitDateSplit := strings.Split(submitDate, "-")
	if len(submitDateSplit) != 3 {
		c.JSON(http.StatusBadRequest, "Pastikan Format submitDate adalah DD-MM-YYYY")
	}

	dueYear, err := strconv.Atoi(dueDateSplit[2])
	if err != nil {
		c.JSON(http.StatusBadRequest, "Tahun harus berupa angka")
	}
	dueMonth, err := strconv.Atoi(dueDateSplit[1])
	if err != nil {
		c.JSON(http.StatusBadRequest, "Bulan harus berupa angka")
	}
	dueDay, err := strconv.Atoi(dueDateSplit[0])
	if err != nil {
		c.JSON(http.StatusBadRequest, "Tanggal harus berupa angka")
	}
	submitYear, err := strconv.Atoi(submitDateSplit[2])
	if err != nil {
		c.JSON(http.StatusBadRequest, "Tahun harus berupa angka")
	}
	submitMonth, err := strconv.Atoi(submitDateSplit[1])
	if err != nil {
		c.JSON(http.StatusBadRequest, "Bulan harus berupa angka")
	}
	submitDay, err := strconv.Atoi(submitDateSplit[0])
	if err != nil {
		c.JSON(http.StatusBadRequest, "Tanggal harus berupa angka")
	}

	bookReturn := h.libraryService.BookReturn(time.Date(dueYear, time.Month(dueMonth), dueDay, 1, 1, 1, 1, time.Local), time.Date(submitYear, time.Month(submitMonth), submitDay, 1, 1, 1, 1, time.Local))

	c.JSON(http.StatusOK, bookReturn)
}
