package routes

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func getClientIP(context *gin.Context) {
	host := context.Request.Host
	remoteAddr := context.Request.RemoteAddr
	forwarded_for := context.Request.Header.Get("x-forwarded-for")
	content_type := context.Request.Header.Get("content-type")

	fmt.Println(host)
	fmt.Println(remoteAddr)
	fmt.Println(forwarded_for)

	if strings.ToLower(content_type) == "application/json" {
		context.JSON(http.StatusOK, gin.H{"host": host, "remoteAddr": remoteAddr, "forwarded_for": forwarded_for})
		return
	}
	context.String(http.StatusOK, "Host: %s\r\nRemote-IP (TCP): %s\r\nRemote-IP (HTTP): %s\r\nContent-Type: %s\r\n", host, remoteAddr, forwarded_for, content_type)

}
