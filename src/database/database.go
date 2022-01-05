package database

import (
	"database/sql"
	"fmt"
	"github.com/projects/COVID_Database/src/models"
	"github.com/projects/COVID_Database/src/utils"
	"log"

	"github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	*gorm.DB
}

var DB *gorm.DB

func Init() *gorm.DB {

	dbUri := getDBUri()
	connection := getConn(dbUri)

	sqlDB, sqlErr := sql.Open("postgres", connection)
	if sqlErr != nil {
		log.Fatalf("Error opening sql database")
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)

	gormDB, gormErr := gorm.Open(postgres.New(
		postgres.Config{Conn: sqlDB}),
		&gorm.Config{SkipDefaultTransaction: true},
	)
	if gormErr != nil {
		log.Fatalf("Error opening gorm database")
	}

	DB = gormDB
	return DB

}

func GetDB() *gorm.DB {
	return DB
}

func getConn(dbUri string) string {
	connection, err := pq.ParseURL(dbUri)
	if err != nil {
		log.Fatalf("Connection error")
	}
	connection += " sslmode=require"
	return connection
}

func getDBUri() string {
	host := utils.GetEnvVariable("HOST")
	db := utils.GetEnvVariable("DATABASE")
	user := utils.GetEnvVariable("USER_NAME")
	port := utils.GetEnvVariable("PORT")
	password := utils.GetEnvVariable("PASSWORD")

	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s", user, password, host, port, db)
}

func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(&models.Hospital{})
	db.AutoMigrate(&models.TestLab{})
	db.AutoMigrate(&models.HealthRecord{})
}
