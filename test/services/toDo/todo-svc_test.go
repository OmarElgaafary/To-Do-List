package test

import (
	"testing"
	"todo/internal/models"
	"todo/internal/service"
	"todo/internal/setupTest"
)

func Test_UpdateToDo(t *testing.T) {
	gdb := setupTest.InitDbInMemory(t,models.ToDo{})
	tdSvc := service.NewItemService(gdb)

	t1 := &models.ToDo{ID: 1000, Title: "thousand", Body: "thousand-body", Finished: false}
	
	err := tdSvc.CreateToDo(t1)
	if err != nil {
		t.Error(err)
	}

	defer tdSvc.DeleteToDoByID(1000)

	var todo *models.ToDo

	t1.Title = "new-thousand-five"
	err = tdSvc.UpdateToDo(t1)
	if err != nil {
		t.Error(err)
	}

	todo, err = tdSvc.GetToDoByID(1000)
	if err != nil {
		t.Error(err)
	}

	if todo.Title != t1.Title {
		t.Errorf("update failed we %v need %v", t1.Title, todo.Title)
	}
}
