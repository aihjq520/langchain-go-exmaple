package example

import (
	"github.com/tmc/langchaingo/llms/ollama"
	"github.com/tmc/langchaingo/prompts"
)

func CreateOllama() *ollama.LLM {
	llm, err := ollama.New(ollama.WithModel("qwen2.5:1.5b"))

	if err != nil {
		panic(err)
	}

	return llm
}

func CreatePrompt() prompts.ChatPromptTemplate {
	return prompts.NewChatPromptTemplate([]prompts.MessageFormatter{
		prompts.NewSystemMessagePromptTemplate("", nil),
		prompts.NewHumanMessagePromptTemplate("{{.text}}", []string{
			"text",
		}),
	})
}
