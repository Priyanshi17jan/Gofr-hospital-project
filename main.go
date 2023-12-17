package main

import (
	"gofr.dev/pkg/gofr"
	"priyanshi_gofr/datastore"
	"priyanshi_gofr/handler"
)

func main() {
	app := gofr.New()
	s := datastore.New()
	h := handler.New(s)
	app.GET("/hospital/{id}", h.GetByID)
	app.POST("/hospital", h.Create)
	app.PUT("/hospital/{id}", h.Update)
	app.DELETE("/hospital/{id}", h.Delete)

	// starting the server on a custom port
	app.Server.HTTP.Port = 9092
	app.Start()
}
