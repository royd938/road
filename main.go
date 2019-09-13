package main

import (
	"log"
	"road/road"
)

func main() {

	r := road.New()

	r.GET("/", func(c *road.Context) {
		c.HTML(200, "<h2>Road</h2>")
	})

	r.GET("/test/super", func(c *road.Context) {
		c.TEXT(200, "OK found")
	})

	r.POST("/test/super", func(c *road.Context) {
		c.TEXT(200, "Post found")
	})

	log.Fatal(r.Start(":5000"))
}
