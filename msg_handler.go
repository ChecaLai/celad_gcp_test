package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"net"
)

func GetLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, address := range addrs {
		// check the address type and if it is not a loopback the display it
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}

func index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": "hello world, celad!","ip":GetLocalIP()})
}
