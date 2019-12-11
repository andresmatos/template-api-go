package storage

import (
	"fmt"
	"template-api-go/pkg/config"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func NewDb(c *config.Config) (*gorm.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		c.DB.Postgres.Host,
		c.DB.Postgres.Port,
		c.DB.Postgres.UserName,
		c.DB.Postgres.Password,
		c.DB.Postgres.Database)

	db, err := gorm.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}
	return db, nil
}
