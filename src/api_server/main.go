package main

import (
	v1 "./handler_v1"
	"./libs"
	ext "./middleware"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	libs.InitDB()

	e.POST("/v1/auth/tokens", v1.Token())

	api_v1 := e.Group("/v1")
	api_v1.Use(middleware.JWT([]byte("secret1")))
	api_v1.GET("/auth/version", v1.Version())
	api_v1.POST("/auth/users", v1.User())

	r2 := e.Group("/v1/auth/renew")
	config := middleware.JWTConfig{
		Claims:     &ext.MyClaim{},
		SigningKey: []byte("secret2"),
	}
	r2.Use(middleware.JWTWithConfig(config))
	r2.POST("", v1.ReAuth())

	e.Start(":3000")
}
