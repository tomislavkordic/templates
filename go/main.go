package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

const SLEEP_DURATION_SECONDS = 30

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, World!")
	})

	r.GET("/exit", func(c *gin.Context) {
		os.Exit(1)
	})

	r.GET("/sleep", func(c *gin.Context) {
		time.Sleep(SLEEP_DURATION_SECONDS * time.Second)
		c.String(http.StatusOK, fmt.Sprintf("Slept for %d seconds.", SLEEP_DURATION_SECONDS))
	})

	r.Run()
}
