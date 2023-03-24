package main

import (
	"bufio"
	"database/sql"
	"encoding/base64"
	"fmt"
	"github.com/labstack/echo/v4"
	_ "github.com/mattn/go-sqlite3"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

func UploadHander(c echo.Context) error {
	email := c.FormValue("estiemail")
	file, err := c.FormFile("estiphoto")
	if err != nil {
		log.Println(err)
	}
	fmt.Println(email)
	fmt.Println(file.Filename)

	ext := filepath.Ext(file.Filename)
	fmt.Println(ext)
	extlist := []string{".jpeg", "jpg", "png", "webp", "avif"}
	if contains(extlist, ext) {

		src, err := file.Open()
		if err != nil {
			log.Println(err)
		}
		defer src.Close()

		// dst, err := os.Create(file.Filename) // may need temp folder
		// if err != nil {
		// 	log.Println(err)
		// }
		// defer dst.Close()

		// // Copy
		// if _, err = io.Copy(dst, src); err != nil {
		// 	log.Println(err)
		// }
		// f, _ := os.Open(dst)

		reader := bufio.NewReader(src)
		content, _ := io.ReadAll(reader)
		encoded := base64.StdEncoding.EncodeToString(content)
		newimagestring := "data:image/jpeg;base64," + encoded

		log.Println("Starting InsertReview")
		var db_file string
		_, boo := os.LookupEnv("ATS_DOCKER_VAR")
		if boo {
			db_file = os.Getenv("ATS_PATH") + "/atsinfo.db"
		} else {
			db_file = "atsinfo.db"
		}

		rawdate := time.Now()
		ndate := rawdate.Format("01-13-2022")

		db, err := sql.Open("sqlite3", db_file) //production

		if err != nil {
			log.Fatal(err)
		}

		defer db.Close()
		nid := UUID()

		res, err := db.Exec("INSERT INTO photos VALUES(?,?,?,?)", &nid, &email, &ndate, &newimagestring)
		if err != nil {
			log.Println(err)
			log.Println("photo insert has failed")
		}
		var ret_val int
		_, err = res.LastInsertId()
		if err != nil {
			log.Println(err)
			ret_val = 1
		} else {
			ret_val = 0
		}

		//tranform to base64 and put in db in new table
		return c.JSON(http.StatusOK, ret_val)

	} else {
		//this is probably a video file
		src, err := file.Open()
		if err != nil {
			log.Println(err)
		}
		defer src.Close()

		dst, err := os.Create(file.Filename) // may need temp folder
		if err != nil {
			log.Println(err)
		}
		defer dst.Close()

		// Copy
		if _, err = io.Copy(dst, src); err != nil {
			log.Println(err)
		}
		rawdate := time.Now()
		ndate := rawdate.Format("01-13-2022")

		nid := UUID()

		var db_file string
		_, boo := os.LookupEnv("ATS_DOCKER_VAR")
		if boo {
			db_file = os.Getenv("ATS_PATH") + "/atsinfo.db"
		} else {
			db_file = "atsinfo.db"
		}

		db, err := sql.Open("sqlite3", db_file)
		if err != nil {
			log.Println(err)
		}
		fmt.Println(dst)
		log.Println(dst)

		res, err := db.Exec("INSERT INTO photos VALUES(?,?,?,?)", &nid, &email, &ndate, &dst)
		if err != nil {
			log.Println(err)
			log.Println("photo insert has failed")
		}
		var ret_val int
		_, err = res.LastInsertId()
		if err != nil {
			log.Println(err)
			ret_val = 1
		} else {
			ret_val = 0
		}

		return c.JSON(http.StatusOK, ret_val)
	}
}

func GetPhotoByEmailHandler(c echo.Context) error {

	rawstr := c.QueryString()
	parts := strings.Split(rawstr, "=")
	email := parts[1]

	log.Println("starting GetAllReviewsHandler")
	var db_file string
	_, boo := os.LookupEnv("ATS_DOCKER_VAR")
	if boo {
		db_file = os.Getenv("ATS_PATH") + "/atsinfo.db"
	} else {
		db_file = "atsinfo.db"
	}

	db, err := sql.Open("sqlite3", db_file) //production

	if err != nil {
		log.Fatal((err))
	}

	defer db.Close()

	rows, err := db.Query("SELECT * FROM photos WHERE email=?", email)
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	var photomap []map[string]string

	for rows.Next() {
		photoinfo := map[string]string{}
		var id string
		var email string
		var photo string

		err = rows.Scan(&id, &email, &photo)
		if err != nil {
			log.Println(err)
		}

		photoinfo["id"] = id
		photoinfo["email"] = email
		photoinfo["photo"] = photo

		photomap = append(photomap, photoinfo)
	}

	return c.JSON(http.StatusOK, photomap)
}
