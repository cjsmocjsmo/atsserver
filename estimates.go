package main

import (
	"compress/gzip"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	_ "github.com/mattn/go-sqlite3"
)

func Create_Estimate_Tables() {
	db, err := sql.Open("sqlite3", "/usr/share/ats_server/atsinfo.db") // production
	// db, err := sql.Open("sqlite3", "atsinfo.db") //testing

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

func InsertEstimateHandler(c echo.Context) error {
	log.Println("Starting InsertEstimate")
	db, err := sql.Open("sqlite3", "/usr/share/ats_server/atsinfo.db") //production
	// db, err := sql.Open("sqlite3", "atsinfo.db") //testing
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
	log.Printf("this is querystring:\n\t %v\n", rawstr)

	parts := strings.Split(rawstr, "SPLIT")
	log.Println(parts)

	rawname := strings.Split(parts[0], "=")
	nname := strings.ReplaceAll(rawname[1], "SPACE", " ")

	naddress := strings.ReplaceAll(parts[1], "SPACE", " ")

	ncity := strings.ReplaceAll(parts[2], "SPACE", " ")

	ntelephone := parts[3]

	rawemail := strings.Replace(parts[4], "AT", "@", 1)
	nemail := strings.ReplaceAll(rawemail, "DOT", ".")

	reqservdate := parts[5]

	rawdate := time.Now()
	ndate := rawdate.Format("13-01-2022")
	ntime := rawdate.Format("15:15:05")

	// ndate := parts[4]
	// ntime := parts[5]

	ncomment := strings.ReplaceAll(parts[6], "SPACE", " ")

	nid := UUID()

	res, err := db.Exec("INSERT INTO estimates VALUES(?,?,?,?,?,?,?,?,?,?)", nid, nname, naddress, ncity, ntelephone, nemail, reqservdate, ndate, ntime, ncomment)
	if err != nil {
		log.Println(err)
		log.Println("estimates insert has failed")
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
	log.Printf("this is insert estimates:\n\t %v", r1)

	return c.JSON(http.StatusOK, r1)
}

func GetAllEstimatesHandler(c echo.Context) error {
	db, err := sql.Open("sqlite3", "/usr/share/ats_server/atsinfo.db") //production
	// db, err := sql.Open("sqlite3", "atsinfo.db") //testing
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
			var address string
			var city string
			var telephone string
			var email string
			var reqservdate string
			var date string
			var time string
			var comment string

			err = rows.Scan(&id, &name, &address, &city, &telephone, &email, &reqservdate, &date, &time, &comment)
			if err != nil {
				log.Println(err)
			}

			est["id"] = id
			est["name"] = name
			est["address"] = address
			est["city"] = city
			est["telephone"] = telephone
			est["email"] = email
			est["reqservdate"] = reqservdate
			est["date"] = date
			est["time"] = time
			est["comment"] = comment
			est_map = append(est_map, est)
		}
	}

	log.Printf("this is est_map:\n\t %v", est_map)

	return c.JSON(http.StatusOK, est_map)
}

func CompletEstimateHandler(c echo.Context) error {
	log.Println("Complete estimate has started")
	to_be_del := c.QueryParam("estvid")

	db, err := sql.Open("sqlite3", "/usr/share/ats_server/atsinfo.db") //production
	// db, err := sql.Open("sqlite3", "atsinfo.db") //testing
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

	log.Println("Complete estimate is finished")

	return c.JSON(http.StatusOK, result)
}

func EstimatesGzipHandler(c echo.Context) error {
	log.Println("starting GetAllReviewsHandler")
	db, err := sql.Open("sqlite3", "/usr/share/ats_server/atsinfo.db") //production
	// db, err := sql.Open("sqlite3", "atsinfo.db") //testing
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
		var address string
		var city string
		var telephone string
		var email string
		var reqservdate string
		var date string
		var time string
		var comment string

		err = rows.Scan(&id, &name, &address, &city, &telephone, &email, &reqservdate, &date, &time, &comment)
		if err != nil {
			log.Println(err)
		}

		est["id"] = id
		est["name"] = name
		est["address"] = address
		est["city"] = city
		est["telephone"] = telephone
		est["email"] = email
		est["reqservdate"] = reqservdate
		est["date"] = date
		est["time"] = time
		est["comment"] = comment
		estimates = append(estimates, est)

	}

	//convert to json
	jsonstr, err := json.Marshal(estimates)
	if err != nil {
		log.Fatal(err)
	}

	//gzip file and move it to static http folder
	//////////////////////////////////////////////////////////////////////////////////////////////////////////

	f, _ := os.Create("/usr/share/ats_server/static/dbbackup.tag.gz") //production
	// f, _ := os.Create("static/est_db.tag.gz") //test

	//////////////////////////////////////////////////////////////////////////////////////////////////////////
	w, _ := gzip.NewWriterLevel(f, gzip.BestCompression)
	w.Write([]byte(jsonstr))
	w.Close()

	log.Println("Estimates gzip is complete")

	return c.JSON(http.StatusOK, "Backup Created")
}
