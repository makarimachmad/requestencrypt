package main

import (
	"log"
	"requestencrypt/kriptografi"
	"requestencrypt/router"

	"github.com/labstack/echo/v4"
)

func main(){
	e := echo.New()
	
	err := kriptografi.EncryptionInit()
	if err != nil{
		log.Println("error")
	}

	router.SetupRoutes(e)
	e.Logger.Fatal(e.Start(":1323"))
}