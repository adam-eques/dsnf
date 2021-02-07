package server

import (
	"github.com/gin-gonic/gin"
	"github.com/midepeter/train-ticker/handler"
)

func StartServer() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("v1")
	{
		v1.GET("/bookings", handler.GetAllTicket)
		v1.POST("/tickets", handler.CreateTicket)
		v1.PUT("/tickets/:id", handler.UpdateTicket)
		v1.GET("/tickets/:id", handler.GetTicket)
		v1.DELETE("/tickets/:id", handler.DeleteTicket)
	}

	return r
}
