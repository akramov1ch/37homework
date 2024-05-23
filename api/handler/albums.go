package handler

import (
	m "37hw/models"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *handler) CreateAlbum(c *gin.Context) {
	newAlbum := m.Album{}

	if err := c.BindJSON(&newAlbum); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	newAlbum.Id = uuid.New().String()

	err := h.storage.Album().CreateAlbum(c.Request.Context(), newAlbum)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	c.IndentedJSON(http.StatusCreated, gin.H{"message": "Successfully created !!!"})
}

func (h *handler) UpdateAlbums(c *gin.Context) {
	id := c.Param("id")
	newAlbum := m.Album{}

	if err := c.BindJSON(&newAlbum); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	album, err := h.storage.Album().UpdateAlbumById(c.Request.Context(), newAlbum, id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	c.IndentedJSON(http.StatusOK, album)
}

func (h *handler) GetAlbumByID(c *gin.Context) {
	id := c.Param("id")

	album, err := h.storage.Album().GetAlbumsById(c.Request.Context(), id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	c.IndentedJSON(http.StatusOK, album)
}

func (h *handler) GetAlbums(c *gin.Context) {
	album, err := h.storage.Album().GetAlbums(c.Request.Context())
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	c.IndentedJSON(http.StatusOK, album)
}

func (h *handler) GetAlbumByTitle(c *gin.Context) {
	title := strings.ToLower(c.Param("title"))

	album, err := h.storage.Album().GetAlbumsByTitle(c.Request.Context(), title)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	c.IndentedJSON(http.StatusOK, album)
}

func (h *handler) GetAlbumByArtist(c *gin.Context) {
	artist := strings.ToLower(c.Param("artist"))

	album, err := h.storage.Album().GetAlbumsByArtist(c.Request.Context(), artist)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	c.IndentedJSON(http.StatusOK, album)
}

func (h *handler) GetAlbumByPrice(c *gin.Context) {
	price := c.Param("price")

	cprice, err := strconv.ParseFloat(price, 64)
	if err != nil {
		log.Println("Error converting string to float64: ", err)
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	album, err := h.storage.Album().GetAlbumsByPrice(c.Request.Context(), cprice)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	c.IndentedJSON(http.StatusOK, album)
}

func (h *handler) DeleteAlbumById(c *gin.Context) {
	id := c.Param("id")

	err := h.storage.Album().DeletAlbumsById(c.Request.Context(), id)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Value deleted successfully!!!"})
}
