package test

import (
	"testing"
	"todo/internal/db"
	"todo/internal/migrate"
	"todo/internal/models"
	"todo/internal/service"
)

func Test_Update(t *testing.T) {
	err := db.Init()
	if err != nil {
		t.Error("error: database initalizing failed", err)
	}

	err = migrate.Run(db.GetDB())
	if err != nil {
		t.Error("error: migrate initizaling failed", err)
	}

	tdsvc := service.NewItemService(db.GetDB())

	t1 := &models.ToDo{ID: 1000, Title: "thousand", Body: "thousand-body", Finished: false}
	err = tdsvc.CreateToDo(t1)
	t.Error(err)

	defer tdsvc.DeleteToDoByID(1000)

	var todo *models.ToDo

	// todo, err = tdsvc.GetToDoByID(1000)
	// if err != nil {
	// 	t.Error(err)
	// }

	t1.Title = "new-thousand-five"
	err = tdsvc.UpdateToDo(t1)
	if err != nil {
		t.Error(err)
	}

	todo, err = tdsvc.GetToDoByID(1000)
	if err != nil {
		t.Error(err)
	}

	if todo.Title != t1.Title {
		t.Errorf("update failed we %v need %v", t1.Title, todo.Title)
	} else {
		t.Log("update success!")
	}

}
