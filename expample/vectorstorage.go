package example

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/tmc/langchaingo/chains"
	"github.com/tmc/langchaingo/embeddings"
	"github.com/tmc/langchaingo/vectorstores/pinecone"
)

func VectorStorage() {
	// client, err := qdrant.NewClient(&qdrant.Config{
	// 	Host:   "d8af6006-f46d-4fe7-83b7-40c4b67bda4d.us-west-1-0.aws.cloud.qdrant.io",
	// 	Port:   6334,
	// 	APIKey: os.Getenv("QDRANT_API_KEY"),
	// 	UseTLS: true,
	// })

	ctx := context.Background()

	ollama := CreateOllama()

	e, err := embeddings.NewEmbedder(ollama)

	if err != nil {
		log.Fatal(err)
	}

	if err != nil {
		log.Fatal(err)
	}

	store, err := pinecone.New(
		pinecone.WithHost("https://langchain-go-zm3n7xr.svc.aped-4627-b74a.pinecone.io"),
		pinecone.WithEmbedder(e),
		pinecone.WithAPIKey(os.Getenv("VECTORDATABASE_API_KEY")),
		pinecone.WithNameSpace("langchain-go"),
	)

	if err != nil {
		log.Fatal(err)
	}

	ans, err := store.SimilaritySearch(ctx, "如何提高开发效率", 3)

	chain := chains.LoadMapReduceQA(ollama)

	resp, err := chain.Call(ctx, map[string]any{"input_documents": ans, "question": "如何提高开发效率"})

	fmt.Print(resp)

	// dat, err := os.Open("./1.txt")

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// 加载文件
	// doc, err := documentloaders.NewText(dat).Load(context.Background())

	// 分割
	// spliter := textsplitter.NewRecursiveCharacter(textsplitter.WithChunkSize(500), textsplitter.WithChunkOverlap(0))
	// docment, err := textsplitter.SplitDocuments(spliter, doc)

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// _, err = store.AddDocuments(ctx, docment)

	// if err != nil {
	// 	log.Fatal(err)
	// }
}

//https://github.com/tmc/langchaingo/blob/main/examples/pinecone-vectorstore-example/pinecone_vectorstore_example.go
