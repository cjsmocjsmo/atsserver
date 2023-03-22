package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func Create_Admin_Tables() {
	var db_file string
	_, boo := os.LookupEnv("ATS_DOCKER_VAR")
	if boo {
		db_file = os.Getenv("ATS_PATH") + "/atsinfo.db"
	} else {
		db_file = "./atsinfo.db" //testing
	}

	fmt.Println(db_file)

	db, err := sql.Open("sqlite3", db_file) // production

	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}

	defer db.Close()

	sts := `
DROP TABLE IF EXISTS admin;
CREATE TABLE admin(id INTEGER PRIMARY KEY, name TEXT, email TEXT, date TEXT, time TEXT, token TEXT, password TEXT);
`
	_, err = db.Exec(sts)

	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}

	log.Println("table admin created")
}

func Create_Estimate_Tables() {
	var db_file string
	_, boo := os.LookupEnv("ATS_DOCKER_VAR")
	if boo {
		db_file = os.Getenv("ATS_PATH") + "/atsinfo.db"
	} else {
		db_file = "atsinfo.db" //testing
	}

	db, err := sql.Open("sqlite3", db_file) // production

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	sts := `
DROP TABLE IF EXISTS estimates;
CREATE TABLE estimates(id INTEGER PRIMARY KEY, name TEXT, address TEXT, city TEXT, telephone TEXT, email TEXT, reqservdate TEXT, date TEXT, time TEXT, comment TEXT);
DROP TABLE IF EXISTS est_completed;
CREATE TABLE est_completed(id INTEGER PRIMARY KEY, estid TEXT);
DROP TABLE IF EXISTS est_working;
CREATE TABLE est_working(id INTEGER PRIMARY KEY, estid TEXT);
`
	_, err = db.Exec(sts)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("table reviews created")
}

func Create_Reviews_Tables() {
	var db_file string
	_, boo := os.LookupEnv("ATS_DOCKER_VAR")
	if boo {
		db_file = os.Getenv("ATS_PATH") + "/atsinfo.db"
	} else {
		db_file = "atsinfo.db" //testing
	}

	db, err := sql.Open("sqlite3", db_file) // production

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	sts := `
DROP TABLE IF EXISTS reviews;
CREATE TABLE reviews(id INTEGER PRIMARY KEY, name TEXT, email TEXT, date TEXT, time TEXT, review TEXT, rating TEXT);
DROP TABLE IF EXISTS revs_accepted;
CREATE TABLE revs_accepted(id INTEGER PRIMARY KEY, revid TEXT);
DROP TABLE IF EXISTS revs_rejected;
CREATE TABLE revs_rejected(id INTEGER PRIMARY KEY, revid TEXT);
DROP TABLE IF EXISTS revs_jailed;
CREATE TABLE revs_jailed(id INTEGER PRIMARY KEY, revid TEXT);
`
	_, err = db.Exec(sts)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("table reviews created")
}
