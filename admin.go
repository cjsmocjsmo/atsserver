package main

import (
	// "github.com/labstack/echo/v4"
	"database/sql"

	_ "github.com/mattn/go-sqlite3"

	// "encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
	// "net/http"
	// "strings"
)

type UserS struct {
	// ID       string `yaml:"id"`
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
		pattern = "/media/charliepi/HD/ats/atsserver/users/*.yaml" //testing
	}
	matches, err := filepath.Glob(pattern)
	if err != nil {
		fmt.Println(err)
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
	fmt.Printf("this is user_list %v", user_list)
	return user_list
}

func Insert_Admins(x UserS) int {
	id := UUID()
	fmt.Println(x.Name)
	fmt.Println(x.Email)
	fmt.Println(x.Date)
	fmt.Println(x.Time)
	fmt.Println(x.Token)
	fmt.Println(x.Password)
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

	res, err := db.Exec("INSERT INTO admin VALUES(?,?,?,?,?,?,?)", id, x.Name, x.Email, x.Date, x.Time, x.Token, x.Password)
	if err != nil {
		fmt.Printf("this is insert err %v", err)
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

	fmt.Printf("insert admin return val %v", ret_val)
	log.Printf("insert admin return val %v", ret_val)
	return ret_val
}

func Create_Admin() {
	alist := parse_admin_list()
	fmt.Println(alist)
	// for _, admin := range alist {
	// 	Insert_Admins(admin)
	// }
}

// func parse_query_string(x string) (string, string, string) {
// 	parts := strings.Split(x, "_")
// 	rawname := strings.Split(parts[0], "=")
// 	return rawname[1], parts[1], parts[2]
// }

// func login(c echo.Context) error {
// 	rawstr := c.QueryString()
// 	t, e, p := parse_query_string(rawstr)

// 	return c.JSON(http.StatusOK, "eat me")
// }

// func logout(c echo.Context) error {

// 	return c.JSON(http.StatusOK, "eat me")
// }
