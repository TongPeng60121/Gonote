package usecases

import (
	models "gonote/models"
	repository "gonote/repository"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// 新建城市
func CreateCities(db *gorm.DB, c *gin.Context) {
	var newCities []models.City
	if err := c.ShouldBindJSON(&newCities); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	repository.CreateCities(db, newCities)
	c.JSON(201, newCities)
}

// 取得所有城市
func GetAllCities(db *gorm.DB, c *gin.Context) {
	allCities := repository.GetAllCities(db)
	c.JSON(200, allCities)
}

// 取得多個城市
func GetCities(db *gorm.DB, c *gin.Context) {
	var query models.Query

	// 使用 ShouldBindJSON 绑定 JSON 数据到 query 变量
	if err := c.ShouldBindJSON(&query); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// 使用 query.IDs 和 query.Names 调用 repository.GetCities 函数
	cities, err := repository.GetCities(db, query.IDs, query.Names)
	if err != nil {
		c.JSON(404, gin.H{"error": "Cities not found"})
		return
	}

	// 返回 JSON 数据
	c.JSON(200, cities)
}

// 更新城市信息
func UpdateCities(db *gorm.DB, c *gin.Context) {
	var updatedCities []models.City
	if err := c.ShouldBindJSON(&updatedCities); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	repository.UpdateCities(db, updatedCities)
	c.JSON(200, gin.H{"message": "Cities updated successfully"})
}

// 删除城市
func DeleteCities(db *gorm.DB, c *gin.Context) {
	var ids []uint
	if err := c.ShouldBindJSON(&ids); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	repository.DeleteCities(db, ids)
	c.JSON(200, gin.H{"message": "Cities deleted successfully"})
}
