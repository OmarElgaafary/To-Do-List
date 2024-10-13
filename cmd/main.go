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

	// t1 := &models.ToDo{ID: 106, Title: "hundred", Body: "hundred-body", Finished: false}
	log.Println("svc successful")
	//CREATING TO DO
	// err = tdsvc.CreateToDo(t1)
	// if err != nil {
	// 	log.Fatal("error: error occured while creating object", err)
	// }

	var todo *models.ToDo

	todo, err = tdsvc.GetToDoByID(105)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("todo 105 title before update is :", todo.Title)
	}

	todo.Title = "hundred-five"
	err = tdsvc.UpdateToDo(todo)
	if err != nil {
		log.Fatal(err)
	}

	todo, err = tdsvc.GetToDoByID(105)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("todo 105 title AFTER update is :", todo.Title)
	}

	// err = tdsvc.DeleteToDoByID(100)
	// if err != nil {
	// 	log.Fatal(err)
	// }

}
