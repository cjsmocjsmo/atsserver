package main

import (
	// "bufio"
	"compress/gzip"
	"database/sql"
	"encoding/json"
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
	// logfile := os.Getenv("ATS_LOG_PATH") + "/ATS.log"
	logfile := "/media/charliepi/HD/ats/atsserver/ATS.log"
	// If the file doesn't exist, create it or append to the file
	file, err := os.OpenFile(logfile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(file)
	log.Println("ATS logging started")
}

func Create_Tables() {
	// db, err := sql.Open("sqlite3", "/usr/share/ats_server/atsinfo.db") // production
	db, err := sql.Open("sqlite3", "atsinfo.db") //testing

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	sts := `
DROP TABLE IF EXISTS reviews;
CREATE TABLE reviews(id INTEGER PRIMARY KEY, uuid TEXT, name TEXT, email TEXT, date TEXT, time TEXT, review TEXT, rating TEXT);
DROP TABLE IF EXISTS estimates;
CREATE TABLE estimates(id INTEGER PRIMARY KEY, uuid TEXT, name TEXT, address TEXT, city TEXT, telephone TEXT, email TEXT, reqservdate TEXT, comment TEXT, photo TEXT);
`
	_, err = db.Exec(sts)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("table reviews created")
}

func Insert_Comment_One() {
	// db, err := sql.Open("sqlite3", "/usr/share/ats_server/atsinfo.db") // production
	db, err := sql.Open("sqlite3", "atsinfo.db") //testing

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	rev1 := `
INSERT INTO reviews (id, uuid, name, email, date, time, review, rating) VALUES('1', '001', 'Scott Mason', 'smason@gmail.com', '2023-03-15', '04PM',
'Very responsive and easy to communicate with. Curtis and crew showed up when scheduled.  Very knowledgeable and professional. Mike did a great job in the tree and zip lined the branches perfectly with Curtis directing. Although he did get bit by the large thorns in the Locust tree. Would definitely recommend them to anyone looking to get problem trees down safely.', '5');
`
	_, err = db.Exec(rev1)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("insert review 001 complete")
}

func Insert_comment_two() {
	// db, err := sql.Open("sqlite3", "/usr/share/ats_server/atsinfo.db") // production
	db, err := sql.Open("sqlite3", "atsinfo.db") //testing

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
	rev2 := `
INSERT INTO reviews (id, uuid, name, email, date, time, review, rating) VALUES('2', '002', 'Dan do1058', 'Dando1058@gmail.com', '2023-03-15', '11am',
'I contacted Curtis about removing several, dangerous trees on my property.  He showed up on time and ready to work. He did exactly what I expected him to do. He does exceptional work. I will continue to call Curtis when I need a tree removed. I would highly recommend AlphaTree.', '5');
`
	_, err = db.Exec(rev2)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("insert review 002 complete")
}

func Insert_comment_three() {
	// db, err := sql.Open("sqlite3", "/usr/share/ats_server/atsinfo.db") // production
	db, err := sql.Open("sqlite3", "atsinfo.db") //testing

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	rev3 := `
INSERT INTO reviews (id, uuid, name, email, date, time, review, rating) VALUES('3', '003', 'Kurt R', 'KurtR@gmail.com', '2023-03-15', '01PM',
'Curtis and crew took down an 80 foot fir near a fence and house. NO DAMAGE!!! Cleanup was thorough and they cut the rounds into 14 inch rounds for later splitting. Crew had a great attitude. Will use them again.', '5');
`
	_, err = db.Exec(rev3)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("insert review 003 complete")
}

func UUID() string {
	aseed := time.Now()
	aSeed := aseed.UnixNano()
	rand.Seed(aSeed)
	u := rand.Int63n(aSeed)
	uuid := strconv.FormatInt(u, 10)
	return uuid
}

func Insert_Initial_Comments() {

	Insert_Comment_One()
	Insert_comment_two()
	Insert_comment_three()
}

func TestHandler(c echo.Context) error {
	test := "Hello from ats_comments."
	return c.JSON(http.StatusOK, test)
}

func InsertReviewHandler(c echo.Context) error {
	db, err := sql.Open("sqlite3", "/usr/share/ats_server/atsinfo.db") //production
	// db, err := sql.Open("sqlite3", "atsinfo.db") //testing
	if err != nil {
		log.Fatal((err))
	}

	defer db.Close()
	uid := UUID()
	nid := uid
	nuuid := uid
	nname := c.QueryParam("name")
	nemail := c.QueryParam("email")
	ndate := c.QueryParam("date")
	ntime := c.QueryParam("time")
	nreview := c.QueryParam("review")
	nrating := c.QueryParam("rating")

	res, err := db.Exec("INSERT INTO reviews VALUES(?,?,?,?,?,?,?,?)", nid, nuuid, nname, nemail, ndate, ntime, nreview, nrating)
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

func GetAllReviewsHandler(c echo.Context) error {
	log.Println("starting GetAllReviewsHandler")
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

	var reviews []map[string]string

	for rows.Next() {
		rev := map[string]string{}
		var id string
		var uuid string
		var name string
		var email string
		var date string
		var time string
		var review string
		var rating string

		err = rows.Scan(&id, &uuid, &name, &email, &date, &time, &review, &rating)
		if err != nil {
			log.Println(err)
		}

		rev["id"] = id
		rev["uuid"] = uuid
		rev["name"] = name
		rev["email"] = email
		rev["date"] = date
		rev["time"] = time
		rev["review"] = review
		rev["rating"] = rating
		reviews = append(reviews, rev)

	}
	return c.JSON(http.StatusOK, reviews)
}

func DumpGzipHandler(c echo.Context) error {
	log.Println("starting GetAllReviewsHandler")
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

	var reviews []map[string]string

	for rows.Next() {
		rev := map[string]string{}
		var id string
		var uuid string
		var name string
		var email string
		var date string
		var time string
		var review string
		var rating string

		err = rows.Scan(&id, &uuid, &name, &email, &date, &time, &review, &rating)
		if err != nil {
			log.Println(err)
		}

		rev["id"] = id
		rev["uuid"] = uuid
		rev["name"] = name
		rev["email"] = email
		rev["date"] = date
		rev["time"] = time
		rev["review"] = review
		rev["rating"] = rating
		reviews = append(reviews, rev)

	}

	//convert to json
	jsonstr, err := json.Marshal(reviews)
	if err != nil {
		log.Fatal(err)
	}

	//gzip file and move it to static http folder

	// f, _ := os.Create("/usr/share/ats_server/static/dbbackup.tag.gz") //production
	f, _ := os.Create("static/dbbackup.tag.gz") //test
	w, _ := gzip.NewWriterLevel(f, gzip.BestCompression)
	w.Write([]byte(jsonstr))
	w.Close()

	return c.JSON(http.StatusOK, "Backup Created")

}

// func AcceptReviewHandler(c echo.Context) error {

// 	return c.JSON(http.StatusOK, ActionMedia)
// }

// func RejectReviewHandler(c echo.Context) error {

// 	return c.JSON(http.StatusOK, ActionMedia)
// }
