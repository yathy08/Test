package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"main.go/database"
	routes "main.go/handlers"
	"main.go/helper"
)

func Init() {
	helper.LoadEnv()
	database.DBconnect()
	database.InitRedis()
}

func main() {
	Init()
	r := gin.Default()

	r.LoadHTMLGlob("templates/*")

	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	r.GET("/", routes.Login)
	r.POST("/", routes.Postlogin)
	r.GET("/home", routes.UserHome)
	r.GET("/signup", routes.Signup)
	r.POST("/signup", routes.Postsignup)
	r.GET("/logout", routes.Logout)
	r.GET("/admin", routes.Admin)
	r.POST("/admin", routes.PostAdmin)
	r.GET("/valadmin", routes.Valadmin)
	r.GET("/search", routes.Search)
	r.POST("/adduser", routes.AddUser)
	r.GET("/delete/:ID", routes.Delete)
	r.GET("/update/:ID", routes.Update)
	r.POST("/update/:ID", routes.Updateuser)
	r.GET("/block/:ID", routes.Block)
	r.GET("/adminlogout", routes.Adminlogout)
	r.POST("/api/send-otp", routes.SendOTPHandler())
	r.POST("/api/signup/verify", routes.SignupVerify)
	r.GET("/ReportDisaster", routes.GetReportDisaster)
	r.POST("/api/ReportDisaster", routes.PostReportDisaster)
	r.POST("/api/save-report", routes.PostReportDisaster)
	r.GET("/RequestAssistance", routes.GetRequestAssistance)
	r.POST("/api/RequestAssistance", routes.PostRequestAssistance)
	r.GET("/AlertPotentialDisaster", routes.GetAlertPotentialDisaster)
	r.POST("/api/AlertPotentialDisaster", routes.PostAlertPotentialDisaster)
	r.POST("/resources", routes.CreateResource)
	r.GET("/resources/:id", routes.GetResourceByID)
	r.PUT("/resources/:id", routes.UpdateResource)
	r.DELETE("/resources/:id", routes.DeleteResource)
	r.GET("/fetchAllDisasters", routes.GetAllDisasters)
	r.PUT("/updateDisaster/:id", routes.UpdateDisaster)
	r.POST("/addDisaster", routes.AddDisaster)
	r.DELETE("/deleteDisaster/:id", routes.DeleteDisaster)
    r.POST("/allocateResource/:id", routes.AllocateResource)
	r.Run(":3000")
}
