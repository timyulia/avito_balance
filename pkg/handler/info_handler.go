package handler

import (
	"balance"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// @Summary      make a report
// @Tags         info
// @Accept       json
// @Produce      json
//@Param year   path int true "year"
//@Param month   path int true "month"
// @Success      200
// @Failure      400  {object}  errorResponse
// @Failure      500  {object}  errorResponse
// @Router       /bill/info/report/{year}/{month} [get]
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

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, linkResponse{
		Link: "/bill/info/report",
	})
}

// @Summary      give a name to a service
// @Tags         info
// @Accept       json
// @Produce      json
// @Param        input body balance.Service true "enter service id and a name"
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

// @Summary      get history
// @Tags         info
// @Accept       json
// @Description  choose order by amount or date and enter in sort field
// @Produce      json
// @Param id   path int true "id"
// @Param sort   path string true "sort"
// @Success      200 {object} []balance.History
// @Failure      400  {object}  errorResponse
// @Failure      500  {object}  errorResponse
// @Router       /bill/info/history/{id}/{sort} [get]
func (h *Handler) getHistory(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}
	sort := c.Param("sort")
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid sort param")
		return
	}
	p := generatePaginationFromRequest(c)

	hists, err := h.services.GetHistory(id, sort, &p)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return

	}
	c.JSON(http.StatusOK, gin.H{
		"data": hists,
	})

}
