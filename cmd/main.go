package main

import (
	"log"
	"todo/internal/db"
	"todo/internal/handlers"
	"todo/internal/migrate"
	"todo/internal/server"
	"todo/internal/service"
)

func main() {

	err := db.Init()
	if err != nil {
		log.Fatal("error: database initalizing failed", err)
	} else {
		log.Println("database initalized successfully")
	}

	err = migrate.Run(db.GetDB())
	if err != nil {
		log.Fatal("error: migrate initizaling failed", err)
	} else {
		log.Println("migration successful")
	}

	todoService := service.NewItemService(db.GetDB())

	todoHandler := handlers.NewtoDoHandler(todoService)

	// var x string
	// x = "106"
	// fmt.Printf("type of x is %T", x)
	// y, err := strconv.ParseUint(x, 10, 64)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Print("type of y from reflect is ", reflect.TypeOf(y))
	// fmt.Printf("type of y is %T", y)

	server.RunServer(todoHandler)

}
