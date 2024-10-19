package test

import (
	"testing"
	// "todo/internal/db"
	// "todo/internal/migrate"
	// "todo/internal/models"
	// "todo/internal/service"
)

func Test_Update(t *testing.T) {
	// err := db.Init() // INITIATES DATABASE
	// if err != nil {
	// 	t.Error("error: database initalizing failed", err)
	// }

	// err = migrate.Run(db.GetDB()) // RUNS MIGRATION
	// if err != nil {
	// 	t.Error("error: migrate initizaling failed", err)
	// }

	// tdsvc := service.NewItemService(db.GetDB()) // GET GORM SERVICE

	// // T1 TEST MODEL
	// t1 := &models.ToDo{ID: 1000, Title: "thousand", Body: "thousand-body", Finished: false}
	// err = tdsvc.CreateToDo(t1)
	// if err != nil {
	// 	t.Error(err)
	// }

	// defer tdsvc.DeleteToDoByID(1000) // DEFER DELETE SERVICE

	// var todo *models.ToDo // T2 MODEL

	// t1.Title = "new-thousand-five"
	// err = tdsvc.UpdateToDo(t1) // UPDATE T1
	// if err != nil {
	// 	t.Error(err)
	// }

	// todo, err = tdsvc.GetToDoByID(1000) // GETS ID 1000
	// if err != nil {
	// 	t.Error(err)
	// }

	// if todo.Title != t1.Title {
	// 	t.Errorf("update failed we %v need %v", t1.Title, todo.Title)
	// } else {
	// 	t.Log("update success!")
	// }

}
