package main

import (
	"api/server"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
	"regexp"
)

func main() {
	logfile, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(logfile, os.Stdout)

	r := gin.New()
	r.Use(CorsMiddleware())
	r.GET("version", func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	})

	v1 := r.Group("/api")

	v1.Any("/file-manager", server.FileManager)

	v1.GET("/get-detail", server.GetDetail)

	err := r.Run(GetRunPort())

	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(0)
	}
}

func CorsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")
		var filterHost = [...]string{origin}
		// filterHost 做过滤器，防止不合法的域名访问
		var isAccess = false
		for _, v := range filterHost {
			match, _ := regexp.MatchString(v, origin)
			if match {
				isAccess = true
			}
		}
		if isAccess {
			// 核心处理方式
			c.Header("Access-Control-Allow-Origin", "*")
			c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
			c.Header("Access-Control-Allow-Methods", "GET, OPTIONS, POST, PUT, DELETE")
			c.Set("content-type", "application/json")
		}
		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "Options Request!")
		}

		c.Next()
	}
}

func GetRunPort() string {
	port := os.Getenv("GOPORT")
	if port == "" {
		port = "8019"
	}
	return fmt.Sprintf(":%s", port)
}