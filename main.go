package main

import (
	"github.com/gin-gonic/gin"
	"github.com/haahaakon/what-is-my-ipv6/routes"
)

func main() {
	server := gin.Default()
	routes.RegisterRoutes(server)
	server.Run(":8080")

}
