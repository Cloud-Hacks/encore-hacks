package config

type ActiveConfig struct {
	CompanyName        string  `json:"companyName"`
	Industry           string  `json:"industry"`
	City               string  `json:"city"`
	Revenue            float64 `json:"revenue"`
	Cogs               float64 `json:"cogs"`
	Depreciation       float64 `json:"depreciation"`
	LongTermAssets     float64 `json:"longTermAssets"`
	ShortTermAssets    float64 `json:"shortTermAssets"`
	LongTermLiability  float64 `json:"longTermLiability"`
	ShortTermLiability float64 `json:"shortTermLiability"`
	OperatingExpense   float64 `json:"operatingExpense"`
	RetainedEarnings   float64 `json:"retainedEarnings"`
	YearsInBusiness    uint    `json:"yearsInBusiness"`
	PresetLevel        string  `json:"presetLevel,omitempty"`
}

var Active ActiveConfig = ActiveConfig{
	CompanyName:        "State Farm",
	Industry:           "Insurance",
	City:               "Dallas",
	Revenue:            float64(1000000),
	Cogs:               float64(500000),
	Depreciation:       float64(100000),
	LongTermAssets:     float64(100000),
	ShortTermAssets:    float64(100000),
	LongTermLiability:  float64(100000),
	ShortTermLiability: float64(100000),
	OperatingExpense:   float64(100000),
	RetainedEarnings:   float64(100000),
	YearsInBusiness:    uint(5),
}

var Configs []ActiveConfig = []ActiveConfig{
	{
		//company that is doing exceptional
		CompanyName:        "Sue's Bakery",
		PresetLevel:        "Low Financial Risk",
		Industry:           "Bakery",
		City:               "Southlake",
		Revenue:            float64(800000),
		Cogs:               float64(500000),
		Depreciation:       float64(15000),
		LongTermAssets:     float64(150000),
		ShortTermAssets:    float64(250000),
		LongTermLiability:  float64(50000),
		ShortTermLiability: float64(100000),
		OperatingExpense:   float64(200000),
		RetainedEarnings:   float64(100000),
		YearsInBusiness:    uint(7),
	},
	{
		//company that is doing Average
		CompanyName:        "Sue's Bakery",
		PresetLevel:        "Possible Financial Risk",
		Industry:           "Bakery",
		City:               "Southlake",
		Revenue:            float64(800000),
		Cogs:               float64(300000),
		Depreciation:       float64(50000),
		LongTermAssets:     float64(200000),
		ShortTermAssets:    float64(300000),
		LongTermLiability:  float64(150000),
		ShortTermLiability: float64(250000),
		OperatingExpense:   float64(400000),
		RetainedEarnings:   float64(50000),
		YearsInBusiness:    uint(7),
	},
	{
		//company that is doing Average
		CompanyName:        "Sue's Bakery",
		PresetLevel:        "High Financial Risk",
		Industry:           "Bakery",
		City:               "Southlake",
		Revenue:            float64(800000),
		Cogs:               float64(500000),
		Depreciation:       float64(15000),
		LongTermAssets:     float64(150000),
		ShortTermAssets:    float64(250000),
		LongTermLiability:  float64(100000),
		ShortTermLiability: float64(350000),
		OperatingExpense:   float64(400000),
		RetainedEarnings:   float64(100000),
		YearsInBusiness:    uint(1),
	},
	{
		//company that is doing really bad
		CompanyName:        "Joe's Plumbing",
		PresetLevel:        "High Financial Risk",
		Industry:           "Plumbing",
		City:               "Richardson",
		Revenue:            float64(400000),
		Cogs:               float64(250000),
		Depreciation:       float64(20000),
		LongTermAssets:     float64(100000),
		ShortTermAssets:    float64(100000),
		LongTermLiability:  float64(50000),
		ShortTermLiability: float64(120000),
		OperatingExpense:   float64(150000),
		RetainedEarnings:   float64(10000),
		YearsInBusiness:    uint(3),
	},
	{
		//company that is going middle-of-the-road well
		CompanyName:        "Bob's Electrician Shop",
		PresetLevel:        "Possible Financial Risk",
		Industry:           "Electrician",
		City:               "Dallas",
		Revenue:            float64(650000),
		Cogs:               float64(350000),
		Depreciation:       float64(30000),
		LongTermAssets:     float64(150000),
		ShortTermAssets:    float64(100000),
		LongTermLiability:  float64(70000),
		ShortTermLiability: float64(180000),
		OperatingExpense:   float64(275000),
		RetainedEarnings:   float64(50000),
		YearsInBusiness:    uint(5),
	},
}
