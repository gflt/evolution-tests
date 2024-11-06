// package db

// import (
// 	"log"
// 	"testproj/models"

// 	"gorm.io/driver/postgres"
// 	"gorm.io/gorm"

// 	_ "github.com/lib/pq" // Подключаем драйвер для PostgreSQL
// )

// type Database struct {
// 	conn *gorm.DB
// }

// // NewDatabase создает новое соединение с базой данных
// func NewDatabase(connStr string) (*Database, error) {
// 	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Применить миграцию
// 	if err := applyMigration(db); err != nil {
// 		return nil, err
// 	}
// 	return &Database{conn: db}, nil
// }

// func (d *Database) Ping() error {
// 	db, err := d.conn.DB()
// 	if err != nil {
// 		return err
// 	}
// 	return db.Ping()
// }

// // applyMigration применяет SQL-миграцию
// func applyMigration(db *gorm.DB) error {
// 	models := []interface{}{
// 		&models.Answer{},
// 		&models.Question{},
// 		&models.Test{},
// 		&models.User{},
// 	}

// 	for _, model := range models {
// 		if err := db.AutoMigrate(model); err != nil {
// 			return err
// 		}
// 	}
// 	return nil
// }
// func (d *Database) CreateUser(user *models.User) error {
// 	result := d.conn.Create(user)
// 	return result.Error
// }

// func (db *Database) Close() {
// 	sqlDB, err := db.conn.DB()
// 	if err != nil {
// 		log.Printf("Ошибка при получении raw SQL DB: %v", err)
// 		return
// 	}
// 	if err := sqlDB.Close(); err != nil {
// 		log.Printf("Ошибка при закрытии базы данных: %v", err)
// 	}
// }
