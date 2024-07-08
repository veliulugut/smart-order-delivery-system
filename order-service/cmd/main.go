package main

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"gorm.io/gorm"
)

var (
	conn *gorm.DB
)

func main() {
	fmt.Println("Hello World")
}

func GracefulShutDown(app *fiber.App) {
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	log.Info("Received terminate,graceful shutdown", sig)

	database, err := conn.DB()
	if err != nil {
		log.Error("Database Closing ERROR :", err)
	}

	err = database.Close()
	if err != nil {
		return
	}
	log.Info("Database Closed")

	err = app.Shutdown()
	if err != nil {
		return
	}
}
