package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
)

func Create_ALL_Tables() {
	log.Println("Starting Create All Tables")
	var db_file string
	_, boo := os.LookupEnv("ATS_DOCKER_VAR")
	if boo {
		db_file = os.Getenv("ATS_PATH") + "/atsinfo.db"
	} else {
		db_file = "./atsinfo.db" //testing
	}

	db, err := sql.Open("sqlite3", db_file)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	sts := `
DROP TABLE IF EXISTS videos;
CREATE TABLE videos(id INTEGER PRIMARY KEY, email TEXT, date TEXT, photo TEXT);
DROP TABLE IF EXISTS photos;
CREATE TABLE photos(id INTEGER PRIMARY KEY, email TEXT, date TEXT, photo TEXT);
DROP TABLE IF EXISTS admin;
CREATE TABLE admin(id INTEGER PRIMARY KEY, name TEXT, email TEXT, date TEXT, time TEXT, token TEXT, password TEXT);

DROP TABLE IF EXISTS loggedin;
CREATE TABLE loggedin(id INTERGER PRIMARY KEY, email TEXT)

DROP TABLE IF EXISTS estimates;
CREATE TABLE estimates(id INTEGER PRIMARY KEY, name TEXT, address TEXT, city TEXT, telephone TEXT, email TEXT, reqservdate TEXT, date TEXT, time TEXT, comment TEXT);
DROP TABLE IF EXISTS est_completed;
CREATE TABLE est_completed(id INTEGER PRIMARY KEY, estid TEXT);
DROP TABLE IF EXISTS est_working;
CREATE TABLE est_working(id INTEGER PRIMARY KEY, estid TEXT);
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

	log.Println("Create All Tables complete")
}
