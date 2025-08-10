package route

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sajal-jayanto/go-template/src/controller"
)

func SetupSampleRoutes(version *fiber.App){
  route := version.Group("/sample")

  route.Post("/", controller.SampleCtrl.CreateSample)
  route.Get("/",  controller.SampleCtrl.GetAllSample)

}
