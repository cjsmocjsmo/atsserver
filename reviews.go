package main

import (
	"database/sql"
	"github.com/labstack/echo/v4"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
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

func Insert_All_Initial_Comments() {
	var db_file string
	_, boo := os.LookupEnv("ATS_DOCKER_VAR")
	if boo {
		db_file = os.Getenv("ATS_PATH") + "/atsinfo.db"
	} else {
		db_file = "atsinfo.db"
	}

	db, err := sql.Open("sqlite3", db_file) // production

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	rev := `
INSERT INTO reviews (id, name, email, date, time, review, rating) VALUES('1', 'Scott Mason', 'smason@gmail.com', '2023-03-15', '04PM',
'Very responsive and easy to communicate with. Curtis and crew showed up when scheduled.  Very knowledgeable and professional. Mike did a great job in the tree and zip lined the branches perfectly with Curtis directing. Although he did get bit by the large thorns in the Locust tree. Would definitely recommend them to anyone looking to get problem trees down safely.', '5');
INSERT INTO revs_accepted(id, revid) VALUES('1', '1');
INSERT INTO reviews (id, name, email, date, time, review, rating) VALUES('2', 'Dan do1058', 'Dando1058@gmail.com', '2023-03-15', '11am',
'I contacted Curtis about removing several, dangerous trees on my property.  He showed up on time and ready to work. He did exactly what I expected him to do. He does exceptional work. I will continue to call Curtis when I need a tree removed. I would highly recommend AlphaTree.', '5');
INSERT INTO revs_accepted(id, revid) VALUES('2', '2');
INSERT INTO reviews (id, name, email, date, time, review, rating) VALUES('3', 'Kurt R', 'KurtR@gmail.com', '2023-03-15', '01PM',
'Curtis and crew took down an 80 foot fir near a fence and house. NO DAMAGE!!! Cleanup was thorough and they cut the rounds into 14 inch rounds for later splitting. Crew had a great attitude. Will use them again.', '5');
INSERT INTO revs_accepted(id, revid) VALUES('3', '3');
`
	_, err = db.Exec(rev)

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
		db_file = "atsinfo.db"
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
	ndate := rawdate.Format("2022-01-13")
	ntime := rawdate.Format("15:15:05")

	nreview := strings.ReplaceAll(parts[2], "SPACE", " ") //replace SPACE
	nrating := parts[3]

	res, err := db.Exec("INSERT INTO reviews VALUES(?,?,?,?,?,?,?)", &nid, &nname, &nemail, &ndate, &ntime, &nreview, &nrating)
	if err != nil {
		log.Println(err)
		log.Println("review insert has failed")
	}

	ret_val := 3
	_, err = res.LastInsertId()
	if err != nil {
		log.Println(err)
		ret_val = 2
	} else {
		ret_val = 0
	}

	res2, err2 := db.Exec("INSERT INTO revs_jailed VALUES(?,?)", nid, nid)
	if err2 != nil {
		log.Println(err2)
		log.Println("revs_jailed insert has failed")
	}

	_, err = res2.LastInsertId()
	if err != nil {
		log.Println(err)
		ret_val = 1
	} else {
		ret_val = 0
	}

	log.Printf("this is insert review exit status:\n\t %v", ret_val)

	return c.JSON(http.StatusOK, ret_val)
}

func get_accepted_reviews() []map[string]string {
	log.Println("starting GetAllReviewsHandler")
	var db_file string
	_, boo := os.LookupEnv("ATS_DOCKER_VAR")
	if boo {
		db_file = os.Getenv("ATS_PATH") + "/atsinfo.db"
	} else {
		db_file = "atsinfo.db"
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
			db_file = "atsinfo.db"
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

func get_jailed_reviews() []map[string]string {
	log.Println("starting GetAllReviewsHandler")
	var db_file string
	_, boo := os.LookupEnv("ATS_DOCKER_VAR")
	if boo {
		db_file = os.Getenv("ATS_PATH") + "/atsinfo.db"
	} else {
		db_file = "atsinfo.db"
	}

	db, err := sql.Open("sqlite3", db_file) //production

	if err != nil {
		log.Fatal((err))
	}

	defer db.Close()

	rows, err := db.Query("SELECT * FROM revs_jailed")
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	var jailed []map[string]string

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
		jailed = append(jailed, revv)
	}

	return jailed
}

func GetJailedReviewsHandler(c echo.Context) error {
	log.Println("GetJailedReviews has started")
	reviewz := []map[string]string{}
	jailedrevs := get_jailed_reviews()
	if len(jailedrevs) == 0 {
		return c.JSON(http.StatusOK, "0")
	}

	for _, arev := range jailedrevs {
		log.Println("starting GetJailedReviewsHandler")
		var db_file string
		_, boo := os.LookupEnv("ATS_DOCKER_VAR")
		if boo {
			db_file = os.Getenv("ATS_PATH") + "/atsinfo.db"
		} else {
			db_file = "atsinfo.db"
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

		log.Println("GetJailedReviews is complete")

	}
	return c.JSON(http.StatusOK, reviewz)
}

func AcceptReviewHandler(c echo.Context) error {
	log.Println("Starting AcceptReviewHandler")
	var db_file string
	_, boo := os.LookupEnv("ATS_DOCKER_VAR")
	if boo {
		db_file = os.Getenv("ATS_PATH") + "/atsinfo.db"
	} else {
		db_file = "atsinfo.db"
	}

	db, err := sql.Open("sqlite3", db_file) //production

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
	nid := UUID()

	rawstr := c.QueryString()
	parts := strings.Split(rawstr, "=")
	to_add_id := parts[1]

	log.Printf("this is query param accpeted %v", to_add_id)

	ret_val := 3
	_, err2 := db.Exec("DELETE FROM revs_jailed WHERE id=?", &to_add_id)
	if err2 != nil {
		log.Printf("this is revs acc err2 %v", err2)
		log.Println("revs_jailed deletion has failed")
		ret_val = 2
	}

	Ins1, err3 := db.Exec("INSERT INTO revs_accepted VALUES(?,?)", &nid, &to_add_id)
	if err3 != nil {
		log.Printf("this is revs acc err3 %v", err3)
		log.Println("rev_accepted insert has failed")
	}

	_, err = Ins1.LastInsertId()
	if err != nil {
		log.Printf("this is revs acc err %v", err)
		ret_val = 1
	} else {
		ret_val = 0
	}

	log.Printf("This is accept review return status:\n\t %v", ret_val)

	return c.JSON(http.StatusOK, ret_val)
}

func RejectReviewHandler(c echo.Context) error {
	log.Println("Starting RejectReviewHandler")
	var db_file string
	_, boo := os.LookupEnv("ATS_DOCKER_VAR")
	if boo {
		db_file = os.Getenv("ATS_PATH") + "/atsinfo.db"
	} else {
		db_file = "atsinfo.db"
	}

	db, err := sql.Open("sqlite3", db_file) //production

	if err != nil {
		log.Fatal((err))
	}

	defer db.Close()
	nid := UUID()

	rawstr := c.QueryString()
	parts := strings.Split(rawstr, "=")
	to_add_id := parts[1]

	log.Printf("tis is queryparams %v", to_add_id)

	ret_val := 3
	_, err2 := db.Exec("DELETE FROM revs_jailed WHERE id=?", &to_add_id)
	if err2 != nil {
		log.Printf("this is revs rej err2 %v", err2)
		log.Println("revs_jailed deletion has failed")
		ret_val = 2
	}

	Ins1, err3 := db.Exec("INSERT INTO revs_rejected VALUES(?,?)", &nid, &to_add_id)
	if err3 != nil {
		log.Printf("this is revs rej err3 %v", err3)
		log.Println("revs_rejected insert has failed")
	}

	_, err4 := Ins1.LastInsertId()
	if err4 != nil {
		log.Printf("this is revs rej err4 %v", err4)
		ret_val = 1
	} else {
		ret_val = 0
	}

	log.Printf("This is reject review return status:\n\t %v", ret_val)

	return c.JSON(http.StatusOK, ret_val)
}
