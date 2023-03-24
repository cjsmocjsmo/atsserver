package main

import (
	"database/sql"
	// "encoding/json"
	// "log"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	_ "github.com/mattn/go-sqlite3"
)

// type CountsS struct {
// 	vidcount             int
// 	photoscount          int
// 	admincount           int
// 	estimatescount       int
// 	estiworkingcount     int
// 	esticompletedcount   int
// 	reviewscount         int
// 	reviewsjailedcount   int
// 	reviewsacceptedcount int
// 	reviewsrejectedcount int
// }

func CountzHandler(c echo.Context) error {
	video_count()
	photos_count()
	reviews_count()

	return c.JSON(http.StatusOK, "Count Complete.")
}
func video_count() {
	var db_file string
	_, boo := os.LookupEnv("ATS_DOCKER_VAR")
	if boo {
		db_file = os.Getenv("ATS_PATH") + "/atsinfo.db"
	} else {
		db_file = "./atsinfo.db" //testing
	}

	db, err := sql.Open("sqlite3", db_file) // production
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}

	defer db.Close()

	var video_count string

	query, err := db.Prepare("SELECT count(*) FROM videos")
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
	defer query.Close()

	err = query.QueryRow().Scan(&video_count)

	switch {
	case err == sql.ErrNoRows:
		fmt.Printf("No videos found.")
	case err != nil:
		fmt.Printf("%s", err)
	default:
		fmt.Printf("Counted %s videos\n", video_count)
	}
}

func photos_count() {
	var db_file string
	_, boo := os.LookupEnv("ATS_DOCKER_VAR")
	if boo {
		db_file = os.Getenv("ATS_PATH") + "/atsinfo.db"
	} else {
		db_file = "./atsinfo.db" //testing
	}

	db, err := sql.Open("sqlite3", db_file) // production
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}

	defer db.Close()

	var photos_count string

	query, err := db.Prepare("SELECT count(*) FROM photos")
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
	defer query.Close()

	err = query.QueryRow().Scan(&photos_count)

	switch {
	case err == sql.ErrNoRows:
		fmt.Printf("No photos found.")
	case err != nil:
		fmt.Printf("%s", err)
	default:
		fmt.Printf("Counted %s photos\n", photos_count)
	}
}

// admin_count, err := db.Query("SELECT count(*) FROM admin")
// if err != nil {
// 	log.Println("admin_count has failed")
// 	log.Fatal(err)
// }

// estimates_count, err := db.Query("SELECT count(*) FROM estimates")
// if err != nil {
// 	log.Println("estimates_count has failed")
// 	log.Fatal(err)
// }

// est_completed_count, err := db.Query("SELECT count(*) FROM est_completed")
// if err != nil {
// 	log.Println("est_completed_count has failed")
// 	log.Fatal(err)
// }

// est_working_count, err := db.Query("SELECT count(*) FROM est_working")
// if err != nil {
// 	log.Println("est_working_count has failed")
// 	log.Fatal(err)
// }

func reviews_count() {
	var db_file string
	_, boo := os.LookupEnv("ATS_DOCKER_VAR")
	if boo {
		db_file = os.Getenv("ATS_PATH") + "/atsinfo.db"
	} else {
		db_file = "./atsinfo.db" //testing
	}

	db, err := sql.Open("sqlite3", db_file) // production
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}

	defer db.Close()

	var reviews_count string

	query, err := db.Prepare("SELECT count(*) FROM reviews")
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
	defer query.Close()

	err = query.QueryRow().Scan(&reviews_count)

	switch {
	case err == sql.ErrNoRows:
		fmt.Printf("No reviews found.")
	case err != nil:
		fmt.Printf("%s", err)
	default:
		fmt.Printf("Counted %s reviews\n", reviews_count)
	}
}

// revs_accepted_count, err := db.Query("SELECT count(*) FROM revs_accepted")
// if err != nil {
// 	log.Println("revs_accepted_count has failed")
// 	log.Fatal(err)
// }

// revs_rejected_count, err := db.Query("SELECT count(*) FROM revs_rejected")
// if err != nil {
// 	log.Println("revs_rejected_count has failed")
// 	log.Fatal(err)
// }

// revs_jailed_count, err := db.Query("SELECT count(*) FROM revs_jailed")
// if err != nil {
// 	log.Println("revs_jailed_count has failed")
// 	log.Fatal(err)
// }

// SELECT COUNT(*) FROM t1;
// return c.JSON(http.StatusOK, "fuckit")
