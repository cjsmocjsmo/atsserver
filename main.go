package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	ATS_Logging()
	Create_Tables()
	Insert_Initial_Comments()

	e := echo.New()
	e.Use(middleware.CORS())
	e.GET("/test", TestHandler)
	e.GET("/ins_rev", InsertReviewHandler)
	e.GET("/all_revs", GetAllReviewsHandler)
	e.GET("/backup", DumpGzipHandler)

	// e.Static("/static", "static")      //for pics
	// e.Static("/music", "fsData/music") //for music
	e.Logger.Fatal(e.Start(":8080"))
}
