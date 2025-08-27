// Copyright Todd LLC, All rights reserved.

package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Sample data structure for demonstration
type Item struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// In-memory storage for demo purposes
var items []Item
var nextID = 1

func main() {
	r := gin.Default()

	r.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
		
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		
		c.Next()
	})

	api := r.Group("/api/v1")
	{
		api.GET("/items", getItems)
		
		api.GET("/items/:id", getItemByID)
		
		api.POST("/items", createItem)
		
		api.PUT("/items/:id", updateItem)
		
		api.DELETE("/items/:id", deleteItem)
	}

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "ok",
			"message": "Cerebro API is running",
		})
	})

	r.Run(":8080")
}

func getItems(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"data":    items,
		"count":   len(items),
		"message": "Items retrieved successfully",
	})
}

// GET /api/v1/items/:id
func getItemByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	for _, item := range items {
		if item.ID == id {
			c.JSON(http.StatusOK, gin.H{
				"data":    item,
				"message": "Item retrieved successfully",
			})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
}

// POST /api/v1/items
func createItem(c *gin.Context) {
	var newItem Item
	if err := c.ShouldBindJSON(&newItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newItem.ID = nextID
	nextID++
	items = append(items, newItem)

	c.JSON(http.StatusCreated, gin.H{
		"data":    newItem,
		"message": "Item created successfully",
	})
}

// PUT /api/v1/items/:id
func updateItem(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	var updatedItem Item
	if err := c.ShouldBindJSON(&updatedItem); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, item := range items {
		if item.ID == id {
			updatedItem.ID = id
			items[i] = updatedItem
			c.JSON(http.StatusOK, gin.H{
				"data":    updatedItem,
				"message": "Item updated successfully",
			})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
}

// DELETE /api/v1/items/:id
func deleteItem(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	for i, item := range items {
		if item.ID == id {
			items = append(items[:i], items[i+1:]...)
			c.JSON(http.StatusOK, gin.H{
				"message": "Item deleted successfully",
			})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
}
