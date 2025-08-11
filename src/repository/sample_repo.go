package repository

import (
	"errors"

	"github.com/sajal-jayanto/go-template/db"
	"github.com/sajal-jayanto/go-template/src/models"
	"gorm.io/gorm"
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

func (sampleRepo) GetById(id int) (models.Sample, error){
	var sample models.Sample
	result := db.Connection.First(&sample, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
    return sample, errors.New("sample not found")
  }
	return sample, result.Error
}