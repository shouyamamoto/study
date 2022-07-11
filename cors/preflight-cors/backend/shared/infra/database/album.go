package database

import (
	"database/sql"
	"github/shouyamamoto/study/db"
	"log"
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
