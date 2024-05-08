package comprecsys

import (
	"context"
	"time"

	"encore.dev/storage/sqldb"
)

type Company struct {
	ID       int    `json:"id" example:"0"`
	Name     string `json:"name" example:"Adomate"`
	Industry string `json:"industry" example:"Technology"`
	City     string `json:"city" example:"Dallas"`

	CreatedAt time.Time `json:"createdAt" example:"2024-01-01T00:00:00Z"`
	UpdatedAt time.Time `json:"updatedAt" example:"2024-01-01T00:00:00Z"`
}

// Define a database named 'cmdb', using the database
// migrations  in the "./migrations" folder.
// Encore provisions, migrates, and connects to the database.
// Learn more: https://encore.dev/docs/primitives/databases
var db = sqldb.NewDatabase("cmdb", sqldb.DatabaseConfig{
	Migrations: "./migrations",
})

func AddCompanies(ctx context.Context, cmp *Company) error {
	_, err := db.Exec(context.Background(), `
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

func UpdateCompany(ctx context.Context, cmp *Company) error {
	_, err := db.Exec(ctx, `
		UPDATE companies
		SET name = $1, industry = $2, city = $3
		WHERE id = $4
	`, cmp.Name, cmp.Industry, cmp.City, cmp.ID)
	return err
}

func DeleteCompany(ctx context.Context, id int) error {
	_, err := db.Exec(ctx, `
		DELETE FROM companies
		WHERE id = $1
	`, id)
	return err
}
