package gemini

import (
	"Brilliant/config"
	"context"
	"fmt"
	"log"

	"github.com/google/generative-ai-go/genai"
	option "google.golang.org/api/option"
)

func GetTextModel() *genai.GenerativeModel {
	// Load environment variables
	env, err := config.Load()
	if err != nil {
		panic(err)
	}
	apiKey := os.Getenv("GEMINI_API_KEY")
if apiKey == "" {
    log.Fatal("API Key not found. Please set the GEMINI_API_KEY environment variable.")
}
	fmt.Println("API Key: ", apiKey)

	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		log.Fatal(err)
	}
	// defer client.Close()

	model := client.GenerativeModel("gemini-pro")
	model.SafetySettings = []*genai.SafetySetting{
		{
			Category:  genai.HarmCategoryHarassment,
			Threshold: genai.HarmBlockMediumAndAbove,
		},
		{
			Category:  genai.HarmCategoryHateSpeech,
			Threshold: genai.HarmBlockMediumAndAbove,
		},
		{
			Category:  genai.HarmCategorySexuallyExplicit,
			Threshold: genai.HarmBlockMediumAndAbove,
		},

		{
			Category:  genai.HarmCategoryDangerousContent,
			Threshold: genai.HarmBlockMediumAndAbove,
		},
	}
	return model
}

func GetImageModel() *genai.GenerativeModel {
	// Load environment variables
	env, err := config.Load()
	if err != nil {
		panic(err)
	}
	apiKey := os.Getenv("GEMINI_API_KEY")
if apiKey == "" {
    log.Fatal("API Key not found. Please set the GEMINI_API_KEY environment variable.")
}
	fmt.Println("API Key: ", apiKey)

	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		log.Fatal(err)
	}
	// defer client.Close()

	model := client.GenerativeModel("gemini-pro-vision")

	model.SafetySettings = []*genai.SafetySetting{
		{
			Category:  genai.HarmCategoryHarassment,
			Threshold: genai.HarmBlockMediumAndAbove,
		},
		{
			Category:  genai.HarmCategoryHateSpeech,
			Threshold: genai.HarmBlockMediumAndAbove,
		},
		{
			Category:  genai.HarmCategorySexuallyExplicit,
			Threshold: genai.HarmBlockMediumAndAbove,
		},

		{
			Category:  genai.HarmCategoryDangerousContent,
			Threshold: genai.HarmBlockMediumAndAbove,
		},
	}
	return model
}
