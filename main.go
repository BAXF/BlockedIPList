package main

import (
	"github.com/gin-gonic/gin"
)

func blockIP(c *gin.Context) {
	ip := c.Param("ip")
	c.String(200, "Blocked IP: %s", ip)
}

func unblockIP(c *gin.Context) {
	ip := c.Param("ip")
	c.String(200, "Unblocked IP: %s", ip)
}

func checkIP(c *gin.Context) {
	ip := c.Param("ip")
	c.String(200, "Checked IP: %s", ip)
}

func main() {

	router := gin.Default()
	router.POST("/:ip", blockIP)
	router.DELETE("/:ip", unblockIP)
	router.GET("/:ip", checkIP)

	router.Run(":8080")

}
