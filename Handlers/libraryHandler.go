package handlers

import (
	services "delosTest/Services"
	"net/http"

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

	bookReturn, err := h.libraryService.BookReturn(dueDate, submitDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
	}
	c.JSON(http.StatusOK, bookReturn)
}
