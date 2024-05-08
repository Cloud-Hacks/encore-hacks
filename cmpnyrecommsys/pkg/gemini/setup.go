package gemini

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/afzal442/encore-hacks/pkg/config"
	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

type RecommendationResponse struct {
	Title    string  `json:"title"`
	Summary  string  `json:"summary"`
	Details  string  `json:"details"`
	Category string  `json:"category"`
	Weight   float64 `json:"weight"`
}

type ReccomendationResponseTwo struct {
	Recommendations []RecommendationResponse `json:"recommendations"`
}

func GetRecommendation(msg string, category config.RecommendationTypes) []RecommendationResponse {
	ctx := context.Background()
	// Access your API key as an environment variable (see "Set up your API key" above)
	client, err := genai.NewClient(ctx, option.WithAPIKey("AIzaSyBM221b9rNbUmpIOIKAOb9xhhyv0dGnMs0"))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	// For text-only input, use the gemini-pro model
	model := client.GenerativeModel("gemini-pro")
	// Initialize the chat
	cs := model.StartChat()
	cs.History = []*genai.Content{
		&genai.Content{
			Parts: []genai.Part{
				genai.Text("Hello, I have a company input."),
			},
			Role: "user",
		},
		&genai.Content{
			Parts: []genai.Part{
				genai.Text("You are a recommendation assistant API. I am going to give you a company's message and category and you need to provide me a list of recommendations. Return a list of objects with the keys: 'title', 'summary', 'details', 'category', and 'weight' like\n\n[\n{\"title\": \"LegalAdvice.com\",\n    \"summary\": \"Get personalized and affordable legal advice from licensed attorneys.\",\n    \"details\": \"With LegalAdvice.com, you can get affordable and personalized legal advice from licensed attorneys. Get help with common legal issues like family law, criminal defense, wills and trusts, and more. LegalAdvice.com offers a variety of packages to fit your budget, and you can be confident that you're getting high-quality legal advice from a qualified attorney.\",\n     */\"category\": \"Legal\",\n    \"weight\": 0.9,\n}] should not contain an extra opening and closing bracket [[ ]] also any  */ as well as invalid character ']' after object key:value pair"),
			},
			Role: "model",
		},
	}

	resp, err := cs.SendMessage(ctx, genai.Text(msg + fmt.Sprintf(" Category: %s", category)))
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(resp.Candidates[0].Content.Parts)

	response := resp.Candidates[0].Content.Parts[0]
	var recommendations []RecommendationResponse
	err = json.Unmarshal([]byte(fmt.Sprint(response)), &recommendations)
	if err != nil {
		fmt.Println("ChatCompletion error: Failing back to method 2")
		var recs ReccomendationResponseTwo
		err = json.Unmarshal([]byte(fmt.Sprint(response)), &recs)
		if err != nil {
			fmt.Printf("ChatCompletion error: %v\n", err)
			fmt.Printf("Given response:\n %s\n\n", response)
			return nil
		}
		return recs.Recommendations
	}
	return recommendations
}

type BaselineRiskResponse struct {
	BaselineRisk float64 `json:"baseline_risk"`
}

func GetBaselineRisk(age uint, industry string) BaselineRiskResponse {
	ctx := context.Background()
	// Access your API key as an environment variable (see "Set up your API key" above)
	client, err := genai.NewClient(ctx, option.WithAPIKey("AIzaSyBM221b9rNbUmpIOIKAOb9xhhyv0dGnMs0"))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	// For text-only input, use the gemini-pro model
	model := client.GenerativeModel("gemini-pro")
	// Initialize the chat
	cs := model.StartChat()
	cs.History = []*genai.Content{
		&genai.Content{
			Parts: []genai.Part{
				genai.Text("Hello, I am interested in risk factors for a company."),
			},
			Role: "user",
		},
		&genai.Content{
			Parts: []genai.Part{
				genai.Text("You are a risk assessor assistant API that returns JSON ONLY. I am going to give you a company's age and industry and you need to provide me a float baseline risk factor on a 0.1 - 1.9 scale with 2 being the riskiest. Return a JSON with the key: 'baseline_risk'. Return JSON only."),
			},
			Role: "model",
		},
	}
	resp, err := cs.SendMessage(ctx, genai.Text(fmt.Sprintf("Age: %d, Industry: %s", age, industry)))
	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return BaselineRiskResponse{
			BaselineRisk: 1.0,
		}
	}

	response := resp.Candidates[0].Content.Parts[0].(genai.Text)
	fmt.Println(response)
	var baselineRisk BaselineRiskResponse
	err = json.Unmarshal([]byte(response), &baselineRisk)
	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		fmt.Printf("Given response:\n %s\n\n", response)
		return BaselineRiskResponse{
			BaselineRisk: 1.0,
		}
	}

	return baselineRisk
}

type RiskResponse struct {
	RiskFactor string  `json:"risk_factor"`
	RiskWeight float64 `json:"risk_weight"`
}

type RiskResponseTwo struct {
	RiskWeights []RiskResponse `json:"risk_weights"`
}

type RiskResponseThree struct {
	RiskWeights []RiskResponse `json:"risk_factors"`
}

func GetRiskWeights(age uint, industry string) []RiskResponse {
	ctx := context.Background()
	// Access your API key as an environment variable (see "Set up your API key" above)
	client, err := genai.NewClient(ctx, option.WithAPIKey("AIzaSyBM221b9rNbUmpIOIKAOb9xhhyv0dGnMs0"))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	// For text-only input, use the gemini-pro model
	model := client.GenerativeModel("gemini-pro")
	// Initialize the chat
	cs := model.StartChat()

	cs.History = []*genai.Content{
		&genai.Content{
			Parts: []genai.Part{
				genai.Text("Hello, I am interested in risk factors for a company."),
			},
			Role: "user",
		},
		&genai.Content{
			Parts: []genai.Part{
				genai.Text("You are a risk assessor assistant API that returns JSON ONLY. I am going to give you a company's age, industry and risk factors and you need to provide me the weights for the risk factors that will help me combine it into a single risk metric. The risk weights should sum to 1.0. Return a JSON list of objects with the keys: 'risk_factor' and 'risk_weight'. Finance should always be rated at 0.5. Return JSON only."),
			},
			Role: "model",
		},
	}
	riskFactors := ""
	for _, factor := range config.RecTypes {
		riskFactors += factor.Type + ", "
	}
	resp, err := cs.SendMessage(ctx, genai.Text(fmt.Sprintf("Age: %d, Industry: %s, Risk Factors: Safety, Security Measures, Accident History, Employee Training, Legal", age, industry)))

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return []RiskResponse{
			{
				RiskFactor: "Finance",
				RiskWeight: 0.5,
			},
			{
				RiskFactor: "Safety",
				RiskWeight: 0.1,
			},
			{
				RiskFactor: "Security Measures",
				RiskWeight: 0.1,
			},
			{
				RiskFactor: "Accident History",
				RiskWeight: 0.1,
			},
			{
				RiskFactor: "Employee Training",
				RiskWeight: 0.1,
			},
			{
				RiskFactor: "Legal",
				RiskWeight: 0.1,
			},
		}
	}

	var riskWeights []RiskResponse
	response := resp.Candidates[0].Content.Parts[0].(genai.Text)
	err = json.Unmarshal([]byte(response), &riskWeights)
	if err != nil {
		fmt.Println("Weight ChatCompletion error: Failing back to method 2")
		fmt.Println(response)
		var recs RiskResponseTwo
		err = json.Unmarshal([]byte(response), &recs)
		if err != nil {
			fmt.Println("Weight ChatCompletion error: Failing back to method 3")
			var recs RiskResponseThree
			err = json.Unmarshal([]byte(response), &recs)
			if err != nil {
				fmt.Printf("Weight ChatCompletion error: %v\n", err)
				fmt.Printf("Given response:\n %s\n\n", response)
				return []RiskResponse{
					{
						RiskFactor: "Finance",
						RiskWeight: 0.5,
					},
					{
						RiskFactor: "Safety",
						RiskWeight: 0.1,
					},
					{
						RiskFactor: "Security Measures",
						RiskWeight: 0.1,
					},
					{
						RiskFactor: "Accident History",
						RiskWeight: 0.1,
					},
					{
						RiskFactor: "Employee Training",
						RiskWeight: 0.1,
					},
					{
						RiskFactor: "Legal",
						RiskWeight: 0.1,
					},
				}
			}
		}
		return recs.RiskWeights
	}

	return riskWeights
}

func Conversation(pastMessages genai.Text) genai.Text {
	client, err := genai.NewClient(context.Background(), option.WithAPIKey("AIzaSyBM221b9rNbUmpIOIKAOb9xhhyv0dGnMs0"))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	// For text-only input, use the gemini-pro model
	model := client.GenerativeModel("gemini-pro")
	// Initialize the chat
	cs := model.StartChat()
	cs.History = []*genai.Content{
		&genai.Content{
			Parts: []genai.Part{
				pastMessages,
			},
			Role: "user",
		},
		&genai.Content{
			Parts: []genai.Part{
				genai.Text("You are a Insurance Agent at State Farm here to help answer questions and provide guidance to your clients."),
			},
			Role: "model",
		},
	}

	resp, err := cs.SendMessage(context.Background(), genai.Text("Plz help me with the risk factors for a company."))
	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return genai.Text("I'm sorry, I don't have an answer for that.")
	}

	return resp.Candidates[0].Content.Parts[0].(genai.Text)
}
