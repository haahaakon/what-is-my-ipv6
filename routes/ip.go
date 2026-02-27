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
	fmt.Println(content_type)

	if strings.ToLower(content_type) == "application/json" {
		context.JSON(http.StatusOK, gin.H{"ipv6": forwarded_for})
		return
	}
	context.String(http.StatusOK, "IP: %s\r\n", forwarded_for)

}
