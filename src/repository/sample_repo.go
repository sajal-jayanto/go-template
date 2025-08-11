package repository

import (
	"errors"
	"strconv"

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

func (sampleRepo) RemoveById(id int) error {
	result := db.Connection.Delete(&models.Sample{} ,id)
	if result.Error != nil {
    return result.Error
  }
	if result.RowsAffected == 0 {
		return errors.New("sample with id " + strconv.Itoa(id)  + " not found")
	}
	
  return nil
}

func (sampleRepo) UpdateSampleById(id int, newData string) (models.Sample, error) {
	var sample models.Sample
	result := db.Connection.Model(&models.Sample{}).Where("id = ?", id).Update("data" , newData)
	if result.Error != nil {
    return sample, result.Error
  }

	if result.RowsAffected == 0 {
		return sample, errors.New("sample with id " + strconv.Itoa(id)  + " not found")
	}
	
	err := db.Connection.First(&sample, id).Error
	if err != nil {
		return sample, err
	}
  
	return sample, nil
}
