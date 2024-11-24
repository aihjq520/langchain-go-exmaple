package example

import (
	"context"
	"fmt"
	"log"

	"github.com/tmc/langchaingo/agents"
	"github.com/tmc/langchaingo/chains"
	"github.com/tmc/langchaingo/tools"
	"github.com/tmc/langchaingo/tools/serpapi"
)

func ApiTool() {
	ollma := CreateOllama()

	search, err := serpapi.New()

	if err != nil {
		fmt.Println("serpapi error")
		log.Fatal(err)
	}

	agentTools := []tools.Tool{
		search,
		tools.Calculator{},
	}

	agent := agents.NewOneShotAgent(ollma, agentTools, agents.WithMaxIterations(1))

	executor := agents.NewExecutor(agent)

	// serpapi问题貌似只支持英文 can't obtain real-time information through a simple google api
	question := "What's the date today?  and What's the stock price of apple ？"

	resp, err := chains.Run(context.Background(), executor, question)

	if err != nil {
		fmt.Println("run error")
		log.Fatal(err)
	}

	fmt.Println(resp)
}
