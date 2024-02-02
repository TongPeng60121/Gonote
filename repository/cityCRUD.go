package repository

import (
	models "gonote/models"

	"gorm.io/gorm"
)

// 新建城市
func CreateCities(db *gorm.DB, cities []models.City) {
	db.Create(&cities)
}

// 取得所有城市
func GetAllCities(db *gorm.DB) []models.City {
	var allCities []models.City
	db.Find(&allCities)
	return allCities
}

// 取得多個城市
func GetCities(db *gorm.DB, ids []uint, names []string) ([]models.City, error) {
	var cities []models.City

	// 建構查詢條件
	query := db
	if len(ids) > 0 {
		query = query.Where("id IN (?)", ids)
	}
	if len(names) > 0 {
		query = query.Where("name IN (?)", names)
	}

	if len(ids) == 0 && len(names) == 0 {
		return cities, nil
	}

	if err := query.Find(&cities).Error; err != nil {
		return nil, err
	}

	return cities, nil
}

// 更新城市信息
func UpdateCities(db *gorm.DB, updatedCities []models.City) {
	db.Save(&updatedCities)
}

// 刪除城市
func DeleteCities(db *gorm.DB, ids []uint) {
	db.Delete(&models.City{}, ids)
}
