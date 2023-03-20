package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	ATS_Logging()
	Create_Admin()
	Create_Reviews_Tables()
	Create_Estimate_Tables()
	Insert_Initial_Comments()

	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	e.GET("/test", TestHandler)
	e.GET("/ins_rev", InsertReviewHandler)
	e.GET("/all_revs", GetAllReviewsHandler)
	e.GET("/rev_accept", AcceptReviewHandler)
	e.GET("/rev_reject", RejectReviewHandler)

	e.GET("/ins_est", InsertEstimateHandler)
	e.GET("/all_est", GetAllEstimatesHandler)
	e.GET("/comp_est", CompletEstimateHandler)

	e.GET("/revbup", ReviewsGzipHandler)
	e.GET("/estbup", EstimatesGzipHandler)

	e.Static("/static", "./static") // testing
	// e.Static("/static", "/usr/share/ats_server/static") // production for backup.tar.gz
	e.Logger.Fatal(e.Start(":8080"))
}
