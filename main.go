package main

import(
	"requestencrypt/router"
	"github.com/labstack/echo/v4"
)

func main(){
	e := echo.New()
	
	router.SetupRoutes(e)
	e.Logger.Fatal(e.Start(":1323"))
}