package handlers

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *UserHandler) GETUser(c *gin.Context) {
	users, err := h.DB.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error getting users",
		})
		slog.Error("Error getting users", "Error: ", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}
