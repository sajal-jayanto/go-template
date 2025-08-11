package handler

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sajal-jayanto/go-template/src/models"
	"github.com/sajal-jayanto/go-template/src/repository"
	"github.com/sajal-jayanto/go-template/utils"
)

// singleton instance
type sampleHandler struct{}
var SampleHandler sampleHandler

var validate = validator.New()

func (sampleHandler) CreateSample(ctx *fiber.Ctx) error {
	var sample models.Sample
	if err := ctx.BodyParser(&sample); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(
			fiber.Map{"error": "invalid body request"},
		)
	}

	if err := validate.Struct(sample); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(
			fiber.Map{"error": utils.FormatValidationErrors(err)},
		)
	}

	data, err := repository.SampleRepo.Create(sample)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(
			fiber.Map{"error": err.Error()},
		)
	}

	return ctx.Status(fiber.StatusCreated).JSON(data)
}

func (sampleHandler) GetAllSample(ctx *fiber.Ctx) error {
	data, err := repository.SampleRepo.GetAll()
	if err != nil{
		return ctx.Status(400).JSON(
			fiber.Map{"error": err.Error()},
		)
	}
	
	return ctx.Status(200).JSON(data)
}

func(sampleHandler) GetSampleById(ctx *fiber.Ctx) error {
	id, err := utils.GetParamInt(ctx, "id")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(
			fiber.Map{"error": err.Error()},
		)
	}

	data, err := repository.SampleRepo.GetById(id)
	if err != nil{
		return ctx.Status(fiber.StatusInternalServerError).JSON(
			fiber.Map{"error": err.Error()},
		)
	}
	
	return ctx.Status(200).JSON(data)
}