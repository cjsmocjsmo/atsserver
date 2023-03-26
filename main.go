package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	ATS_Logging()
	Create_ALL_Tables()
	Create_Admin()
	Insert_All_Initial_Comments()

	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	e.GET("/test", TestHandler)

	e.GET("/login", LoginHandler)
	e.GET("/logout", LogoutHandler)

	e.GET("/ins_rev", InsertReviewHandler)
	e.GET("/all_revs", GetAllReviewsHandler)
	e.GET("/rev_accept", AcceptReviewHandler)
	e.GET("/rev_reject", RejectReviewHandler)
	e.GET("/all_jailed", GetJailedReviewsHandler)

	e.GET("/ins_est", InsertEstimateHandler)
	e.GET("/all_est", GetAllEstimatesHandler)
	e.GET("/comp_est", CompletEstimateHandler)

	e.GET("/revbup", ReviewsGzipHandler)
	e.GET("/estbup", EstimatesGzipHandler)

	e.POST("/upload", UploadHandler)
	e.GET("/getphotobyemail", GetPhotoByEmailHandler)

	e.GET("/counts", CountzHandler)

	e.File("/dbbackup", "static/dbbackup.tar.gz") // testing
	e.File("/revbackup", "static/rev_db.tar.gz")
	e.File("/estbackup", "static/est_db.tar.gz")

	e.Static("/static", "static") // production for backup.tar.gz
	// e.Logger.Fatal(e.Start(":8080")) //testing
	e.Logger.Fatal(e.Start(":80")) //production
}
