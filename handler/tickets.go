package handler

import (
	"net/http"

	"github.com/midepeter/train-ticket/database/models"
	"github.com/midepeter/train-ticket/repository"
	"gitub.com/gin-gonic/gin"
	"gorm.io/gorm"
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

func (h *Handler)CreateTicket(c *gin.Context) {
	var ticket []models.Ticket

	if err := h.db.Create(&ticket).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to create ticket",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"ticket": ticket,
	})
}

func (h *Handler)UpdateTicket(c *gin.Context) {
	id := c.Params("id")

	if id := ""{
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Ticket update failed",
		})
	}
	ticket, err := h.Repo.Update(uint(id), data) 
	if err != nil {
		fmt.Println("Unable to update the ticket")
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"ticket": ticket,
	})
}

func (h *Handler)GetTicket(c *gin.Context) {
	id := c.Params("id")
	if id := "" {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to get ticket ",
		})
	}

	ticket, err := h.Repo.Get(uint(id))
	if err != nil {
		fmt.Println("ticket not found")
	}

	c.JSON(http.StatuOK, gin.H{
		"message": "success",
		"ticket": ticket,
	})
}

func (h *Handler)DeleteTicket(c *gin.Context) {
	id := c.Params("id")
	if id := "" {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed",
		})
		return
	}
	_, err := h.Repo.Delete(uint(id), ) 
	if err != nil {
		panic(err)
	}
}
