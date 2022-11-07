package handler

import (
	_ "balance/docs"
	"balance/pkg/service"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := router.Group("/bill")
	{

		//api.GET("/", h.GetAllAccs) //all bills
		api.GET("/:id", h.getBalance)
		api.PUT("/add", h.addMoney)
		api.PUT("/", h.writeOff)
		api.PUT("/reserve", h.reserve)
		api.PUT("/return", h.dereserve)

		info := api.Group("/info")
		{
			info.GET("/report/:year/:month", h.report)
			info.PUT("/specify", h.giveName)
		}
	}
	return router
}
