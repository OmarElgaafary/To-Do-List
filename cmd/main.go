package main

import (
	"log"
	db "todo/internal/db"
	"todo/internal/migrate"
	"todo/internal/models"
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

	tdsvc := service.NewItemService(db.GetDB())
	t1 := &models.ToDo{Title: "To-Do-Title", Body: "To-Do-Body", Finished: false}
	log.Println("svc successful")
	err = tdsvc.CreateToDo(t1)
	if err != nil {
		log.Fatal("error: error occured while creating object", err)
	}

}
