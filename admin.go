package main

import (
	"crypto/sha256"
	"database/sql"
	"encoding/hex"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/labstack/echo/v4"
	_ "github.com/mattn/go-sqlite3"
	"gopkg.in/yaml.v3"
)

type UserS struct {
	Name     string `yaml:"name"`
	Email    string `yaml:"email"`
	Token    string `yaml:"token"`
	Password string `yaml:"password"`
	Date     string `yaml:"date"`
	Time     string `yaml:"time"`
}

func glob_user_dir() []string {
	var pattern string
	_, boo := os.LookupEnv("ATS_DOCKER_VAR")
	if boo {
		pattern = os.Getenv("ATS_PATH") + "/users/*.yaml"
	} else {
		pattern = "/media/charliepi/HD/ats/atsserver/*.yaml" //testing
	}
	matches, err := filepath.Glob(pattern)
	if err != nil {
		log.Println(err)
	}
	return matches
}

func parse_admin_list() []UserS {
	pfiles := glob_user_dir()
	user_list := []UserS{}
	for _, p := range pfiles {
		data, err := os.ReadFile(p)
		if err != nil {
			log.Println(err)
		}
		u := UserS{}
		err = yaml.Unmarshal(data, &u)
		if err != nil {
			log.Println(err)
		}
		user_list = append(user_list, u)
	}
	return user_list
}

func Insert_Admins(x UserS) int {
	id := UUID()
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

	newemail := strings.Replace(x.Email, "AT", "@", 1)
	nemail := strings.ReplaceAll(newemail, "DOT", ".")
	ndate := strings.ReplaceAll(x.Date, "_", "-")

	res, err := db.Exec("INSERT INTO admin VALUES(?,?,?,?,?,?,?)", id, x.Name, nemail, ndate, x.Time, x.Token, x.Password)
	if err != nil {
		log.Println("admin insert has failed")
	}
	var ret_val int
	_, err = res.LastInsertId()
	if err != nil {
		log.Printf("this is last insert id err %v", err)
		ret_val = 1
	} else {
		ret_val = 0
	}
	log.Printf("insert admin return val %v", ret_val)
	return ret_val
}

func insert_loggedin(email string) int {

	var db_file string
	_, boo := os.LookupEnv("ATS_DOCKER_VAR")
	if boo {
		db_file = os.Getenv("ATS_PATH") + "/atsinfo.db" // production

	} else {
		db_file = "atsinfo.db" //testing
	}

	db, err := sql.Open("sqlite3", db_file)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	id := UUID()
	// newemail := strings.Replace(x.Email, "AT", "@", 1)
	// nemail := strings.ReplaceAll(newemail, "DOT", ".")
	// ndate := strings.ReplaceAll(x.Date, "_", "-")

	res, err := db.Exec("INSERT INTO loggedin VALUES(?,?)", id, email)
	if err != nil {
		log.Println("admin insert has failed")
	}
	var ret_val int
	_, err = res.LastInsertId()
	if err != nil {
		log.Printf("this is last insert id err %v", err)
		ret_val = 1
	} else {
		ret_val = 0
	}
	log.Printf("insert admin return val %v", ret_val)
	return ret_val
}

func Create_Admin() {
	alist := parse_admin_list()
	for _, admin := range alist {
		Insert_Admins(admin)
	}
}

func parse_query_string(x string) (string, string, string) {
	parts := strings.Split(x, "_")
	rawname := strings.Split(parts[0], "=")
	return rawname[1], parts[1], parts[2]
}

func get_hash(x string) string {
	h := sha256.New()
	h.Write([]byte(x))
	ash := h.Sum(nil)
	hash := hex.EncodeToString(ash)
	return hash
}

func get_admin_by_email(x string) map[string]string {
	var db_file string
	_, boo := os.LookupEnv("ATS_DOCKER_VAR")
	if boo {
		db_file = os.Getenv("ATS_PATH") + "/atsinfo.db"
	} else {
		db_file = "atsinfo.db"
	}

	db, err := sql.Open("sqlite3", db_file)

	if err != nil {
		log.Fatal((err))
	}

	defer db.Close()

	rows, err := db.Query("SELECT * FROM admin WHERE email=?", x)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	aadmin := map[string]string{}
	for rows.Next() {

		var id string
		var name string
		var email string
		var date string
		var time string
		var token string
		var password string
		err = rows.Scan(&id, &name, &email, &date, &time, &token, &password)
		if err != nil {
			log.Println(err)
		}

		aadmin["id"] = id
		aadmin["name"] = name
		aadmin["email"] = email
		aadmin["date"] = date
		aadmin["time"] = time
		aadmin["token"] = token
		aadmin["password"] = password
	}
	return aadmin
}

func comp_str(x string, y string) bool {
	if string(x) != string(y) {
		return false
	} else {
		return true
	}
}

func LoginHandler(c echo.Context) error {
	rawstr := c.QueryString()
	log.Println(rawstr)
	// t := c.FormValue("token")
	// e := c.FormValue("email")
	// p := c.FormValue("pwd")
	// log.Println(t)
	// log.Println(e)
	// log.Println(p)

	t, e, p := parse_query_string(rawstr)
	thash := get_hash(t)
	ehash := get_hash(e)
	phash := get_hash(p)
	admin_info_db := get_admin_by_email(e)
	edb := get_hash(admin_info_db["email"])

	comp1 := comp_str(thash, admin_info_db["token"])
	comp2 := comp_str(ehash, edb)
	comp3 := comp_str(phash, admin_info_db["password"])

	isLoggedIn := "false"
	if comp1 && comp2 && comp3 {
		insert_loggedin(e)
		isLoggedIn = "true"
	}

	return c.JSON(http.StatusOK, isLoggedIn)
}

func LogoutHandler(c echo.Context) error {
	// rawstr := c.QueryString()
	// parts := strings.Split(rawstr, "=")
	// email := parts[1]

	log.Println("Starting LogoutHandler")
	var db_file string
	_, boo := os.LookupEnv("ATS_DOCKER_VAR")
	if boo {
		db_file = os.Getenv("ATS_PATH") + "/atsinfo.db" //production
	} else {
		db_file = "atsinfo.db"
	}

	db, err := sql.Open("sqlite3", db_file)

	if err != nil {
		log.Fatal((err))
	}

	defer db.Close()

	rm_id := c.QueryParam("email")
	ret_val := 0

	_, err2 := db.Exec("DELETE FROM loggedin WHERE email=?)", &rm_id)
	if err2 != nil {
		log.Println(err2)
		log.Println("revs_jailed deletion has failed")
		ret_val = 1
	}

	return c.JSON(http.StatusOK, ret_val)
}
