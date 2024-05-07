package models

import (
	"github.com/afzal442/encore-hacks/pkg/database"
)

type Preset struct {
	ID                 uint    `json:"id" gorm:"primaryKey,autoIncrement" example:"0"`
	Type               string  `json:"type" gorm:"not null" example:"low/medium/high"`
	CompanyID          uint    `json:"companyId" gorm:"not null" example:"0"`
	Company            Company `json:"company" gorm:"foreignKey:CompanyID"`
	Revenue            float64 `json:"revenue" gorm:"not null" example:"1000000"`
	COGS               float64 `json:"cogs" gorm:"not null" example:"500000"`
	Depreciation       float64 `json:"depreciation" gorm:"not null" example:"100000"`
	LongTermAssets     float64 `json:"longTermAssets" gorm:"not null" example:"100000"`
	ShortTermAssets    float64 `json:"shortTermAssets" gorm:"not null" example:"100000"`
	LongTermLiability  float64 `json:"longTermLiability" gorm:"not null" example:"100000"`
	ShortTermLiability float64 `json:"shortTermLiability" gorm:"not null" example:"100000"`
	OperatingExpense   float64 `json:"operatingExpense" gorm:"not null" example:"100000"`
	RetainedEarnings   float64 `json:"retainedEarnings" gorm:"not null" example:"100000"`
	YearsInBusiness    uint    `json:"yearsInBusiness" gorm:"not null" example:"5"`

	CreatedAt string `json:"createdAt" example:"2021-01-01T00:00:00Z"`
	UpdatedAt string `json:"updatedAt" example:"2021-01-01T00:00:00Z"`
}

func (p *Preset) Create() error {
	return database.DB.Create(&p).Error
}

func GetPresets() ([]Preset, error) {
	var presets []Preset
	err := database.DB.Find(&presets).Error
	return presets, err
}

func GetPresetsByCompanyID(companyID uint) ([]Preset, error) {
	var presets []Preset
	err := database.DB.Where("company_id = ?", companyID).Find(&presets).Error
	return presets, err
}

func (p *Preset) GetPreset() (Preset, error) {
	var preset Preset
	err := database.DB.Where("id = ?", p.ID).First(&preset).Error
	return preset, err
}

func (p *Preset) Update() error {
	return database.DB.Save(&p).Error
}

func (p *Preset) Delete() error {
	return database.DB.Delete(&p).Error
}
