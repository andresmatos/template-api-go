package di

import (
	"go.uber.org/dig"
	"template-api-go/pkg/config"
	"template-api-go/pkg/customer"
	"template-api-go/pkg/storage"
	"template-api-go/pkg/storage/orm"
)

var container = dig.New()

func BuildContainer() *dig.Container {
	// config
	container.Provide(config.NewConfig)

	// DB
	container.Provide(storage.NewDb)

	// costumer
	container.Provide(orm.NewCustomerRepo)
	container.Provide(customer.NewCustomerService)
	return container
}

func Invoke(i interface{}) error {
	return container.Invoke(i)
}
