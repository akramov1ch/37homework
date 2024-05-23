package repo

import (
	"context"
	m "37hw/models"
)

type AlbumsStorageI interface {
	CreateAlbum(ctx context.Context, alb m.Album) error
	GetAlbumsById(ctx context.Context, id string) (m.Album, error)
	GetAlbums(ctx context.Context) (albums []m.Album, err error)
	UpdateAlbumById(ctx context.Context, alb m.Album, id string) (m.Album, error)
	GetAlbumsByTitle(ctx context.Context, title string) (albums []m.Album, err error)
	GetAlbumsByArtist(ctx context.Context, artist string) (albums []m.Album, err error)
	GetAlbumsByPrice(ctx context.Context, price float64) (albums []m.Album, err error)
	DeletAlbumsById(ctx context.Context, id string) error
}
