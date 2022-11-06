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
		api.PUT("/add", h.addMoney)
		api.PUT("/", h.writeOff)
		api.PUT("/reserve", h.reserve)
		api.PUT("/return", h.dereserve)

	}
	return router
}
