package data

import "github.com/afzal442/encore-hacks/pkg/config"

// Altman Z-Score model
// Discriminant Analysis Model
// 90.9% Accurate for non-bankrupt firms
// 97.0% Accurate for bankrupt firms

// 0-4 bounds

func CalculateZ() float64 {
	local := config.Active
	totalAssets := local.LongTermAssets + local.ShortTermAssets
	totalLiabilities := local.LongTermLiability + local.ShortTermLiability
	x1 := (local.ShortTermAssets - local.ShortTermLiability) / totalAssets                         // Working Capital/Total Assets
	x2 := local.RetainedEarnings / totalAssets                                                     // Retained Earnings/Total Assets
	x3 := (local.Revenue - local.Cogs - local.Depreciation - local.OperatingExpense) / totalAssets // Earnings Before Interest and Tax/Total Assets
	x4 := (totalAssets - totalLiabilities) / totalLiabilities                                      // Book Value of Equity/Total Liabilities
	x5 := (local.Revenue) / totalAssets                                                            // Sales/Total Assets

	return 0.717*x1 + 0.847*x2 + 3.107*x3 + 0.420*x4 + 0.998*x5
}

func Situation(z float64) string {
	if z < 1.23 {
		return "High Risk"
	} else if z <= 2.90 {
		return "Medium Risk"
	} else {
		return "Low Risk"
	}
}

func InvertZ(z float64) float64 {
	if z > 4 {
		return 0
	}
	if z < 0 {
		return 4
	}
	return 4 - z
}
