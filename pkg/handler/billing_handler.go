package handler

import (
	"balance"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// @Summary      add money
// @Tags         billing
// @Accept       json
// @Produce      json
// @Param        input body balance.User true "user and amount"
// @Success      200
// @Failure      400  {object}  errorResponse
// @Failure      500  {object}  errorResponse
// @Router       /bill/add [put]
func (h *Handler) addMoney(c *gin.Context) {
	var input balance.User
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err := h.services.Billing.AddMoney(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})

}

// @Summary      reserve
// @Tags         billing
// @Accept       json
// @Produce      json
// @Param        input body balance.Order true "create a new order to reserve money"
// @Success      200
// @Failure      400  {object}  errorResponse
// @Failure      500  {object}  errorResponse
// @Router       /bill/reserve [put]
func (h *Handler) reserve(c *gin.Context) {

	var input balance.Order
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err := h.services.Reserve(input)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

// @Summary      write off
// @Tags         billing
// @Accept       json
// @Produce      json
// @Param        input body balance.Order true "specify an existing order"
// @Success      200
// @Failure      400  {object}  errorResponse
// @Failure      500  {object}  errorResponse
// @Router       /bill [put]
func (h *Handler) writeOff(c *gin.Context) {

	var input balance.Order
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err := h.services.WriteOff(input)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}

// @Summary      get balance
// @Tags         billing
// @Accept       json
// @Produce      json
//@Param id   path int true "id"
// @Success      200 {object} int
// @Failure      400  {object}  errorResponse
// @Failure      500  {object}  errorResponse
// @Router       /bill/{id} [get]
func (h *Handler) getBalance(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}
	amount, err := h.services.GetBalance(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, amount)
}

// @Summary      return money
// @Tags         billing
// @Accept       json
// @Produce      json
// @Param        input body balance.Order true "enter the order id and the user id"
// @Success      200
// @Failure      400  {object}  errorResponse
// @Failure      500  {object}  errorResponse
// @Router       /bill/return [put]
func (h *Handler) dereserve(c *gin.Context) {
	var input balance.Order
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	err := h.services.Dereserve(input)

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})
}
