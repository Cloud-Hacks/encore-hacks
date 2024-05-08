package comprecsys

import (
	"context"
	"fmt"
	"sync"

	"github.com/afzal442/encore-hacks/pkg/config"
	"github.com/afzal442/encore-hacks/pkg/data"
	"github.com/afzal442/encore-hacks/pkg/dto"
	"github.com/afzal442/encore-hacks/pkg/gemini"
	"github.com/google/generative-ai-go/genai"
)

type Config struct {
	CompanyName string   `json:"companyName"`
	Industry    string   `json:"industry"`
	City        string   `json:"city"`
	Presets     []Preset `json:"presets"`
}

type getConfigsResponse struct {
	Configs []config.ActiveConfig `json:"configs"`
}

//encore:api public method=GET path=/configs
func getConfigs(ctx context.Context) (*getConfigsResponse, error) {
	return &getConfigsResponse{
		Configs: []config.ActiveConfig{
			config.Active,
		},
	}, nil
}

type container struct {
	recs []config.Recommendation
	mu   sync.Mutex
}

var Loading = false

//encore:api public method=POST path=/set-config
func setActiveConfig(ctx context.Context, req *config.ActiveConfig) (dto.MessageResponse, error) {

	config.Active = *req

	prompt := fmt.Sprintf("Company: %s, Industry: %s, City: %s", config.Active.CompanyName, config.Active.Industry, config.Active.City)

	cont := container{
		recs: []config.Recommendation{},
	}

	var wg = &sync.WaitGroup{}
	for _, rec := range config.RecTypes {

		rec := rec
		wg.Add(1)
		go func(recType config.RecommendationTypes) {
			defer wg.Done()
			resp := gemini.GetRecommendation(prompt, recType)
			for _, r := range resp {
				cont.mu.Lock()
				cont.recs = append(cont.recs, config.Recommendation{
					Title:     r.Title,
					Summary:   r.Summary,
					Details:   r.Details,
					Type:      rec.Type,
					Weight:    r.Weight,
					Completed: false,
				})
				cont.mu.Unlock()
			}
		}(rec)
	}
	wg.Wait()

	config.Recs = cont.recs
	data.Weights = nil
	data.BaselineRisk = 0

	Loading = false

	return dto.MessageResponse{
		Message: "Active config updated successfully",
	}, nil
}

// For the user show the current config

type overall struct {
	RiskValue          float64             `json:"riskValue"`
	ZValue             float64             `json:"zValue"`
	BankruptcyRisk     string              `json:"bankruptcyRisk"`
	WeightedBaseline   float64             `json:"weightedBaseline"`
	WeightedRecs       float64             `json:"weightedRecs"`
	WeightedBankruptcy float64             `json:"weightedBankruptcy"`
	Active             config.ActiveConfig `json:"active"`
}

//encore:api public method=GET path=/active-config
func getActiveConfig(ctx context.Context) (overall, error) {
	var o overall
	o.Active = config.Active
	o.RiskValue = data.CalculateRisk()
	o.ZValue = data.CalculateZ()
	o.BankruptcyRisk = data.Situation(o.ZValue)
	if o.ZValue < 1.23 {
		added := false
		for _, rec := range config.Recs {
			if rec.Title == "Get Liability Insurance" {
				added = true
			}
		}
		if !added {
			config.Recs = append(config.Recs, config.Recommendation{
				Title:     "Get Liability Insurance",
				Summary:   "Protect Against the Unexpected: Liability Insurance as Your Business Safety Net",
				Details:   "For small businesses navigating financial challenges, liability insurance is a critical tool for risk management. It offers protection against the high costs associated with legal claims and lawsuits, which can be the difference between survival and closure. When facing potential bankruptcy, liability insurance acts as a shield, guarding your hard-earned business against the threats that could exacerbate financial strain.",
				Type:      "Legal",
				Weight:    1,
				Completed: false,
			})
		}
	}

	o.WeightedBankruptcy = data.BankruptcyWeight
	o.WeightedBaseline = data.BaselineWeight
	o.WeightedRecs = data.RecWeight

	if Loading {
		o.RiskValue = -1
	}

	return o, nil
}

type configRecommendation struct {
	ConfRec []config.Recommendation `json:"recommendations"`
}

//encore:api public method=POST path=/recommendations
func getRecommendations(ctx context.Context, category config.RecommendationTypes) (*configRecommendation, error) {
	resp := gemini.GetRecommendation("Give me recommendation for this category", category)
	recommendations := make([]config.Recommendation, len(resp))
	fmt.Println("Recommendations: ", resp)
	for i, r := range resp {
		recommendations[i] = config.Recommendation{
			Title:   r.Title,
			Summary: r.Summary,
			Details: r.Details,
			Type:    category.Type,
			Weight:  r.Weight,
		}
	}
	return &configRecommendation{ConfRec: recommendations}, nil
}

type genAIcontent struct {
	Msg string `json:"msg"`
}

//encore:api public method=POST path=/chat-completion
func ChatCompletion(ctx context.Context, req genAIcontent) (genAIcontent, error) {
	resp := gemini.Conversation(genai.Text(req.Msg))
	return genAIcontent{Msg: string(resp)}, nil
}
