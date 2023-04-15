package main

import (
	"errors"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

/**
The storage layout is as follows:
./files/									// this is where the files are stored
	|-> artists_id/ 						// contain all the files of the artist
		|-> artist_image.jpg				// The artist image
		|-> albums/							// Contains all the albums of the artist
			|-> album_id/					// Contains all the files of the album
				|-> album_image.jpg			// The album image
				|-> songs/					// Contains all the songs of the album
					|-> songs_id.mp3		// The song file
*/

// Config

var Filespath string = "./storage/files"

type FileStorage struct {
	path string
}

func NewFileStorage() (*FileStorage, error) {
	err := initFileStorage()
	if err != nil {
		return nil, err
	}
	return &FileStorage{
		path: Filespath,
	}, nil
}

func initFileStorage() error {
	// check if files folder exists and create it if not
	if _, err := os.Stat(Filespath); errors.Is(err, os.ErrNotExist) {
		err := os.Mkdir(Filespath, os.ModePerm)
		if err != nil {
			log.Println(err)
		}
	}
	return nil
}

// File Handling
