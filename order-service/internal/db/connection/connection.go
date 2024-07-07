package connection

import (
	"sync"

	"gorm.io/gorm"
)

var (
	onlyOnce sync.Once
	client   Client
)

type Client struct {
	PostgresConnection *gorm.DB
}

func New() Client {
	onlyOnce.Do(func() {
		client = Client{
			PostgresConnection: initializePostgres(),
		}
	})

	return client
}
