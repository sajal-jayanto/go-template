package api

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/sajal-jayanto/go-template/src/router"
)

type APIServer struct {
	addr string
}

func NewAPIServer(addr string) *APIServer {
	return &APIServer{
		addr: addr,
	}
}

func (server *APIServer) Run() error {
	app := fiber.New() 
	
	v1 := fiber.New()
	app.Mount("/api/v1", v1)

	// Set up sample routers 
	router.SetupSampleRoutes(v1)
	


	// v2 := fiber.New()
	// app.Mount("/api/v2", v2)

	log.Println("Server started on port", server.addr)
	return app.Listen(server.addr)
} 

