package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
)

func Create_Admin_Tables() {
	bpath := os.Getenv("ATS_PATH") + "/atsinfo.db"
	db, err := sql.Open("sqlite3", bpath) // production

	if err != nil {
		db, err2 := sql.Open("sqlite3", "atsinfo.db") //testing
		if err2 != nil {
			log.Fatal(err2)
		}
		defer db.Close()

		sts := `
DROP TABLE IF EXISTS admin;
CREATE TABLE admin(id INTEGER PRIMARY KEY, name TEXT, email TEXT, date TEXT, time TEXT, token TEXT, pword TEXT);
`
		_, err = db.Exec(sts)

		if err != nil {
			log.Fatal(err)
		}

		log.Println("table admin created")

	}

	defer db.Close()

	sts := `
DROP TABLE IF EXISTS admin;
CREATE TABLE admin(id INTEGER PRIMARY KEY, name TEXT, email TEXT, date TEXT, time TEXT, token TEXT, pword TEXT);
`
	_, err = db.Exec(sts)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("table admin created")
}

func Create_Estimate_Tables() {
	bpath := os.Getenv("ATS_PATH") + "/atsinfo.db"
	db, err := sql.Open("sqlite3", bpath) // production

	if err != nil {
		db, err2 := sql.Open("sqlite3", "atsinfo.db") //testing
		if err2 != nil {
			log.Fatal(err2)
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
	bpath := os.Getenv("ATS_PATH") + "/atsinfo.db"
	db, err := sql.Open("sqlite3", bpath) // production

	if err != nil {
		db, err2 := sql.Open("sqlite3", "atsinfo.db") //testing
		if err2 != nil {
			log.Fatal(err2)
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
