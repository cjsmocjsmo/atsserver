package main

import (
	// "compress/gzip"
	"database/sql"
	// "encoding/json"
	"github.com/labstack/echo/v4"
	_ "github.com/mattn/go-sqlite3"
	"log"
	// "math/rand"
	"net/http"
	// "os"
	// "strconv"
	// "strings"
	// "time"
)

func Create_Admin_Tables() {
	// db, err := sql.Open("sqlite3", "/usr/share/ats_server/atsinfo.db") // production
	db, err := sql.Open("sqlite3", "atsinfo.db") //testing

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	sts := `
DROP TABLE IF EXISTS admin;
CREATE TABLE admin(id INTEGER PRIMARY KEY, name TEXT, email TEXT, date TEXT, time TEXT, review TEXT, rating TEXT);
`
	_, err = db.Exec(sts)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("table admin created")
}

func admin_sign_in(c echo.Context) error {

	return c.JSON(http.StatusOK, "fuck")
}

func admin_sign_out(c echo.Context) error {

	return c.JSON(http.StatusOK, "fuck")
}
