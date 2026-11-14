package database

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewSqlite() (*gorm.DB, error) {
	dbPath := "./storage/learning.db"

	dir := filepath.Dir(dbPath)

	info, err := os.Stat(dir)

	if err == nil && !info.IsDir() {
		return nil, fmt.Errorf("❌ %s exists but is not a directory", dir)
	}

	if err != nil && !os.IsNotExist(err) {
		return nil, fmt.Errorf("❌ check database directory: %w", err)
	}

	if err := os.MkdirAll(dir, 0750); err != nil {
		return nil, fmt.Errorf("❌ create database directory: %w", err)
	}

	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("❌ connect to sqlite: %w", err)
	}

	sqlDB, err := db.DB()

	if err != nil {
		return nil, fmt.Errorf("❌ get underlying database: %w", err)
	}

	if err := sqlDB.Ping(); err != nil {
		return nil, fmt.Errorf("❌ ping sqlite: %w", err)
	}

	log.Println("✅ sqlite database connected successfully")

	return db, nil
}
