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

/* ==== Front requests ==== */
func (a *App) GetAlbums(filters []Param) []Album {
	albums, err := dao.getAlbums(filters...)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return albums
}

func (a *App) GetArtists(filters []Param) []Artist {
	artist, err := dao.getArtists(filters...)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return artist
}

type Person struct {
	Name    string   `json:"name"`
	Age     uint8    `json:"age"`
	Address *Address `json:"address"`
}

type Address struct {
	Street   string `json:"street"`
	Postcode string `json:"postcode"`
}

func (a *App) Greet(p Person) string {
	return fmt.Sprintf("Hello %s (Age: %d)!", p.Name, p.Age)
}
