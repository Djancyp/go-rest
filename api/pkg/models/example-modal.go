package exampleModal

import (
	"github.com/Djancyp/go-rest/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Example struct {
	gorm.Model
	Name string `json:"name"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Example{})
}

func (e *Example) Create() *Example {
	db.NewRecord(e)
	db.Create(&e)
	return e
}

func GetAllExamples() []Example {
	var Example []Example
	db.Find(&Example)
	return Example
}

func GetExampleById(Id int64) (*Example, *gorm.DB) {
	var getExample Example
	db := db.Where("id = ?", Id).First(&getExample)
	return &getExample, db
}
func DeleteExampleById(Id int64) (*Example, *gorm.DB) {
	var deleteExample Example
	db := db.Where("id = ?", Id).Delete(&deleteExample)
	return &deleteExample, db
}

func (e *Example) UpdateExample(Id int64) (*Example, *gorm.DB) {
	var updateExample Example
	db.Model(&updateExample).Updates(e)
	db.Where("id = ?", Id).First(&updateExample)
	db.Save(&e)

	return &updateExample, db
}
