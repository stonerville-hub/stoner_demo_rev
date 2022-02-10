package main

import (
	"github.com/my/repo/handler"
	mongoCRUD "github.com/my/repo/server/mongo"
)

func main() {
	// Create a new client and connect to the server
	mongoCRUD.CheckDBConnection()
	mongoCRUD.InsertRecord()
	handler.Controller()
}
