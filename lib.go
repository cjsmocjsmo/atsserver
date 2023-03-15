package main

import (
	"database/sql"
	"fmt"
	"github.com/labstack/echo/v4"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"
)

func ATS_Logging() {
	logfile := os.Getenv("ATS_LOG_PATH") + "/ATS.log"
	// If the file doesn't exist, create it or append to the file
	file, err := os.OpenFile(logfile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(file)
	log.Println("ATS logging started")
}

func Create_Tables() {
	db, err := sql.Open("sqlite3", "/usr/share/ats_server/atsinfo.db")

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	sts := `
DROP TABLE IF EXISTS reviews;
CREATE TABLE reviews(id INTEGER PRIMARY KEY, uuid TEXT, name TEXT, email TEXT, date TEXT, time TEXT, review TEXT);
DROP TABLE IF EXISTS estimates;
CREATE TABLE estimates(id INTEGER PRIMARY KEY, uuid TEXT, name TEXT, address TEXT, city TEXT, telephone TEXT, email TEXT, reqservdate TEXT, comment TEXT, photo TEXT);
`
	_, err = db.Exec(sts)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("table reviews created")
}

func UUID() string {
	aseed := time.Now()
	aSeed := aseed.UnixNano()
	rand.Seed(aSeed)
	u := rand.Int63n(aSeed)
	uuid := strconv.FormatInt(u, 10)
	return uuid
}

func TestHandler(c echo.Context) error {
	test := "Hello from ats_comments."
	return c.JSON(http.StatusOK, test)
}

func InsertReviewHandler(c echo.Context) error {
	// db, err := sql.Open("sqlite3", "/usr/share/ats_server/atsinfo.db") //production
	db, err := sql.Open("sqlite3", "atsinfo.db") //testing
	if err != nil {
		log.Fatal((err))
	}

	defer db.Close()

	nuuid := UUID()
	nname := c.QueryParam("name")
	nemail := c.QueryParam("email")
	ndate := c.QueryParam("date")
	ntime := c.QueryParam("time")
	nreview := c.QueryParam("review")

	res, err := db.Exec("INSERT INTO reviews VALUES(NULL,?,?,?,?,?,?)", nuuid, nname, nemail, ndate, ntime, nreview)
	if err != nil {
		log.Println(err)
		log.Println("review insert has failed")
	}

	var id int64
	var ret_val int64
	id, err = res.LastInsertId()
	if err != nil {
		log.Println(err)
		ret_val = 0
	} else {
		ret_val = id
	}

	result := strconv.Itoa(int(ret_val))

	return c.JSON(http.StatusOK, result)
}

type RevewInfo struct {
	uuid   string
	name   string
	email  string
	date   string
	time   string
	review string
}

func GetAllReviewsHandler(c echo.Context) error {
	// db, err := sql.Open("sqlite3", "/usr/share/ats_server/atsinfo.db") //production
	db, err := sql.Open("sqlite3", "atsinfo.db") //testing
	if err != nil {
		log.Fatal((err))
	}

	defer db.Close()

	rows, err := db.Query("SELECT * FROM reviews")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var reviews []RevewInfo

	for rows.Next() {
		var rev RevewInfo
		var uuid string
		var name string
		var email string
		var date string
		var time string
		var review string

		err = rows.Scan(&uuid, &name, &email, &date, &time, &review)
		if err != nil {
			log.Println(err)
		}

		rev.uuid = uuid
		rev.name = name
		rev.email = email
		rev.date = date
		rev.time = time
		rev.review = review
		reviews = append(reviews, rev)

	}

	return c.JSON(http.StatusOK, reviews)
}

// func AcceptReviewHandler(c echo.Context) error {

// 	return c.JSON(http.StatusOK, ActionMedia)
// }

// func RejectReviewHandler(c echo.Context) error {

// 	return c.JSON(http.StatusOK, ActionMedia)
// }
