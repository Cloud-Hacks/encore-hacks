package models

import (
	"github.com/afzal442/encore-hacks/pkg/database"
	"time"
)

type Company struct {
	ID       uint   `json:"id" gorm:"primaryKey,autoIncrement" example:"0"`
	Name     string `json:"name" gorm:"not null" example:"Adomate"`
	Industry string `json:"industry" gorm:"not null" example:"Technology"`
	City     string `json:"city" gorm:"not null" example:"Dallas"`

	CreatedAt time.Time `json:"createdAt" example:"2021-01-01T00:00:00Z"`
	UpdatedAt time.Time `json:"updatedAt" example:"2021-01-01T00:00:00Z"`
}

func (c *Company) Create() error {
	return database.DB.Create(&c).Error
}

func GetCompanies() ([]Company, error) {
	var companies []Company
	err := database.DB.Find(&companies).Error
	return companies, err
}

func (c *Company) Update() error {
	return database.DB.Save(&c).Error
}

func (c *Company) Delete() error {
	return database.DB.Delete(&c).Error
}
