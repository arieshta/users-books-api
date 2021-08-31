package main

import (
	"users-api/config"
	"users-api/routes"
	m "users-api/middlewares"
)

func main() {
	config.InitDB()
	e := routes.New()

	// logger middleware
	m.LogMiddlewares(e)
	e.Logger.Fatal(e.Start(":8000"))
}
