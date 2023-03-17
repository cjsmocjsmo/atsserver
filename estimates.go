package main

import (
	"compress/gzip"
	"database/sql"
	"encoding/json"
	"github.com/labstack/echo/v4"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func Create_Estimate_Tables() {
	// db, err := sql.Open("sqlite3", "/usr/share/ats_server/atsinfo.db") // production
	db, err := sql.Open("sqlite3", "atsinfo.db") //testing

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	sts := `
DROP TABLE IF EXISTS estimates;
CREATE TABLE estimates(id INTEGER PRIMARY KEY, name TEXT, address TEXT, city TEXT, telephone TEXT, email TEXT, reqservdate TEXT, comment TEXT);
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

func InsertEstimateHandler(c echo.Context) error {
	// db, err := sql.Open("sqlite3", "/usr/share/ats_server/atsinfo.db") //production
	db, err := sql.Open("sqlite3", "atsinfo.db") //testing
	if err != nil {
		log.Fatal((err))
	}

	defer db.Close()

	// MUST BE IN THE FORMAT
	// entry=BradSPACEPittSPLIT
	// 924SPACEHullSPACEAveSPLIT
	// PortSPACEOrchardSPLIT
	// 903-465-7811SPLITfooATgmailDOTcomSPLIT03-15-2023SPLIT0800SPLITASPACEtreeSPACEfellSPACEneedsSPACEcleanSPACEupSPLIT
	rawstr := c.QueryString()

	parts := strings.Split(rawstr, "SPLIT")
	rawname := strings.Split(parts[0], "=")
	nname := strings.ReplaceAll(rawname[1], "SPACE", " ")

	rawemail := strings.Replace(parts[3], "AT", "@", 1)
	nemail := strings.ReplaceAll(rawemail, "DOT", ".")

	ndate := parts[4]
	ntime := parts[5]

	naddress := strings.ReplaceAll(parts[1], "SPACE", " ")
	ncity := strings.ReplaceAll(parts[2], "SPACE", " ")

	ncomment := strings.ReplaceAll(parts[6], "SPACE", " ")

	nid := UUID()

	res, err := db.Exec("INSERT INTO reviews VALUES(?,?,?,?,?,?,?,?)", nid, nname, nemail, ndate, ntime, naddress, ncity, ncomment)
	if err != nil {
		log.Println(err)
		log.Println("review insert has failed")
	}

	var ret_val int
	_, err = res.LastInsertId()
	if err != nil {
		log.Println(err)
		ret_val = 1
	} else {
		ret_val = 0
	}

	res2, err2 := db.Exec("INSERT INTO est_working VALUES(?,?)", nid, nid)
	if err2 != nil {
		log.Println(err2)
		log.Println("review insert has failed")
	}

	var ret_val2 int
	_, err = res2.LastInsertId()
	if err != nil {
		log.Println(err)
		ret_val = 1
	} else {
		ret_val = 0
	}

	r1 := []string{strconv.Itoa(ret_val), strconv.Itoa(ret_val2)}

	return c.JSON(http.StatusOK, r1)
}

func GetAllEstimatesHandler(c echo.Context) error {
	// db, err := sql.Open("sqlite3", "/usr/share/ats_server/atsinfo.db") //production
	db, err := sql.Open("sqlite3", "atsinfo.db") //testing
	if err != nil {
		log.Fatal((err))
	}

	defer db.Close()

	rows, err := db.Query("SELECT * FROM est_working")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var estworking []string

	for rows.Next() {
		var id string
		var estid string

		err = rows.Scan(&id, &estid)
		if err != nil {
			log.Println(err)
		}
		estworking = append(estworking, estid)
	}

	est_map := []map[string]string{}

	for _, estid := range estworking {
		rows, err := db.Query("SELECT * FROM estimates WHERE id=?", estid)
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()
		for rows.Next() {
			est := map[string]string{}
			var id string

			var name string
			var email string
			var date string
			var time string
			var address string
			var city string
			var comment string

			err = rows.Scan(&id, &name, &email, &date, &time, &address, &city, &comment)
			if err != nil {
				log.Println(err)
			}

			est["id"] = id
			est["name"] = name
			est["email"] = email
			est["date"] = date
			est["time"] = time
			est["address"] = address
			est["city"] = city
			est["comment"] = comment
			est_map = append(est_map, est)
		}
	}

	return c.JSON(http.StatusOK, est_map)
}

func CompletEstimateHandler(c echo.Context) error {
	to_be_del := c.QueryParam("estvid")

	// db, err := sql.Open("sqlite3", "/usr/share/ats_server/atsinfo.db") //production
	db, err := sql.Open("sqlite3", "atsinfo.db") //testing
	if err != nil {
		log.Fatal((err))
	}

	defer db.Close()
	del1, err2 := db.Exec("DELETE FROM est_working WHERE id=?)", &to_be_del)
	if err2 != nil {
		log.Println(err2)
		log.Println("est_working deletion has failed")
	}
	var ret_val int
	_, err = del1.LastInsertId()
	if err != nil {
		log.Println(err)
		ret_val = 1
	} else {
		ret_val = 0
	}
	//delete from working add to completed
	res2, err2 := db.Exec("INSERT INTO est_completed VALUES(?,?)", to_be_del, to_be_del)
	if err2 != nil {
		log.Println(err2)
		log.Println("est_completed insert has failed")
	}

	var ret_val2 int
	_, err = res2.LastInsertId()
	if err != nil {
		log.Println(err)
		ret_val = 1
	} else {
		ret_val = 0
	}

	result := []int{ret_val, ret_val2}

	return c.JSON(http.StatusOK, result)
}

func EstimatesGzipHandler(c echo.Context) error {
	log.Println("starting GetAllReviewsHandler")
	// db, err := sql.Open("sqlite3", "/usr/share/ats_server/atsinfo.db") //production
	db, err := sql.Open("sqlite3", "atsinfo.db") //testing
	if err != nil {
		log.Fatal((err))
	}

	defer db.Close()

	rows, err := db.Query("SELECT * FROM estimates")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	estimates := []map[string]string{}

	for rows.Next() {
		est := map[string]string{}
		var id string

		var name string
		var email string
		var date string
		var time string
		var address string
		var city string
		var comment string

		err = rows.Scan(&id, &name, &email, &date, &time, &address, &city, &comment)
		if err != nil {
			log.Println(err)
		}

		est["id"] = id
		est["name"] = name
		est["email"] = email
		est["date"] = date
		est["time"] = time
		est["address"] = address
		est["city"] = city
		est["comment"] = comment
		estimates = append(estimates, est)

	}

	//convert to json
	jsonstr, err := json.Marshal(estimates)
	if err != nil {
		log.Fatal(err)
	}

	//gzip file and move it to static http folder

	// f, _ := os.Create("/usr/share/ats_server/static/dbbackup.tag.gz") //production
	f, _ := os.Create("static/est_db.tag.gz") //test
	w, _ := gzip.NewWriterLevel(f, gzip.BestCompression)
	w.Write([]byte(jsonstr))
	w.Close()

	return c.JSON(http.StatusOK, "Backup Created")
}
