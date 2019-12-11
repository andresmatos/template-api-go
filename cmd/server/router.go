package server

import (
	"template-api-go/pkg/config"
	"template-api-go/pkg/customer"
	"template-api-go/pkg/http/rest"

	"github.com/gin-gonic/gin"
)

func (s *server) MapRoutes(c *config.Config) {

	// Group : v1
	apiV1 := s.router.Group("apptest/api/v1")

	s.healthRoutes(apiV1)
	s.customerRoutes(apiV1)

}

func (s *server) healthRoutes(api *gin.RouterGroup) {
	healthRoutes := api.Group("/health")
	{
		h := rest.NewHealthCtrl()
		healthRoutes.GET("/", h.Ping)
	}
}


func (s *server) customerRoutes(api *gin.RouterGroup) {
	customerRoutes := api.Group("/users")
	{
		var cSvc customer.Service
		s.cont.Invoke(func(u customer.Service) {
			cSvc = u
		})

		c := rest.NewCustomerCtrl(cSvc)

		customerRoutes.GET("/", c.GetAll)
		customerRoutes.POST("/", c.Store)
		customerRoutes.GET("/:id", c.GetByID)
		customerRoutes.PUT("/:id", c.Update)
		customerRoutes.DELETE("/:id", c.Delete)
	}
}