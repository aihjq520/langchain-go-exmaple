package example

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/tmc/langchaingo/chains"
	"github.com/tmc/langchaingo/documentloaders"
	"github.com/tmc/langchaingo/llms"
)

func Summary() {
	ollama := CreateOllama()

	dat, err := os.Open("./1.txt")

	if err != nil {
		log.Fatal(err)
	}

	// 加载文件
	doc, err := documentloaders.NewText(dat).Load(context.Background())

	// spliter := textsplitter.NewRecursiveCharacter(textsplitter.WithChunkSize(100), textsplitter.WithChunkOverlap(0))
	// docment, err := spliter.SplitText(doc[0].PageContent)

	if err != nil {
		log.Fatal(err)
	}

	// 分割后得到的向量个数
	// fmt.Println(len(docment))

	chain := chains.LoadRefineSummarization(ollama)

	output, err := chain.Call(context.Background(), map[string]any{"input_documents": doc})

	if err != nil {
		log.Fatal(err)
	}
	data := map[string]any{
		"text": output,
	}

	prompt := CreatePrompt()

	msg, _ := prompt.FormatMessages(data)

	content := []llms.MessageContent{
		llms.TextParts(llms.ChatMessageTypeSystem, "你是个翻译助手，翻译英文为中文"),  // role
		llms.TextParts(llms.ChatMessageTypeHuman, msg[1].GetContent()), // question
	}

	ollama.GenerateContent(context.Background(), content, llms.WithStreamingFunc(func(ctx context.Context, chunk []byte) error {
		fmt.Print(string(chunk))
		return nil
	}))
}
