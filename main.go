package main

import (
	"backend/controllers"
	_ "backend/resource"
	_ "backend/utils"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	r := gin.Default()
	r.GET("/report", controllers.ReportCon.GetReport)
	r.GET("/search", controllers.ReportCon.Search)
	r.Run()
}
