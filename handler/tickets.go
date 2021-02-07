package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/midepeter/train-ticket/database/models"
	"github.com/midepeter/train-ticket/utils/utils"
)

func (h *Handler) GetAllTickets(c *gin.Context) {
	var tickets []models.Ticket

	if err := h.db.Find(&tickets).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to load tickets",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"tickets": tickets,
	})
}

func (h *Handler) CreateTicket(c *gin.Context) {
	var ticket []models.Ticket

	if err := h.db.Create(&ticket).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to create ticket",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"ticket":  ticket,
	})
}

func (h *Handler) UpdateTicket(c *gin.Context) {
	var ticket *models.Ticket
	id := c.Query("id")
	if id == "" {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Ticket update failed",
		})
	}

	id64 := utils.ConvertToUint64(id)
	err := h.Repo.UpdateBooking(uint(id64), ticket)
	if err != nil {
		fmt.Println("Unable to update the ticket")
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"ticket":  ticket,
	})
}

func (h *Handler) GetTicket(c *gin.Context) {
	var ticket *models.Ticket

	id := c.Query("id")
	if id == "" {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to get ticket ",
		})
	}

	id64 := utils.ConvertToUint64(id)
	err := h.Repo.GetBooking(uint(id64), ticket)
	if err != nil {
		fmt.Println("ticket not found")
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"ticket":  ticket,
	})
}

func (h *Handler) DeleteTicket(c *gin.Context) {
	var ticket *models.Ticket
	id := c.Query("id")
	if id == "" {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed",
		})
		return
	}

	id64 := utils.ConvertToUint64(id)
	err := h.Repo.DeleteBooking(uint(id64), ticket)
	if err != nil {
		panic(err)
	}
}
