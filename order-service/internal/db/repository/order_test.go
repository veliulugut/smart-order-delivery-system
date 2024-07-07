package repository_test

import (
	"order-service/internal/db/models"
	"order-service/internal/db/repository"
	"testing"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&models.Order{}, &models.OrderItem{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

/*
	func Test_CreateOrder(t *testing.T) {
		db, err := setupTestDB()
		assert.NoError(t, err)

		repo := repository.New(db)

		order := &models.Order{
			ID:           1,
			UserID:       1,
			RestaurantID: 1,
			Status:       "pending",
			TotalPrice:   10.0,
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
			OrderItems: []models.OrderItem{
				{
					ID:        1,
					OrderID:   1,
					ProductID: 1,
					Quantity:  1,
					Price:     10.0,
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				},
			},
		}

		if err := repo.CreateOrder(order); err != nil {
			t.Fatalf("Sipariş oluşturma hatası: %v", err)
		}

		if order.ID == 0 || order.OrderItems[0].ID == 0 {
			t.Fatal("ID'ler atanmadı")
		}

		var fetchedOrder models.Order
		if err := db.Preload("OrderItems").First(&fetchedOrder, order.ID).Error; err != nil {
			t.Fatalf("Sipariş getirme hatası: %v", err)
		}

		if fetchedOrder.ID != order.ID || len(fetchedOrder.OrderItems) != 1 || fetchedOrder.OrderItems[0].ID != order.OrderItems[0].ID {
			t.Fatal("Veri doğrulama hatası")
		}

}
*/
func Test_CreateOrder(t *testing.T) {
	db, err := setupTestDB()
	if err != nil {
		t.Fatalf("Veritabanı kurulum hatası: %v", err)
	}

	repo := repository.New(db)
	order := &models.Order{
		UserID:       1,
		RestaurantID: 1,
		Status:       "pending",
		TotalPrice:   10.0,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
		OrderItems: []models.OrderItem{
			{ProductID: 1, Quantity: 1, Price: 10.0, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		},
	}

	if err := repo.CreateOrder(order); err != nil {
		t.Fatalf("Sipariş oluşturma hatası: %v", err)
	}

	if order.ID == 0 || order.OrderItems[0].ID == 0 {
		t.Fatal("ID'ler atanmadı")
	}

	var fetchedOrder models.Order
	if err := db.Preload("OrderItems").First(&fetchedOrder, order.ID).Error; err != nil {
		t.Fatalf("Sipariş getirme hatası: %v", err)
	}

	if fetchedOrder.ID != order.ID || len(fetchedOrder.OrderItems) != 1 || fetchedOrder.OrderItems[0].ID != order.OrderItems[0].ID {
		t.Fatal("Veri doğrulama hatası")
	}
}
