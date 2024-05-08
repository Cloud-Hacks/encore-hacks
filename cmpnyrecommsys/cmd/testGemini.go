package main

import (
	"fmt"
	"log"

	"github.com/afzal442/encore-hacks/models"
)

/* type container struct {
	mu   sync.Mutex
	recs []config.Recommendation
} */

func main() {
	/* err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	} */
	fmt.Println("Im online...")
	// database.ConnectDatabase(database.Config())

	/* config.Active = config.ActiveConfig{
		//company that is going middle-of-the-road well
		CompanyName:        "Bob's Electrician Shop",
		Industry:           "Electrician",
		City:               "Dallas",
		Revenue:            float64(900000),
		Cogs:               float64(400000),
		Depreciation:       float64(30000),
		LongTermAssets:     float64(150000),
		ShortTermAssets:    float64(280000),
		LongTermLiability:  float64(70000),
		ShortTermLiability: float64(110000),
		OperatingExpense:   float64(250000),
		RetainedEarnings:   float64(100000),
		YearsInBusiness:    uint(5),
	}

	config.Active = config.ActiveConfig{
		//company that is doing really bad
		CompanyName:        "Joe's Plumbing",
		Industry:           "Plumbing",
		City:               "Richardson",
		Revenue:            float64(400000),
		Cogs:               float64(250000),
		Depreciation:       float64(20000),
		LongTermAssets:     float64(50000),
		ShortTermAssets:    float64(100000),
		LongTermLiability:  float64(80000),
		ShortTermLiability: float64(120000),
		OperatingExpense:   float64(300000),
		RetainedEarnings:   float64(-50000),
		YearsInBusiness:    uint(3),
	}

	config.Active = config.ActiveConfig{
		//company that is doing exceptional
		CompanyName:        "Sue's Bakery",
		Industry:           "Bakery",
		City:               "Southlake",
		Revenue:            float64(800000),
		Cogs:               float64(300000),
		Depreciation:       float64(15000),
		LongTermAssets:     float64(150000),
		ShortTermAssets:    float64(250000),
		LongTermLiability:  float64(50000),
		ShortTermLiability: float64(100000),
		OperatingExpense:   float64(200000),
		RetainedEarnings:   float64(100000),
		YearsInBusiness:    uint(7),
	} */

	//prompt := fmt.Sprintf("Company: %s, Industry: %s, City: %s", config.Active.CompanyName, config.Active.Industry, config.Active.City)
	//
	//cont := container{
	//	recs: []config.Recommendation{},
	//}
	//
	//var wg = &sync.WaitGroup{}
	//for _, rec := range config.RecTypes {
	//
	//	rec := rec
	//	wg.Add(1)
	//	go func(recType config.RecommendationTypes) {
	//		defer wg.Done()
	//		resp := gpt.GetRecommendation(prompt, recType)
	//		for _, r := range resp {
	//			cont.mu.Lock()
	//			cont.recs = append(cont.recs, config.Recommendation{
	//				Title:     r.Title,
	//				Summary:   r.Summary,
	//				Details:   r.Details,
	//				Type:      string(rec),
	//				Completed: false,
	//			})
	//			cont.mu.Unlock()
	//		}
	//	}(rec)
	//}
	//
	//wg.Wait()
	//
	//fmt.Println(len(cont.recs))
	//for _, rec := range cont.recs {
	//	fmt.Println(rec.Title)
	//	fmt.Println(rec.Summary)
	//	fmt.Println(rec.Details)
	//	fmt.Println("----")
	//}

	/* 	err = models.Migrate()
	   	if err != nil {
	   		log.Fatalf("Migration failed: %v", err)
	   	} */
	mv := models.Company{}
	mv.Name = "Bob's Electrician Shop"
	mv.Industry = "Electrician"
	mv.City = "Dallas"
	err := models.AddCompanies(&mv)
	if err != nil {
		log.Fatal(err)
	}
	cmp, err := models.GetCompanies()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(cmp)
	/* 	resp := gemini.GetRiskWeights(config.Active.YearsInBusiness, config.Active.Industry)

	   	for _, res := range resp {
	   		fmt.Println(res.RiskFactor)
	   		fmt.Println(res.RiskWeight)
	   	} */

}
