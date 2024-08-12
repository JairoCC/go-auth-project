package main

import (
	"exmaple.com/go-auth-project/db"
	"exmaple.com/go-auth-project/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()
	routes.RegisterRoutes(server)
	server.Run(":8080") //localhost:8080
}
