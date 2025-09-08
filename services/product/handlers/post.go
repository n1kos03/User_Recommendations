package handlers

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/n1kos03/User_Recommendations/common/models"
)

func (h *ProductService) POSTProduct(c *gin.Context) {
	var product models.Product

	err := c.ShouldBindBodyWithJSON(&product)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid request body",
		})
		return
	}

	err = h.DB.InsertProduct(&product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error inserting product",
		})
		slog.Error("Error inserting product", "Error: ", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Product created successfully",
		"product": product,
	})
}
