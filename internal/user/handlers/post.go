package handlers

import (
	"log/slog"
	"net/http"

	"github.com/n1kos03/User_Recommendations/internal/models"

	"github.com/gin-gonic/gin"
)

func (h *UserService) POSTUser(c *gin.Context) {
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

	userKafkaEvent := models.UserMessageEvent{
		UserID:  user.ID,
		Product: user.FavoriteProduct,
	}

	if err := h.Producer.SendMessage("user_updates", []any{userKafkaEvent}); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error sending message to Kafka",
		})
		slog.Error("Error sending message to Kafka", "Error: ", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User created successfully",
	})
}
