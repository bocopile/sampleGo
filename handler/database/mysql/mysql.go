package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"sampleGo/config"
	"sync"
	"time"
)

var (
	dbManagerInstance *DBManager
	once              sync.Once
)

type DBManager struct {
	DB *sql.DB
}

func initDBManager() {
	cfg, err := config.LoadConfig("config.json")
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		cfg.Database.Username, cfg.Database.Password, cfg.Database.Host, cfg.Database.Port, cfg.Database.DBName)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}

	db.SetMaxOpenConns(100)
	db.SetMaxIdleConns(10)
	db.SetConnMaxLifetime(time.Minute * 5)

	dbManagerInstance = &DBManager{DB: db}
}

// GetDBManager는 싱글톤 DBManager 인스턴스를 반환합니다.
func GetDBManager() *DBManager {
	once.Do(initDBManager)
	return dbManagerInstance
}

func (manager *DBManager) GetDB() *sql.DB {
	return manager.DB
}

func (manager *DBManager) Close() {
	if err := manager.DB.Close(); err != nil {
		log.Fatalf("Error closing database connection: %v", err)
	}
}
