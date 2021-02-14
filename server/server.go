package server

import (
	"github.com/gin-gonic/gin"
	"github.com/midepeter/train-ticket/handler"
	"gorm.io/gorm"
)

<<<<<<< HEAD
func StartServer() *gin.Engine {
=======
func StartServer(db *gorm.DB) *gin.Engine {
>>>>>>> c3d406c1586667a68bca168dd8d85b5bc4e01b07
	h := handler.NewHandler(db)

	r := gin.Default()

	r.GET("/bookings", h.GetAllTickets)
	r.POST("/tickets", h.CreateTicket)
	r.PUT("/tickets/:id", h.UpdateTicket)
	r.GET("/tickets/:id", h.GetTicket)
	r.DELETE("/tickets/:id", h.DeleteTicket)

	return r
}
