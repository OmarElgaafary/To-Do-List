package migrate

import (
	"todo/internal/models"

	"gorm.io/gorm"
)

func Run(gDB *gorm.DB) error {
	return gDB.AutoMigrate(&models.ToDo{})
}
