package config

type Recommendation struct {
	Title     string  `json:"title"`
	Summary   string  `json:"summary"`
	Details   string  `json:"details"`
	Type      string  `json:"type"`
	Weight    float64 `json:"weight"`
	Completed bool    `json:"completed"`
}

var Recs []Recommendation = []Recommendation{
	{
		Title:     "Tax Recommendation",
		Summary:   "Lie to the IRS",
		Details:   "Lying to the IRS is an extremely effective way to save money on taxes. So long as you don't get caught...",
		Type:      "Tax",
		Weight:    0.5,
		Completed: false,
	},
}

type RecommendationTypes struct {
	Type string `json:"type"`
}

var RecTypes []RecommendationTypes = []RecommendationTypes{
	{
		Type: "Safety",
	},
	{
		Type: "Accident History",
	},
	{
		Type: "Finance",
	},
	{
		Type: "Employee Training",
	},
	{
		Type: "Legal",
	},
	{
		Type: "Tax",
	},
}

var RiskFactorScores map[string]float64 = map[string]float64{
	"Safety":            2,
	"Accident History":  2,
	"Finance":           2,
	"Employee Training": 2,
	"Legal":             2,
	"Tax":               2,
}
