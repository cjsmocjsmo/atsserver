package main

import (
	//	"compress/gzip"
	//
	// "database/sql"
	// "encoding/json"
	"github.com/labstack/echo/v4"
	_ "github.com/mattn/go-sqlite3"
	// "log"
	// "math/rand"
	"net/http"
	// "os"
	// "strconv"
	// "time"
)

func InsertEstimateHandler(c echo.Context) error {

	return c.JSON(http.StatusOK, "Backup Created")
}

func GetAllEstimatesHandler(c echo.Context) error {

	return c.JSON(http.StatusOK, "Backup Created")
}

func GzipEstimatesHandler(c echo.Context) error {

	return c.JSON(http.StatusOK, "Backup Created")
}
