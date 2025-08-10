package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sajal-jayanto/go-template/src/models"
	"github.com/sajal-jayanto/go-template/src/repository"
)

func CreateSample(ctx *fiber.Ctx) error {
	var sample models.Sample
	if err := ctx.BodyParser(&sample); err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": "invalid request body"})
	}

	data, err := repository.CreateSample(sample)
	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.Status(201).JSON(data)
}

func GetAllSample(ctx *fiber.Ctx) error {
	data, err := repository.GetAllSample()
	if err != nil{
		return ctx.Status(400).JSON(fiber.Map{"error": err.Error()})
	}
	
	return ctx.Status(200).JSON(data)
}