package main

import (
	"fmt"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	_ "github.com/lib/pq"
	"github.com/rdyc/go-echo/driver"
	"github.com/rdyc/go-echo/routers"
)

const (
	host   = "localhost"
	port   = "5432"
	user   = "appuser"
	pass   = "P@ssw0rd!"
	dbname = "IdentityServer4Admin"
)

func main() {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")

	if dbHost == "" {
		dbHost = host
	}
	if dbPort == "" {
		dbPort = port
	}
	if dbName == "" {
		dbName = dbname
	}
	if dbUser == "" {
		dbUser = user
	}
	if dbPass == "" {
		dbPass = pass
	}

	db, err := driver.ConnectSQL(host, port, user, pass, dbname)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	// create new echo instance
	e := echo.New()

	// register middlewares
	// e.Use(middleware.Logger())
	// e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))

	// router group
	v1 := e.Group("/v1")

	// register routers with group
	routers.UserRouter(v1, db)

	// error handler on startup
	e.Logger.Fatal(e.Start(":4444"))
}