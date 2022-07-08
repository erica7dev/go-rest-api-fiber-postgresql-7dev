package main

import (
	"github.com/erica7dev/fiberapi/pkg/books"
	"github.com/erica7dev/fiberapi/pkg/common/config"
	"github.com/erica7dev/fiberapi/pkg/common/db"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main(){
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config",err)
	}

	app := fiber.New()
	db := db.Init(c.DBUrl)

	books.RegisterRoutes(app, db)

	app.Listen(c.Port)
}