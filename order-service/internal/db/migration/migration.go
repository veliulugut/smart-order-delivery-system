package migration

import (
	"log"
	"order-service/internal/db/models"
	"sync"

	"gorm.io/gorm"
)

var onlyOnce sync.Once

func Migrate(connection *gorm.DB) {
	onlyOnce.Do(func() {

		log.Println("Migrating the database...")

		err := connection.AutoMigrate(
			models.Order{},
			models.OrderItem{},
		)

		if err != nil {
			log.Fatalf("Could not migrate: %v", err)
		}

	})

}
