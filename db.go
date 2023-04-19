package main

import (
	"database/sql"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

type Dao struct {
	path string
}

// newDao creates a new Dao instance
func newDao(path string) (*Dao, error) {
	dao := &Dao{path: path}
	err := dao.initDB(path)
	return dao, err
}

// initDB creates the database file and tables if they don't exist
func (e *Dao) initDB(path string) error {
	// check if db file exists and create it if not
	if _, err := os.Stat(path); err != nil {
		_, err := os.Create(path)
		if err != nil {
			return err
		}
	}

	// then check if the table exists and create it if not
	db, err := e.openDB()
	if err != nil {
		return err
	}

	//Create a table for artists
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS artists (
			person_id INTEGER PRIMARY KEY,
			first_name TEXT NOT NULL,
			last_name TEXT
		)`)
	if err != nil {
		return err
	}
	//Create a table for albums
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS albums (
			album_id INTEGER PRIMARY KEY,
			artist_id INTEGER,
			album_name TEXT NOT NULL,
			album_image TEXT, 
			FOREIGN KEY (artist_id) REFERENCES artists (person_id)
		)`)
	if err != nil {
		return err
	}
	//Create a table for songs
	// I've add the foregin key to an artist and album because I want to be able to get all the songs from an album and all the songs from an artist easily
	// and i want to later be able to have a song in multiple albums
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS songs (
			song_id INTEGER PRIMARY KEY,
			album_id INTEGER,
			artist_id INTEGER,
			song_name TEXT NOT NULL,
			song_path TEXT NOT NULL,
			FOREIGN KEY (album_id) REFERENCES albums (album_id)
			FOREIGN KEY (artist_id) REFERENCES artists (person_id)
		)`)

	if err != nil {
		return err
	}
	defer db.Close()
	return nil
}

// openDB opens the database file
func (e *Dao) openDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", e.path)
	if err != nil {
		return nil, err
	}
	return db, nil
}

/* ==== DB TYPES ==== */
type Artist struct {
	ID        *int64 `json:"id" db:"person_id"`
	FirstName string `json:"firstName" db:"first_name"`
	LastName  string `json:"lastName" db:"last_name"`
}
type Album struct {
	ID        *int64 `json:"id" db:"album_id"`
	Artist    int64  `json:"artist" db:"artist_id"`
	Name      string `json:"name" db:"album_name"`
	ImagePath string `json:"imagePath" db:"album_image"`
}
type Song struct {
	ID     *int64 `json:"id" db:"song_id"`
	Album  int64  `json:"album" db:"album_id"`
	Artist int64  `json:"artist" db:"artist_id"`
	Name   string `json:"name" db:"song_name"`
	Path   string `json:"path" db:"song_path"`
}

/* ==== DB QUERIES ==== */

/* --- INSERT --- */
// Insert a new song into the database and return the id of the inserted song
func (e *Dao) insertSong(song Song, album *Album, artist *Artist) (int64, error) {
	var err error

	/*
		// if artist is not nil then insert the artist
		if album != nil {
			song.Album, err = e.insertAlbum(*album, *artist)
			if err != nil {
				return 0, err
			}
		}

		// if artist is not nil then insert the artist
		if artist != nil {
			song.Artist, err = e.insertArtist(*artist)
			if err != nil {
				return 0, err
			}
		}*/

	//Insert a new song
	db, err := e.openDB()
	if err != nil {
		return 0, err
	}
	res, err := db.Exec(`INSERT INTO songs (album_id, artist_id, song_name, song_path) VALUES (?, ?, ?, ?)`, song.Album, song.Artist, song.Name, song.Path)
	if err != nil {
		return 0, err
	}
	// get the id of the inserted song
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

// Insert a new album into the database and return the id of the inserted album
func (e *Dao) insertAlbum(album Album, artist Artist) (int64, error) {
	var err error

	/*
		// if artist is not nil then insert the artist
		if artist.ID == nil {
			album.Artist, err = e.insertArtist(artist)
			if err != nil {
				return 0, err
			}
		}
	*/

	//Insert a new album
	db, err := e.openDB()
	if err != nil {
		return 0, err
	}
	res, err := db.Exec(`INSERT INTO albums (artist_id, album_name, album_image) VALUES (?, ?, ?)`, album.Artist, album.Name, album.ImagePath)
	if err != nil {
		return 0, err
	}
	// get the id of the inserted album
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

// Insert a new artist into the database and return the id of the inserted artist
func (e *Dao) insertArtist(artist Artist) (int64, error) {
	//Insert a new artist
	db, err := e.openDB()
	if err != nil {
		return 0, err
	}
	res, err := db.Exec(`INSERT INTO artists (first_name, last_name) VALUES (?, ?)`, artist.FirstName, artist.LastName)
	if err != nil {
		return 0, err
	}
	// get the id of the inserted artist
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

/* --- SELECT --- */

type Param struct {
	Field string `json:"key"`
	Value any    `json:"value"`
}

func formatWhereClause(filters []Param) (string, []any) {
	var whereClause string
	var values []any
	for i, filter := range filters {
		//Build the where clause
		if i == 0 {
			whereClause += " WHERE "
		} else {
			whereClause += " AND "
		}
		whereClause += filter.Field + " = ?"
		//Build the values array
		values = append(values, filter.Value)
	}
	return whereClause, values
}

//TODO refactor this function to be more generic and implement a heritance system

//TODO add the option to get a specific number of songs and to get songs from a specific index

// Get songs from the database
// TODO add filters fonctionality for songs
func (e *Dao) getSongs(filters ...Param) ([]Song, error) {
	db, err := e.openDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	query := "SELECT * FROM songs"
	where, values := formatWhereClause(filters)
	query += where

	rows, err := db.Query(query, values...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var songs []Song
	for rows.Next() {
		var song Song
		err := rows.Scan(&song.ID, &song.Album, &song.Artist, &song.Name, &song.Path)
		if err != nil {
			return nil, err
		}
		songs = append(songs, song)
	}
	return songs, nil
}

// Get all the albums from the database
// TODO add filters fonctionality for albums
func (e *Dao) getAlbums(filters ...Param) ([]Album, error) {
	db, err := e.openDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	query := "SELECT * FROM albums"
	where, values := formatWhereClause(filters)
	query += where

	rows, err := db.Query(query, values...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var albums []Album
	for rows.Next() {
		var album Album
		err := rows.Scan(&album.ID, &album.Artist, &album.Name, &album.ImagePath)
		if err != nil {
			return nil, err
		}
		albums = append(albums, album)
	}
	return albums, nil
}

// Get all the artists from the database
// TODO add filters fonctionality for artists
func (e *Dao) getArtists(filters ...Param) ([]Artist, error) {
	db, err := e.openDB()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	query := "SELECT * FROM artists"
	where, values := formatWhereClause(filters)
	query += where

	rows, err := db.Query(query, values...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var artists []Artist
	for rows.Next() {
		var artist Artist
		err := rows.Scan(&artist.ID, &artist.FirstName, &artist.LastName)
		if err != nil {
			return nil, err
		}
		artists = append(artists, artist)
	}
	return artists, nil
}
