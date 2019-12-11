package server

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"template-api-go/pkg/config"
	"template-api-go/pkg/customer"

	"github.com/gin-gonic/gin"
	"go.uber.org/dig"
)

type server struct {
	router *gin.Engine
	cont   *dig.Container
}

func NewServer(e *gin.Engine, c *dig.Container) *server {
	return &server{
		router: e,
		cont:   c,
	}
}

func (s *server) SetupDB() error {
	var db *gorm.DB

	if err := s.cont.Invoke(func(d *gorm.DB) { db = d }); err != nil {
		return err
	}

	db.AutoMigrate(&customer.Customer{})
	return nil
}

// Start start serving the application
func (s *server) Start() error {
	var cfg *config.Config
	if err := s.cont.Invoke(func(c *config.Config) { cfg = c }); err != nil {
		return err
	}
	return s.router.Run(fmt.Sprintf(":%s", cfg.Port))
}
