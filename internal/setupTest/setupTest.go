package setupTest


import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"testing"
)

// SetupDB creates an in-memory SQLite database for testing purposes.
// It automatically migrates the provided model and sets up cleanup.
// The database and all its data are automatically destroyed when the test completes.
// There's no need to manually delete data or clean up the database.
func InitDbInMemory(t *testing.T, model interface{}) *gorm.DB {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to open test database: %v", err)
	}

	if err := db.AutoMigrate(model); err != nil {
		t.Fatalf("Failed to migrate model: %v", err)
	}

	t.Cleanup(func() {
		sqlDB, err := db.DB()
		if err != nil {
			t.Logf("Failed to get DB in cleanup: %v", err)
			return
		}
		if err := sqlDB.Close(); err != nil {
			t.Logf("Failed to close test DB in cleanup: %v", err)
		}
	})

	return db
}