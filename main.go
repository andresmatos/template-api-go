package main

import (
	"fmt"
	"os"
	"template-api-go/cmd/server"

	"template-api-go/pkg/config"
	"template-api-go/pkg/di"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(-1)
	}
}

func run() error {

	g := gin.Default()
	d := di.BuildContainer()
	c, _ := config.NewConfig()

	g.Use(cors.Default())

	svr := server.NewServer(g, d)
	svr.MapRoutes(c)
	if err := svr.SetupDB(); err != nil {
		return err
	}
	return svr.Start()
}
