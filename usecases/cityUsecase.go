package usecases

import (
	models "gonote/models"
	repository "gonote/repository"

	"gorm.io/gorm"
)

// 新建城市
func CreateCities(db *gorm.DB, cities []models.City) {
	repository.CreateCities(db, cities)
}

// 取得所有城市
func GetAllCities(db *gorm.DB) []models.City {
	allCities := repository.GetAllCities(db)
	return allCities
}

// 取得多個城市
func GetCities(db *gorm.DB, ids []uint, names []string) ([]models.City, error) {
	cities, err := repository.GetCities(db, ids, names)
	return cities, err
}

// 更新城市信息
func UpdateCities(db *gorm.DB, updatedCities []models.City) {
	repository.UpdateCities(db, updatedCities)
}

// 删除城市
func DeleteCities(db *gorm.DB, ids []uint) {
	repository.DeleteCities(db, ids)
}
