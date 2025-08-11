package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sajal-jayanto/go-template/src/handler"
)

func SetupSampleRoutes(version *fiber.App){
  route := version.Group("/sample")

  route.Post("/", handler.SampleHandler.CreateSample)
  route.Get("/",  handler.SampleHandler.GetAllSample)
  route.Get("/:id", handler.SampleHandler.GetSampleById)
  route.Patch("/:id", handler.SampleHandler.UpdateSampleById)
  route.Delete("/:id", handler.SampleHandler.RemoveSampleById)
}
