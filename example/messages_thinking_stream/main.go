package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"

	claude "github.com/potproject/claude-sdk-go"
)

func main() {
	apiKey := os.Getenv("API_KEY")
	c := claude.NewClient(apiKey)
	m := claude.RequestBodyMessages{
		Model:     "claude-3-7-sonnet-20250219",
		MaxTokens: 8192,
		Thinking:  claude.UseThinking(4096),
		Messages: []claude.RequestBodyMessagesMessages{
			{
				Role:    claude.MessagesRoleUser,
				Content: "Guess the Earth's population in 2100",
			},
		},
	}
	ctx := context.Background()
	stream, err := c.CreateMessagesStream(ctx, m)
	if err != nil {
		panic(err)
	}
	defer stream.Close()
	streamType := ""
	for {
		res, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			panic(err)
		}
		if res.Content[0].Type == claude.ResponseBodyMessagesContentTypeThinking && streamType != claude.ResponseBodyMessagesContentTypeThinking {
			fmt.Println("[thinking]")
			streamType = claude.ResponseBodyMessagesContentTypeThinking
		}
		if res.Content[0].Type == claude.ResponseBodyMessagesContentTypeText && streamType != claude.ResponseBodyMessagesContentTypeText {
			fmt.Println()
			fmt.Println("[text]")
			streamType = claude.ResponseBodyMessagesContentTypeText
		}

		fmt.Printf("%s", res.Content[0].Thinking)
		fmt.Printf("%s", res.Content[0].Text)
	}
	fmt.Println()
	// Output:
	// [thinking]
	// I need to make a prediction about the Earth's population in the year 2100. This is asking me to make a forecast about demographic trends over the next ~75 years.

	// To make a reasonable estimate, I should consider:
	// 1. Current population (as of 2023, the world population is about 8 billion people)
	// 2. Current population growth trends (the rate has been slowing)
	// 3. UN and other expert demographic projections

	// The UN Population Division regularly makes projections for future population growth. In their most recent projections, they typically provide several scenarios:
	// - Low variant (assuming lower fertility rates)
	// - Medium variant (their "most likely" scenario)
	// - High variant (assuming higher fertility rates)

	// Based on my knowledge, the UN's medium projection for 2100 is somewhere around 10-11 billion people. The range across different scenarios is quite wide, from potentially as low as 8-9 billion (if fertility rates drop significantly worldwide) to as high as 12-13 billion in high-growth scenarios.

	// Some key factors that will influence the actual outcome:
	// - Fertility rates in developing regions, especially Africa (currently with the highest growth rates)
	// - Further demographic transition in middle-income countries
	// - Aging populations in developed countries
	// - Advances in medicine and healthcare
	// - Climate change impacts
	// - Migration patterns
	// - Potential unforeseen events (pandemics, etc.)

	// I should note that population projections become less certain the further into the future they go, and 2100 is quite far away. There's significant uncertainty in any estimate.

	// My best estimate based on current projections would be around 10-11 billion people by 2100, with the understanding that this is subject to many variables and unknowns.owns.
	// [text]
	// # Earth's Population in 2100: Projection

	// Based on demographic trends and expert forecasts, Earth's population in 2100 is projected to be approximately **10-11 billion people**.

	// This estimate reflects the United Nations' medium-variant projection, which is generally considered the most likely scenario. However, there's significant uncertainty in long-term population forecasts due to several factors:

	// ## Key Variables
	// - Declining fertility rates worldwide
	// - Aging populations in developed nations
	// - Economic development in currently high-growth regions
	// - Advances in healthcare and life expectancy
	// - Climate change impacts
	// - Migration patterns
	// - Policy decisions regarding family planning

	// Different models produce a range of estimates from as low as 8.5 billion to as high as 13 billion people by 2100, depending on how these variables evolve over time.

	// Would you like me to elaborate on any particular aspect of future population trends?
}
