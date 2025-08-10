package repository

import (
	"fmt"

	"github.com/sajal-jayanto/go-template/db"
	"github.com/sajal-jayanto/go-template/src/models"
)

func CreateSample(sample models.Sample) (models.Sample, error) {
	result := db.Connection.Create(&sample)
	return sample, result.Error
}

func GetAllSample() ([]models.Sample, error){
	var samples []models.Sample
	result := db.Connection.Find(&samples)
	fmt.Println("DSalfjsdkjf")
	return samples, result.Error
}