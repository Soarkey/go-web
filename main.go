package main

import (
	"net/http"

	"go-web/engine"
)

func main() {
	e := engine.New()
	e.GET("/", func(c *engine.Context) {
		c.HTML(http.StatusOK, "<h1>Hello GO</h1>")
	})
	e.GET("/hello", func(c *engine.Context) {
		// path: /hello?name=xxx
		c.String(http.StatusOK, "hello %s, welcome to path %s\n", c.Query("name"), c.Path)
	})
	e.GET("/hello/:name", func(c *engine.Context) {
		// path: /hello/xxx
		c.String(http.StatusOK, "hello %s, welcome to path %s\n", c.Param("name"), c.Path)
	})
	e.GET("/assets/*filepath", func(c *engine.Context) {
		c.JSON(http.StatusOK, engine.H{"filepath": c.Param("filepath")})
	})
	if err := e.Run(":9999"); err != nil {
		return
	}
}
