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
	return tdsvc.gdb.Create(&todo).Error
}

// GET FUNCTION
func (tdsvc *toDoService) GetToDoByID(ID uint64) (*models.ToDo, error) {
	var todo models.ToDo
	getRes := tdsvc.gdb.First(&todo, ID)
	if getRes.Error != nil {
		return nil, getRes.Error
	}
	return &todo, nil
}

// SAVE FUNCTION
func (tdsvc *toDoService) UpdateToDo(todo *models.ToDo) error {
	return tdsvc.gdb.Save(&todo).Error
}

// DELETE FUNCTION

func (tdsvc *toDoService) DeleteToDoByID(ID uint64) error {
	return tdsvc.gdb.Delete(&models.ToDo{}, ID).Error
}
