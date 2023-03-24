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

type CountsS struct {
	vids          string
	photos        string
	admin         string
	estimates     string
	esticompleted string
	estiworking   string
	revs          string
	revsaccepted  string
	revsjailed    string
	revsrejected  string
}

func CountzHandler(c echo.Context) error {
	// vc := video_count()
	// pc := photos_count()
	// ac := admin_count()
	// ec := estimates_count()
	// ecompc := est_completed_count()
	// eworkc := est_working_count()
	// rc := reviews_count()
	// racceptc := revs_accepted_count()
	// rjailedc := revs_jailed_count()
	// rrejectc := revs_rejected_count()

	result := []CountsS{}

	r := CountsS{}
	r.vids = video_count()
	r.photos = photos_count()
	r.admin = admin_count()
	r.estimates = estimates_count()
	r.esticompleted = est_completed_count()
	r.estiworking = est_working_count()
	r.revs = reviews_count()
	r.revsaccepted = revs_accepted_count()
	r.revsjailed = revs_jailed_count()
	r.revsrejected = revs_rejected_count()

	result = append(result, r)

	fmt.Println(r)

	return c.JSON(http.StatusOK, result)
}

func video_count() string {
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
		return video_count
	case err != nil:
		fmt.Printf("%s", err)
		return video_count
	default:
		fmt.Printf("Counted %s videos\n", video_count)
		return video_count
	}
}

func photos_count() string {
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
		return photos_count
	case err != nil:
		fmt.Printf("%s", err)
		return photos_count
	default:
		fmt.Printf("Counted %s photos\n", photos_count)
		return photos_count
	}
}

func admin_count() string {
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

	var admin_count string

	query, err := db.Prepare("SELECT count(*) FROM admin")
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
	defer query.Close()

	err = query.QueryRow().Scan(&admin_count)

	switch {
	case err == sql.ErrNoRows:
		fmt.Printf("No admin found.")
		return admin_count
	case err != nil:
		fmt.Printf("%s", err)
		return admin_count
	default:
		fmt.Printf("Counted %s admin\n", admin_count)
		return admin_count
	}
}

func estimates_count() string {
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

	var estimates_count string

	query, err := db.Prepare("SELECT count(*) FROM estimates")
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
	defer query.Close()

	err = query.QueryRow().Scan(&estimates_count)

	switch {
	case err == sql.ErrNoRows:
		fmt.Printf("No estimates found.")
		return estimates_count
	case err != nil:
		fmt.Printf("%s", err)
		return estimates_count
	default:
		fmt.Printf("Counted %s estimates\n", estimates_count)
		return estimates_count
	}
}

func est_completed_count() string {
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

	var est_completed_count string

	query, err := db.Prepare("SELECT count(*) FROM est_completed")
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
	defer query.Close()

	err = query.QueryRow().Scan(&est_completed_count)

	switch {
	case err == sql.ErrNoRows:
		fmt.Printf("No est_completed found.")
		return est_completed_count
	case err != nil:
		fmt.Printf("%s", err)
		return est_completed_count
	default:
		fmt.Printf("Counted %s est_completed\n", est_completed_count)
		return est_completed_count
	}
}

func est_working_count() string {
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

	var est_working_count string

	query, err := db.Prepare("SELECT count(*) FROM est_working")
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
	defer query.Close()

	err = query.QueryRow().Scan(&est_working_count)

	switch {
	case err == sql.ErrNoRows:
		fmt.Printf("No est_working found.")
		return est_working_count
	case err != nil:
		fmt.Printf("%s", err)
		return est_working_count
	default:
		fmt.Printf("Counted %s est_working\n", est_working_count)
		return est_working_count
	}
}

func reviews_count() string {
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
		return reviews_count
	case err != nil:
		fmt.Printf("%s", err)
		return reviews_count
	default:
		fmt.Printf("Counted %s reviews\n", reviews_count)
		return reviews_count
	}
}

func revs_accepted_count() string {
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

	var revs_accepted_count string

	query, err := db.Prepare("SELECT count(*) FROM revs_accepted")
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
	defer query.Close()

	err = query.QueryRow().Scan(&revs_accepted_count)

	switch {
	case err == sql.ErrNoRows:
		fmt.Printf("No revs_accepted found.")
		return revs_accepted_count
	case err != nil:
		fmt.Printf("%s", err)
		return revs_accepted_count
	default:
		fmt.Printf("Counted %s revs_accepted\n", revs_accepted_count)
		return revs_accepted_count
	}
}

func revs_rejected_count() string {
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

	var revs_rejected_count string

	query, err := db.Prepare("SELECT count(*) FROM revs_rejected")
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
	defer query.Close()

	err = query.QueryRow().Scan(&revs_rejected_count)

	switch {
	case err == sql.ErrNoRows:
		fmt.Printf("No revs_rejected found.")
		return revs_rejected_count
	case err != nil:
		fmt.Printf("%s", err)
		return revs_rejected_count
	default:
		fmt.Printf("Counted %s revs_rejected\n", revs_rejected_count)
		return revs_rejected_count
	}
}

func revs_jailed_count() string {
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

	var revs_jailed_count string

	query, err := db.Prepare("SELECT count(*) FROM revs_jailed")
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
	defer query.Close()

	err = query.QueryRow().Scan(&revs_jailed_count)

	switch {
	case err == sql.ErrNoRows:
		fmt.Printf("No revs_jailed found.")
		return revs_jailed_count
	case err != nil:
		fmt.Printf("%s", err)
		return revs_jailed_count
	default:
		fmt.Printf("Counted %s revs_jailed\n", revs_jailed_count)
		return revs_jailed_count
	}
}
