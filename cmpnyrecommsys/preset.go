package comprecsys

import (
	"context"
)

type Preset struct {
	ID                 int     `json:"id" example:"0"`
	Type               string  `json:"type" example:"low/medium/high"`
	CompanyID          int     `json:"companyId" example:"0"`
	Company            Company `json:"company" gorm:"foreignKey:CompanyID"`
	Revenue            float64 `json:"revenue" example:"1000000"`
	COGS               float64 `json:"cogs" example:"500000"`
	Depreciation       float64 `json:"depreciation" example:"100000"`
	LongTermAssets     float64 `json:"longTermAssets" example:"100000"`
	ShortTermAssets    float64 `json:"shortTermAssets" example:"100000"`
	LongTermLiability  float64 `json:"longTermLiability" example:"100000"`
	ShortTermLiability float64 `json:"shortTermLiability" example:"100000"`
	OperatingExpense   float64 `json:"operatingExpense" example:"100000"`
	RetainedEarnings   float64 `json:"retainedEarnings" example:"100000"`
	YearsInBusiness    int     `json:"yearsInBusiness" example:"5"`

	CreatedAt string `json:"createdAt" example:"2021-01-01T00:00:00Z"`
	UpdatedAt string `json:"updatedAt" example:"2021-01-01T00:00:00Z"`
}

func AddPreset(ctx context.Context, p *Preset) error {
	_, err := db.Exec(ctx, `
		INSERT INTO presets (type, company_id, revenue, cogs, depreciation, long_term_assets, short_term_assets, long_term_liability, short_term_liability, operating_expense, retained_earnings, years_in_business)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
	`, p.Type, p.CompanyID, p.Revenue, p.COGS, p.Depreciation, p.LongTermAssets, p.ShortTermAssets, p.LongTermLiability, p.ShortTermLiability, p.OperatingExpense, p.RetainedEarnings, p.YearsInBusiness)
	return err
}

func GetPresets(ctx context.Context) (*[]Preset, error) {
	p := &Preset{}
	rows, err := db.Query(ctx, `SELECT * FROM presets`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var presets []Preset
	for rows.Next() {
		if err := rows.Scan(&p.ID, &p.Type, &p.CompanyID, &p.Revenue, &p.COGS, &p.Depreciation, &p.LongTermAssets, &p.ShortTermAssets, &p.LongTermLiability, &p.ShortTermLiability, &p.OperatingExpense, &p.RetainedEarnings, &p.YearsInBusiness, &p.CreatedAt, &p.UpdatedAt); err != nil {
			return nil, err
		}
		presets = append(presets, *p)
	}
	return &presets, nil
}

func UpdatePreset(ctx context.Context, p *Preset) error {
	_, err := db.Exec(ctx, `
		UPDATE presets
		SET type = $1, company_id = $2, revenue = $3, cogs = $4, depreciation = $5, long_term_assets = $6, short_term_assets = $7, long_term_liability = $8, short_term_liability = $9, operating_expense = $10, retained_earnings = $11, years_in_business = $12
		WHERE id = $13
	`, p.Type, p.CompanyID, p.Revenue,
		p.COGS, p.Depreciation, p.LongTermAssets, p.ShortTermAssets, p.LongTermLiability, p.ShortTermLiability, p.OperatingExpense, p.RetainedEarnings, p.YearsInBusiness, p.ID)
	return err
}

func DeletePreset(ctx context.Context, p *Preset) error {
	_, err := db.Exec(ctx, `DELETE FROM presets WHERE id = $1`, p.ID)
	return err
}
