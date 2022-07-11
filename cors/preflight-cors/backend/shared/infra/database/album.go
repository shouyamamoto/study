package database

import (
	"database/sql"
	"github/shouyamamoto/study/db"
	"log"
	"math/rand"
	"net/http"
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
)

func dbConn() (db *sql.DB) {
	cfg := mysql.Config{
		User:   os.Getenv("DBUSER"),
		Passwd: os.Getenv("DBPASS"),
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "recordings",
	}

	var err error
	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	return db
}

func GetAlbums(c echo.Context) error {
	var albums []db.Album
	d := dbConn()
	rows, err := d.Query("SELECT * FROM album ORDER BY id ASC")
	if err != nil {
		return err
	}
	for rows.Next() {
		var alb db.Album
		if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
			return err
		}
		albums = append(albums, alb)
	}

	return c.JSON(http.StatusOK, albums)
}

func AddAlbum(c echo.Context) error {
	albs := []db.Album{
		{
			Title:  "Blue Train",
			Artist: "John Coltrane",
			Price:  56.99,
		},
		{
			Title:  "Giant Steps",
			Artist: "John Coltrane",
			Price:  63.99,
		},
		{
			Title:  "Jeru",
			Artist: "Gerry Mulligan",
			Price:  17.99,
		},
		{
			Title:  "Sarah Vaughan",
			Artist: "Sarah Vaughan",
			Price:  34.98,
		},
	}
	i := rand.Intn(4)
	d := dbConn()
	_, err := d.Exec("INSERT INTO album (title, artist, price) VALUES (?, ?, ?)", albs[i].Title, albs[i].Artist, albs[i].Price)
	if err != nil {
		return err
	}

	return nil
}

func Delete(c echo.Context) error {
	d := dbConn()
	id := c.Param("id")
	delForm, err := d.Prepare("DELETE FROM album WHERE id = ?")
	if err != nil {
		return err
	}
	delForm.Exec(id)
	defer d.Close()

	return nil
}
