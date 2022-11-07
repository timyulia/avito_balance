package handler

import (
	"balance"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// @Summary      make a report
// @Tags         billing
// @Accept       json
// @Produce      json
// @Param        input body balance.Report true "choose a year and a month"
// @Success      200
// @Failure      400  {object}  errorResponse
// @Failure      500  {object}  errorResponse
// @Router       /bill/info/report/:year/:month [get]
func (h *Handler) report(c *gin.Context) {
	year, err := strconv.Atoi(c.Param("year"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid year param")
		return
	}
	month, err := strconv.Atoi(c.Param("month"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid month param")
		return
	}
	err = h.services.MakeReport(year, month)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	file, err := c.FormFile("report.csv")
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	// Upload the file to specific dst.
	dst := "csv"
	err = c.SaveUploadedFile(file, dst)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))

	//c.JSON(http.StatusOK, statusResponse{
	//	Status: "ok",
	//})
}

// @Summary      give name to a service
// @Tags         billing
// @Accept       json
// @Produce      json
// @Param        input body balance.Report true "enter the service id and a name"
// @Success      200
// @Failure      400  {object}  errorResponse
// @Failure      500  {object}  errorResponse
// @Router       /bill/info/specify [put]
func (h *Handler) giveName(c *gin.Context) {
	var input balance.Report
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	err := h.services.GiveName(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}
