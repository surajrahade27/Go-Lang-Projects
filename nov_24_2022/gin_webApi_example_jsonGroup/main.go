package main

import (
	"encoding/json"
	"io"

	"net/http"
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

// album represents data about a record album.
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {
	val, _ := c.Get("jsonalbums") // val is a type of interface{}
	albums := val.([]album)       // type casting
	c.IndentedJSON(http.StatusOK, albums)
}

// postAlbums adds an album from JSON received in the request body.
func postAlbums(c *gin.Context) {
	var newAlbum album
	val, _ := c.Get("jsonalbums") // val is a type of interface{}
	albums := val.([]album)       // type casting
	// Call BindJSON to bind the received JSON to
	// newAlbum.
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Add the new album to the slice.
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func getAlbumByID(c *gin.Context) {
	id := c.Param("id")
	val, _ := c.Get("jsonAlbums") // val is a type of interface{}
	albums := val.([]album)       // type casting
	// Loop through the list of albums, looking for
	// an album whose ID value matches the parameter.
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

// JsonLoadHandler : will read json data from a file database and inject the same in the context
//
// For demo purposes we are just querying a small json file database
// Queries and filtering can follow in downstream in handlers
func JsonLoadHandler(c *gin.Context) {
	data := []album{}
	file, err := os.Open("MOCKDATA.json")
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	}
	byt, err := io.ReadAll(file)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	}
	if json.Unmarshal(byt, &data) != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	}
	c.Set("jsonAlbums", data)
	log.Debugf("We have about %d Albums loaded", len(data))
}

func main() {

	router := gin.Default()
	// router.GET("/albums", getAlbums)
	// router.GET("/albums/:id", getAlbumByID)
	// router.POST("/albums", postAlbums)

	albums := router.Group("/albums", JsonLoadHandler)
	albums.GET("/", func(c *gin.Context) {
		val, _ := c.Get("jsonalbums") // val is a type of interface{}
		data := val.([]album)         // type casting
		c.AbortWithStatusJSON(http.StatusOK, gin.H{
			"total": len(data),
		})
	})
	albums.GET("/:id", getAlbumByID)
	router.Run("localhost:8080")
}
