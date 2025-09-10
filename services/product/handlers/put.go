package handlers

import (
	"log/slog"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *ProductService) PUTProduct(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error converting parametr: \"id\"",
		})
		slog.Error("Error converting parametr: \"id\"", "Error: ", err)
		return
	}

	price, err := strconv.Atoi(c.Query("price"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error converting parametr: \"price\"",
		})
		slog.Error("Error converting parametr: \"price\"", "Error: ", err)
		return
	}

	product, err := h.DB.UpdateProductPrice(id, price)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error while updating price for product",
		})
		slog.Error("Error while updating price for product", "Error: ", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"updated product": product,
	})
}