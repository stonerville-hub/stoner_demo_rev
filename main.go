package main

import (
	// auth "simple-api/controllers/v1/auth"
	"github.com/gin-gonic/gin"

	"github.com/my/repo/handler"
	s "github.com/my/repo/server"
	utils "github.com/my/repo/utility"
)
func init(){
	s.LoadConnection()
	s.CheckDBConnection()
}

func main() {
    r := gin.Default()
    r.LoadHTMLGlob("templates/*.html")
	// r.Static("/css", "../templates/css")
	r.GET(utils.HOME, handler.HomePage)
    r.GET(utils.LOADNEWDATA, handler.LoadNewData)
    r.GET(utils.GET_CUSTOMERBYID, handler.GetCustomerByID)

	r.Run()

}
