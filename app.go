package main

import (
	"context"
	"fmt"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

/* ==== Front get ==== */

// Return the songs with the filters
func (a *App) GetSongs(filters []Param) []Song {
	songs, err := dao.getSongs(filters...)
	if err != nil {
		fmt.Println(err)
		songs = make([]Song, 0)
	}
	return songs
}

// Return the songs with the filters
func (a *App) GetAlbums(filters []Param) []Album {
	albums, err := dao.getAlbums(filters...)
	if err != nil {
		fmt.Println(err)
		albums = make([]Album, 0)
	}
	return albums
}

// Return the songs with the filters
func (a *App) GetArtists(filters []Param) []Artist {
	artist, err := dao.getArtists(filters...)
	if err != nil {
		fmt.Println(err)
		artist = make([]Artist, 0)
	}
	return artist
}

/* ==== Front post ==== */

func (a *App) PostArtists(artists Artist) bool {
	_, err := dao.insertArtist(artists)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

func (a *App) PostAlbums(albums Album) bool {
	_, err := dao.insertAlbum(albums, Artist{ID: &albums.Artist})
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

func (a *App) PostSongs(songs Song) bool {
	_, err := dao.insertSong(songs, &Album{ID: &songs.Album}, &Artist{ID: &songs.Artist})
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}
