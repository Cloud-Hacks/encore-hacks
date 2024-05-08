package models

import (
	"context"
	"time"

	"encore.dev/storage/sqldb"
)

type Company struct {
	ID       uint   `json:"id" gorm:"primaryKey,autoIncrement" example:"0"`
	Name     string `json:"name" gorm:"not null" example:"Adomate"`
	Industry string `json:"industry" gorm:"not null" example:"Technology"`
	City     string `json:"city" gorm:"not null" example:"Dallas"`

	CreatedAt time.Time `json:"createdAt" example:"2024-01-01T00:00:00Z"`
	UpdatedAt time.Time `json:"updatedAt" example:"2024-01-01T00:00:00Z"`
}

/* func Migrate() error {
	return database.DB.AutoMigrate(&Company{})
}

func (c *Company) Create() error {
	return database.DB.Create(&c).Error
}

func AddCompany(name, industry, city string) error {
	company := Company{
		Name:     name,
		Industry: industry,
		City:     city,
	}
	return company.Create()
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
} */

// Define a database named 'url', using the database
// migrations  in the "./migrations" folder.
// Encore provisions, migrates, and connects to the database.
// Learn more: https://encore.dev/docs/primitives/databases
var db = sqldb.NewDatabase("url", sqldb.DatabaseConfig{
	Migrations: "../migrations",
})

func AddCompanies(cmp *Company) error {
	_, err := db.Exec(context.Background(), `
		INSERT INTO companies (name, industry, city)
		VALUES ($1, $2, $3)
	`, cmp.Name, cmp.Industry, cmp.City)
	return err
}

func GetCompanies() ([]Company, error) {
	err := sqldb.QueryRow(context.Background(),
		`SELECT * FROM companies`,
	).Scan(&Company{})
	return []Company{}, err
}
