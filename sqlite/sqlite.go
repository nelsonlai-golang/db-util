package sqlite

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// DB connects to SQLite database
func DB(path string) *gorm.DB {
	return _connect(path)
}

// Memory connects to in-memory SQLite database
func Memory() *gorm.DB {
	return _connect("file::memory:?cache=shared")
}

func _connect(dsn string) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})
	if err != nil {
		log.Println(err)
	}
	return db
}
