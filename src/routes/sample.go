package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sajal-jayanto/go-template/src/handlers"
)

func SetupSampleRoutes(version *fiber.App){
  route := version.Group("/sample")

	route.Post("/", handlers.CreateSample)
  route.Get("/",  handlers.GetAllSample)

}