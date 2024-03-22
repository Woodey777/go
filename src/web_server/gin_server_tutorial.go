package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func getAlbums(c *gin.Context) {
	// get data from db..

	c.JSON(http.StatusOK, albums)
}

func addAlbums(c *gin.Context) {
	var alb album

	if err := c.BindJSON(&alb); err != nil {
		fmt.Printf("binding JSON: %w", err)
		return
	}

	albums = append(albums, alb)
	c.JSON(http.StatusCreated, alb)
}

func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	for _, alb := range albums {
		if alb.ID == id {
			c.JSON(http.StatusOK, alb)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.POST("/albums", addAlbums)

	router.Run("localhost:55000")
}
