package connection

import (
	"fmt"
	"order-service/internal/db/migration"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func initializePostgres() *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s application_name='%s' search_path=%s sslmode='%s' timezone=Europe/Istanbul",
		os.Getenv("POSTGRES_DB_HOST"),
		os.Getenv("POSTGRES_DB_USER"),
		os.Getenv("POSTGRES_DB_PASSWORD"),
		os.Getenv("POSTGRES_DB_NAME"),
		os.Getenv("POSTGRES_DB_PORT"),
		os.Getenv("POSTGRES_DB_APP_NAME"),
		os.Getenv("POSTGRES_DB_APP_NAME"),
		os.Getenv("POSTGRES_DB_SSL_MODE"),
	)

	connection, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		fmt.Println("Error connecting to database: ", err)
		return nil
	}

	migration.Migrate(connection)

	return connection
}
