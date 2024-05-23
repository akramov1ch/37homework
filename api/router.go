package api

import (
	h "37hw/api/handler"
	"37hw/storage"

	"github.com/gin-gonic/gin"
)

type Option struct {
	Storage storage.IStorage
}

func New(option Option) *gin.Engine {
	router := gin.Default()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	handler := h.New(&h.HandlerConfig{
		Storage: option.Storage,
	})

	album := router.Group("/album")
	album.POST("", handler.CreateAlbum)
	album.GET("", handler.GetAlbums)
	album.GET("/:id", handler.GetAlbumByID)
	album.PUT("/:id", handler.UpdateAlbums)
	album.GET("/title/:title", handler.GetAlbumByTitle)
	album.GET("/artist/:artist", handler.GetAlbumByArtist)
	album.GET("/price/:price", handler.GetAlbumByPrice)
	album.DELETE("/:id", handler.DeleteAlbumById)

	return router
}
