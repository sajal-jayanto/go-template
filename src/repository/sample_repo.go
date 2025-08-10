package repository

import (
	"github.com/sajal-jayanto/go-template/db"
	"github.com/sajal-jayanto/go-template/src/models"
)

// singleton instance
type sampleRepo struct{}
var SampleRepo sampleRepo

func (sampleRepo) Create(sample models.Sample) (models.Sample, error) {
	result := db.Connection.Create(&sample)
	return sample, result.Error
}

func (sampleRepo) GetAll() ([]models.Sample, error){
	var samples []models.Sample
	result := db.Connection.Find(&samples)
	return samples, result.Error
}