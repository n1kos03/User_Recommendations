package handlers

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *UserService) PUTUser(c *gin.Context) {
	user, err := h.DB.UpdateUser(c.Param("id"), c.Query("favorite_product"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error updating user",
		})
		slog.Error("Error updating user", "Error: ", err)
		return
	}

	if err := h.Producer.SendMessage("user_updates", []any{user}); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error sending message to Kafka",
		})
		slog.Error("Error sending message to Kafka", "Error: ", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User updated successfully",
	})
}
