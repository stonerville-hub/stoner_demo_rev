package main

import (
	// auth "simple-api/controllers/v1/auth"
	// "time"
	// "github.com/chenyahui/gin-cache"
	// "github.com/chenyahui/gin-cache/persist"
	"github.com/gin-gonic/gin"
	// "github.com/go-redis/redis/v8"
	"github.com/my/repo/handler"
	s "github.com/my/repo/server"
	u "github.com/my/repo/utility"
)
func init(){
	s.LoadConnection()
	s.CheckDBConnection()
}

func main() {
    r := gin.Default()
    r.LoadHTMLGlob("templates/*.html")
	// r.Static("/css", "../templates/css")
	r.GET(u.HOME, handler.HomePage)
    r.GET(u.LOADNEWDATA, handler.LoadNewData)
    r.GET(u.GET_CUSTOMERBYID, handler.GetCustomerByID)

	r.Run(":8080")

}
