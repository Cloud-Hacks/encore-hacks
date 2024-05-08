package model

import (
	"context"
	"time"

	"encore.dev/storage/sqldb"
)

type Company struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Industry string `json:"industry"`
	City     string `json:"city"`

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

// Define a database named 'mydb', using the database
// migrations  in the "./migrations" folder.
// Encore provisions, migrates, and connects to the database.
// Learn more: https://encore.dev/docs/primitives/databases
var db = sqldb.NewDatabase("mydb", sqldb.DatabaseConfig{
	Migrations: "./migrations",
})

func AddCompanies(ctx context.Context, cmp *Company) error {
	_, err := db.Exec(ctx, `
		INSERT INTO companies (name, industry, city)
		VALUES ($1, $2, $3)
	`, cmp.Name, cmp.Industry, cmp.City)
	return err
}

func GetCompanies(ctx context.Context) (*[]Company, error) {
	cmp := &Company{}
	rows, err := db.Query(ctx, `SELECT * FROM companies`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var companies []Company
	for rows.Next() {
		if err := rows.Scan(&cmp.ID, &cmp.Name, &cmp.Industry, &cmp.City, &cmp.CreatedAt, &cmp.UpdatedAt); err != nil {
			return nil, err
		}
		companies = append(companies, *cmp)
	}
	return &companies, nil
}

type GetCompaniesResponse struct {
	Companies []Company `json:"companies"`
}

//encore:api public method=GET path=/companies
func GetCompaniesHandler(ctx context.Context) (*GetCompaniesResponse, error) {
	companies, err := GetCompanies(ctx)
	if err != nil {
		return nil, err
	}
	return &GetCompaniesResponse{Companies: *companies}, nil
}

//encore:api public method=POST path=/companies
func AddCompaniesHandler(ctx context.Context, cmp *Company) error {
	return AddCompanies(ctx, cmp)
}
