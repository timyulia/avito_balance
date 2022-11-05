package handler

import (
	"balance/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()
	api := router.Group("/bill")
	{

		//api.GET("/", h.GetAllAccs) //all bills
		api.GET("/:id", h.getBalance)
		api.PUT("/", h.addMoney)
		api.PUT("/:id", h.writeOff)
		api.PUT("/reserve/:id", h.reserve)
		api.PUT("/:id/:service", h.dereserve)

	}
	return router
}
