package main

import (
	"github.com/gin-gonic/gin"
	"go-clean-code/routes"
)

func main() {
	engine := gin.Default()
	routes.NewRoutes(engine)
	err := engine.Run()
	if err != nil {
		return
	}
}
