package main

import (
	// "github.com/labstack/echo/v4"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	// "encoding/json"
	"fmt"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"path/filepath"
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
	pattern := "/usr/share/ats_server/users/*.yaml"
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
	fmt.Println(user_list)
	return user_list
}

func Insert_Admins(x UserS) int {
	id := UUID()
	fmt.Println(x.Name)
	bpath := os.Getenv("ATS_PATH") + "/atsinfo.db"
	db, err := sql.Open("sqlite3", bpath) // production

	if err != nil {
		db, err2 := sql.Open("sqlite3", "atsinfo.db") //testing
		if err2 != nil {
			log.Fatal(err2)
		}
		defer db.Close()

		res, err := db.Exec("INSERT INTO admin VALUES(?,?,?,?,?,?,?)", &id, &x.Name, &x.Email, &x.Date, &x.Time, &x.Token, &x.Password)
		if err != nil {
			log.Println(err)
			log.Println("admin insert has failed")
		}
		var ret_val int
		_, err = res.LastInsertId()
		if err != nil {
			log.Println(err)
			ret_val = 1
		} else {
			ret_val = 0
		}

		log.Println("insert admin complete")
		return ret_val
	}

	defer db.Close()

	res, err := db.Exec("INSERT INTO admin VALUES(?,?,?,?,?,?,?)", &id, &x.Name, &x.Email, &x.Date, &x.Time, &x.Token, &x.Password)
	if err != nil {
		log.Println(err)
		log.Println("admin insert has failed")
	}
	var ret_val int
	_, err = res.LastInsertId()
	if err != nil {
		log.Println(err)
		ret_val = 1
	} else {
		ret_val = 0
	}

	log.Println("insert admin complete")
	return ret_val
}

func Create_Admin() {
	alist := parse_admin_list()
	for _, admin := range alist {
		Insert_Admins(admin)
	}
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
