package handlers

import (
	"User_Recommendations/internal/models"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *UserHandler) POSTUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindBodyWithJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request body",
		})
		slog.Error("Invalid request body", "Error: ", err)
		return
	}

	err := h.DB.InsertUser(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error inserting user",
		})
		slog.Error("Error inserting user", "Error: ", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User created successfully",
	})
}