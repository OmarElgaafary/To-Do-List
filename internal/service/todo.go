package service

import (
	"fmt"
	"strconv"
	"todo/internal/models"

	"gorm.io/gorm"
)

type ToDoService struct { // constructor struct
	gdb *gorm.DB
}

func NewItemService(gormDB *gorm.DB) *ToDoService { // constructor init
	return &ToDoService{gdb: gormDB}
}

// CREATE FUNCTION
func (tdsvc *ToDoService) CreateToDo(todo *models.ToDo) error {
	return tdsvc.gdb.Create(&todo).Error
}

// GET FUNCTION
func (tdsvc *ToDoService) GetToDoByID(strID string) (*models.ToDo, error) {
	ID, err := strconv.ParseUint(strID, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("failed to convert id %v", err)
	}
	var todo models.ToDo
	getRes := tdsvc.gdb.First(&todo, ID)
	if getRes.Error != nil {
		return nil, getRes.Error
	}
	return &todo, nil
}

// SAVE FUNCTION
func (tdsvc *ToDoService) UpdateToDo(todo *models.ToDo) error {
	return tdsvc.gdb.Save(&todo).Error
}

// DELETE FUNCTION

func (tdsvc *ToDoService) DeleteToDoByID(ID uint64) error {
	return tdsvc.gdb.Delete(&models.ToDo{}, ID).Error
}
