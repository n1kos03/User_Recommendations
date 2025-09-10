package handlers

import (
	"log/slog"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *ProductService) GETProduct(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error converting parametr: \"id\" ",
		})
		slog.Error("Error converting paramtr \"id\"", "Error:", err)
		return
	}

	product, err := h.DB.GetProductByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error while getting product",
		})
		slog.Error("Error while getting product.", "Error:", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"product": product,
	})
}

func (h *ProductService) GETProductByTag(c *gin.Context) {
	tags := c.QueryArray("tags")

	products, err := h.DB.GetProductsByTag(tags)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Something go wrong.",
		})
		slog.Error("Error with query parametr", "Error: ", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"products": products,
	})
}