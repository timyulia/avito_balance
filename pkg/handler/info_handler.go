package handler

import (
	"balance"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

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
	//
	//router := gin.Default()
	//router.Static("/image", "./path-to-image-dir")

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

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
