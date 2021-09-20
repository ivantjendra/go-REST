package controllers

import (
	"assignment2/config"
	"assignment2/models"
	"encoding/json"
	"io/ioutil"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Read(c *gin.Context) {
	db := config.GetDB()
	var orders []models.Order
	err := db.Preload("Items").Find(&orders).Error
	if err != nil {
		c.JSON(500, err)
	} else {
		c.JSON(200, orders)
	}
}

func Create(c *gin.Context) {
	db := config.GetDB()
	jsonData, _ := ioutil.ReadAll(c.Request.Body)
	var order models.Order
	json.Unmarshal(jsonData, &order)

	err := db.Create(&order).Error
	if err != nil {
		c.JSON(500, err)
	} else {
		c.JSON(201, order)
	}
}

func Update(c *gin.Context) {
	db := config.GetDB()
	jsonData, _ := ioutil.ReadAll(c.Request.Body)
	var order models.Order
	json.Unmarshal(jsonData, &order)

	err := db.Session(&gorm.Session{FullSaveAssociations: true}).Save(&order).Error
	if err != nil {
		c.JSON(500, err)
	} else {
		c.JSON(200, order)
	}
}

func Delete(c *gin.Context) {
	db := config.GetDB()
	err := db.Delete(&models.Order{}, c.Param("id")).Error
	if err != nil {
		c.JSON(500, err)
	} else {
		c.JSON(200, gin.H{"message": "Deleted successfully"})
	}
}
