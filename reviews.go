package main

import (
	"compress/gzip"
	"database/sql"
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	_ "github.com/mattn/go-sqlite3"
)

func ATS_Logging() {
	var logfile string
	_, boo := os.LookupEnv("ATS_DOCKER_VAR")
	if boo {
		logfile = os.Getenv("ATS_LOG_PATH")
	} else {
		logfile = "/media/charliepi/HD/ats/atsserver/ATS.log"
	}

	file, err := os.OpenFile(logfile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(file)
	log.Println("ATS logging started")
}

func Insert_Comment_One() {
	var db_file string
	_, boo := os.LookupEnv("ATS_DOCKER_VAR")
	if boo {
		db_file = os.Getenv("ATS_PATH") + "/atsinfo.db"
	} else {
		db_file = "/media/charliepi/HD/ats/atsserver/atsinfo.db" //testing
	}

	db, err := sql.Open("sqlite3", db_file) // production

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	rev1 := `
INSERT INTO reviews (id, name, email, date, time, review, rating) VALUES('1', 'Scott Mason', 'smason@gmail.com', '2023-03-15', '04PM',
'Very responsive and easy to communicate with. Curtis and crew showed up when scheduled.  Very knowledgeable and professional. Mike did a great job in the tree and zip lined the branches perfectly with Curtis directing. Although he did get bit by the large thorns in the Locust tree. Would definitely recommend them to anyone looking to get problem trees down safely.', '5');
INSERT INTO revs_accepted(id, revid) VALUES('1', '1');
`
	_, err = db.Exec(rev1)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("insert review 001 complete")
}

func Insert_comment_two() {
	var db_file string
	_, boo := os.LookupEnv("ATS_DOCKER_VAR")
	if boo {
		db_file = os.Getenv("ATS_PATH") + "/atsinfo.db"
	} else {
		db_file = "/media/charliepi/HD/ats/atsserver/atsinfo.db" //testing
	}

	db, err := sql.Open("sqlite3", db_file) //testing
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
	rev2 := `
INSERT INTO reviews (id, name, email, date, time, review, rating) VALUES('2', 'Dan do1058', 'Dando1058@gmail.com', '2023-03-15', '11am',
'I contacted Curtis about removing several, dangerous trees on my property.  He showed up on time and ready to work. He did exactly what I expected him to do. He does exceptional work. I will continue to call Curtis when I need a tree removed. I would highly recommend AlphaTree.', '5');
INSERT INTO revs_accepted(id, revid) VALUES('2', '2');
`
	_, err = db.Exec(rev2)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("insert review 002 complete")
}

func Insert_comment_three() {
	var db_file string
	_, boo := os.LookupEnv("ATS_DOCKER_VAR")
	if boo {
		db_file = os.Getenv("ATS_PATH") + "/atsinfo.db"
	} else {
		db_file = "/media/charliepi/HD/ats/atsserver/atsinfo.db"
	}

	db, err := sql.Open("sqlite3", db_file) // production

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	rev3 := `
INSERT INTO reviews (id, name, email, date, time, review, rating) VALUES('3', 'Kurt R', 'KurtR@gmail.com', '2023-03-15', '01PM',
'Curtis and crew took down an 80 foot fir near a fence and house. NO DAMAGE!!! Cleanup was thorough and they cut the rounds into 14 inch rounds for later splitting. Crew had a great attitude. Will use them again.', '5');
INSERT INTO revs_accepted(id, revid) VALUES('3', '3');
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
	log.Println("Starting InsertReview")
	var db_file string
	_, boo := os.LookupEnv("ATS_DOCKER_VAR")
	if boo {
		db_file = os.Getenv("ATS_PATH") + "/atsinfo.db"
	} else {
		db_file = "/atsinfo.db"
	}

	db, err := sql.Open("sqlite3", db_file) //production

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
	nid := UUID()

	// MUST BE IN THE FORMAT
	// entry=JoSPACEBlowSPLITcjsmoATgmailDOTcomSPLITjobSPACEwellSPACEdoneSPLIT5
	rawstr := c.QueryString()
	parts := strings.Split(rawstr, "SPLIT")

	rawname := strings.Split(parts[0], "=")
	nname := strings.ReplaceAll(rawname[1], "SPACE", " ") // replace SPACE

	rawemail := strings.Replace(parts[1], "AT", "@", 1)
	nemail := strings.Replace(rawemail, "DOT", ".", 1)

	rawdate := time.Now()
	ndate := rawdate.Format("13-01-2022")
	ntime := rawdate.Format("15:15:05")

	// ntime := parts[3]
	nreview := strings.ReplaceAll(parts[3], "SPACE", " ") //replace SPACE
	nrating := parts[2]

	log.Println(nname)
	log.Println(nemail)
	log.Println(ndate)
	log.Println(ntime)
	log.Println(nreview)
	log.Println(nrating)

	res, err := db.Exec("INSERT INTO reviews VALUES(?,?,?,?,?,?,?)", &nid, &nname, &nemail, &ndate, &ntime, &nreview, &nrating)
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

	res2, err2 := db.Exec("INSERT INTO revs_jailed VALUES(?,?)", nid, nid)
	if err2 != nil {
		log.Println(err2)
		log.Println("revs_jailed insert has failed")
	}
	var ret_val2 int
	_, err = res2.LastInsertId()
	if err != nil {
		log.Println(err)
		ret_val2 = 1
	} else {
		ret_val2 = 0
	}

	result := []string{strconv.Itoa(ret_val), strconv.Itoa(ret_val2)}
	log.Printf("this is insert review exit status:\n\t %v", result)

	return c.JSON(http.StatusOK, result)
}

func get_accepted_reviews() []map[string]string {
	log.Println("starting GetAllReviewsHandler")
	var db_file string
	_, boo := os.LookupEnv("ATS_DOCKER_VAR")
	if boo {
		db_file = os.Getenv("ATS_PATH") + "/atsinfo.db"
	} else {
		db_file = "/atsinfo.db"
	}

	db, err := sql.Open("sqlite3", db_file) //production

	if err != nil {
		log.Fatal((err))
	}

	defer db.Close()

	rows, err := db.Query("SELECT * FROM revs_accepted")
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	var accepted []map[string]string

	for rows.Next() {
		revv := map[string]string{}
		var id string
		var revid string

		err = rows.Scan(&id, &revid)
		if err != nil {
			log.Println(err)
		}

		revv["id"] = id
		revv["revid"] = revid
		accepted = append(accepted, revv)
	}

	return accepted
}

func GetAllReviewsHandler(c echo.Context) error {
	log.Println("GetAllReviews has started")
	reviewz := []map[string]string{}
	allrevs := get_accepted_reviews()
	if len(allrevs) == 0 {
		return c.JSON(http.StatusOK, "0")
	}

	for _, arev := range allrevs {
		log.Println("starting GetAllReviewsHandler")
		var db_file string
		_, boo := os.LookupEnv("ATS_DOCKER_VAR")
		if boo {
			db_file = os.Getenv("ATS_PATH") + "/atsinfo.db"
		} else {
			db_file = "/atsinfo.db"
		}

		db, err := sql.Open("sqlite3", db_file) //production

		if err != nil {
			log.Fatal((err))
		}

		defer db.Close()

		rows, err := db.Query("SELECT * FROM reviews WHERE id=?", arev["revid"])
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		for rows.Next() {
			frev := map[string]string{}
			var id string
			var name string
			var email string
			var date string
			var time string
			var review string
			var rating string

			err = rows.Scan(&id, &name, &email, &date, &time, &review, &rating)
			if err != nil {
				log.Println(err)
			}

			frev["id"] = id
			frev["name"] = name
			frev["email"] = email
			frev["date"] = date
			frev["time"] = time
			frev["review"] = review
			frev["rating"] = rating
			reviewz = append(reviewz, frev)

		}

		log.Println("GetAllReviews is complete")
	}
	return c.JSON(http.StatusOK, reviewz)
}

func ReviewsGzipHandler(c echo.Context) error {
	log.Println("starting GetAllReviewsHandler")
	var db_file string
	_, boo := os.LookupEnv("ATS_DOCKER_VAR")
	if boo {
		db_file = os.Getenv("ATS_PATH") + "/atsinfo.db"
	} else {
		db_file = "/atsinfo.db"
	}

	db, err := sql.Open("sqlite3", db_file) //production

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	rows, err := db.Query("SELECT * FROM reviews")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	reviews := []map[string]string{}

	for rows.Next() {
		rev := map[string]string{}
		var id string

		var name string
		var email string
		var date string
		var time string
		var review string
		var rating string

		err = rows.Scan(&id, &name, &email, &date, &time, &review, &rating)
		if err != nil {
			log.Println(err)
		}

		rev["id"] = id
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
	buppath := os.Getenv("ATS_PATH") + "/static/dbbackup.tag.gz"
	f, _ := os.Create(buppath) //production
	// f, _ := os.Create("static/rev_db.tag.gz") //test
	w, _ := gzip.NewWriterLevel(f, gzip.BestCompression)
	w.Write([]byte(jsonstr))
	w.Close()
	log.Println("rev_db.tag.gz has been created")

	return c.JSON(http.StatusOK, "Backup Created")

}

func AcceptReviewHandler(c echo.Context) error {
	log.Println("Starting AcceptReviewHandler")
	var db_file string
	_, boo := os.LookupEnv("ATS_DOCKER_VAR")
	if boo {
		db_file = os.Getenv("ATS_PATH") + "/atsinfo.db"
	} else {
		db_file = "/atsinfo.db"
	}

	db, err := sql.Open("sqlite3", db_file) //production

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
	nid := UUID()

	to_add_id := c.QueryParam("revid")

	del1, err2 := db.Exec("DELETE FROM revs_jailed VALUES(?)", &to_add_id)
	if err2 != nil {
		log.Println(err2)
		log.Println("revs_jailed deletion has failed")
	}
	var ret_val int
	_, err = del1.LastInsertId()
	if err != nil {
		log.Println(err)
		ret_val = 1
	} else {
		ret_val = 0
	}

	Ins1, err2 := db.Exec("INSERT INTO revs_accepted VALUES(?,?)", &nid, &to_add_id)
	if err2 != nil {
		log.Println(err2)
		log.Println("rev_accepted insert has failed")
	}
	var ret_val2 int
	_, err = Ins1.LastInsertId()
	if err != nil {
		log.Println(err)
		ret_val2 = 1
	} else {
		ret_val2 = 0
	}

	result := []int{ret_val, ret_val2}

	log.Printf("This is accept review return status:\n\t %v", result)

	return c.JSON(http.StatusOK, result)
}

func RejectReviewHandler(c echo.Context) error {
	log.Println("Starting RejectReviewHandler")
	var db_file string
	_, boo := os.LookupEnv("ATS_DOCKER_VAR")
	if boo {
		db_file = os.Getenv("ATS_PATH") + "/atsinfo.db"
	} else {
		db_file = "/atsinfo.db"
	}

	db, err := sql.Open("sqlite3", db_file) //production

	if err != nil {
		log.Fatal((err))
	}

	defer db.Close()
	nid := UUID()

	to_add_id := c.QueryParam("revid")

	del1, err2 := db.Exec("DELETE FROM revs_jailed WHERE id=?)", &to_add_id)
	if err2 != nil {
		log.Println(err2)
		log.Println("revs_jailed deletion has failed")
	}
	var ret_val int
	_, err = del1.LastInsertId()
	if err != nil {
		log.Println(err)
		ret_val = 1
	} else {
		ret_val = 0
	}

	Ins1, err2 := db.Exec("INSERT INTO revs_rejected VALUES(?,?)", &nid, &to_add_id)
	if err2 != nil {
		log.Println(err2)
		log.Println("revs_rejected insert has failed")
	}
	var ret_val2 int
	_, err = Ins1.LastInsertId()
	if err != nil {
		log.Println(err)
		ret_val2 = 1
	} else {
		ret_val2 = 0
	}

	result := []int{ret_val, ret_val2}

	log.Printf("This is reject review return status:\n\t %v", result)

	return c.JSON(http.StatusOK, result)
}
