package postgres

import (
	"context"
	m "37hw/models"
	"log"
	"strings"

	"github.com/jmoiron/sqlx"
)

type AlbumRepo struct {
	DB *sqlx.DB
}

func NewAlbumsrepo(db *sqlx.DB) *AlbumRepo {
	return &AlbumRepo{
		DB: db,
	}
}

func (a *AlbumRepo) CreateAlbum(ctx context.Context, alb m.Album) error {
	tx, err := a.DB.Begin()
	if err != nil {
		log.Fatal("Error beginning transaction: ", err)
	}

	query := `
		INSERT INTO albums (
			title, 
			artist, 
			price) 
		VALUES(	$1, $2, $3)
		RETURNING
			id,         
			title,       
			artist,      
			price
	`

	_, err = tx.ExecContext(ctx, query, strings.ToLower(alb.Title), strings.ToLower(alb.Artist), alb.Price)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

func (a *AlbumRepo) GetAlbumsById(ctx context.Context, id string) (m.Album, error) {
	query := `
	SELECT * FROM albums WHERE id = $1
	`
	row := a.DB.QueryRowContext(ctx, query, id)

	album := m.Album{}

	err := row.Scan(&album.Id, &album.Title, &album.Artist, &album.Price)
	if err != nil {
		return album, err
	}

	return album, nil
}

func (a *AlbumRepo) GetAlbums(ctx context.Context) (albums []m.Album, err error) {
	query := `
	SELECT * FROM albums
	`
	row, err := a.DB.QueryContext(ctx, query)
	if err != nil {
		return albums, err
	}

	for row.Next() {
		album := m.Album{}
		err := row.Scan(&album.Id, &album.Title, &album.Artist, &album.Price)
		if err != nil {
			log.Fatal("Error scanning sql query: ", err)
		}
		albums = append(albums, album)
	}
	return albums, nil
}

func (a *AlbumRepo) UpdateAlbumById(ctx context.Context, alb m.Album, id string) (m.Album, error) {
	album := m.Album{}

	tx, err := a.DB.Begin()
	if err != nil {
		return album, err
	}

	query := `
		UPDATE albums
		SET title = $1, artist = $2, price = $3
		WHERE id = $4
		RETURNING
			id,         
			title,       
			artist,      
			price
	`

	row := tx.QueryRowContext(ctx, query, alb.Title, alb.Artist, alb.Price, id)

	err = row.Scan(&album.Id, &album.Title, &album.Artist, &album.Price)
	if err != nil {
		tx.Rollback()
		return album, err
	}

	err = tx.Commit()
	if err != nil {
		return album, err
	}

	return album, nil
}

func (a *AlbumRepo) GetAlbumsByTitle(ctx context.Context, title string) (albums []m.Album, err error) {
	query := `
        SELECT * FROM albums WHERE title = $1
    `
	rows, err := a.DB.QueryContext(ctx, query, title)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var album m.Album
		err := rows.Scan(&album.Id, &album.Title, &album.Artist, &album.Price)
		if err != nil {
			return nil, err
		}
		albums = append(albums, album)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return albums, nil
}

func (a *AlbumRepo) GetAlbumsByArtist(ctx context.Context, artist string) (albums []m.Album, err error) {
	query := `
        SELECT * FROM albums WHERE artist = $1
    `
	rows, err := a.DB.QueryContext(ctx, query, artist)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var album m.Album
		err := rows.Scan(&album.Id, &album.Title, &album.Artist, &album.Price)
		if err != nil {
			return nil, err
		}
		albums = append(albums, album)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return albums, nil
}

func (a *AlbumRepo) GetAlbumsByPrice(ctx context.Context, price float64) (albums []m.Album, err error) {
	query := `
        SELECT * FROM albums WHERE price = $1
    `
	rows, err := a.DB.QueryContext(ctx, query, price)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var album m.Album
		err := rows.Scan(&album.Id, &album.Title, &album.Artist, &album.Price)
		if err != nil {
			return nil, err
		}
		albums = append(albums, album)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return albums, nil
}

func (a *AlbumRepo) DeletAlbumsById(ctx context.Context, id string) error {
	query := `
	DELETE FROM albums WHERE id = $1;
	`
	_, err := a.DB.ExecContext(ctx, query, id)

	if err != nil {
		return err
	}

	return nil
}
