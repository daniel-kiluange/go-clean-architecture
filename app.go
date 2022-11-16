package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"go-clean-code/routes"
)

func main() {

	engine := gin.Default()
	routes.NewRoutes(engine)
	serverErr := engine.Run()
	if serverErr != nil {
		return
	}
}
