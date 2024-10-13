package service

import (
	"todo/internal/models"

	"gorm.io/gorm"
)

type toDoService struct { // constructor struct
	gdb *gorm.DB
}

func NewItemService(gormDB *gorm.DB) *toDoService { // constructor init
	return &toDoService{gdb: gormDB}
}

// CREATE FUNCTION
func (tdsvc *toDoService) CreateToDo(todo *models.ToDo) error {
	gormRes := tdsvc.gdb.Create(todo)
	if gormRes.Error != nil {
		return gormRes.Error
	}

	return nil
}

// GET FUNCTION

// SAVE FUNCTION

// DELETE FUNCTION
