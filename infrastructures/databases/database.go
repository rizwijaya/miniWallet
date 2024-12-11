package database

import (
	"time"

	log "github.com/rizwijaya/miniWallet/infrastructures/logger"

	"github.com/rizwijaya/miniWallet/infrastructures/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDatabase(config config.LoadConfig) *gorm.DB {
	dsn := "host=" + config.Database.Host + " user=" + config.Database.Username + " password=" + config.Database.Password + " dbname=" + config.Database.Name + " port=" + config.Database.Port + " sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	// set connection-pooling
	sqlDb, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}

	log.Info("Setting connection pool database...")
	sqlDb.SetConnMaxIdleTime(10)
	sqlDb.SetMaxOpenConns(100)
	sqlDb.SetConnMaxLifetime(time.Hour)

	// Set the time zone explicitly for the session
	if err := db.Exec("SET TIME ZONE 'Asia/Jakarta'").Error; err != nil {
		log.Fatal("failed to set time zone:", err)
	}

	err = db.AutoMigrate(&Wallet{}, &Transaction{})
	if err != nil {
		log.Fatal("failed to migrate", err)
		return nil
	}
	log.Info("migration successfully done")

	return db
}
