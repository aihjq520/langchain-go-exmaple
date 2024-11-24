package example

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/tmc/langchaingo/chains"
	"github.com/tmc/langchaingo/documentloaders"
)

func SummaryPdf() {
	f, err := os.Open("./clean-code.pdf")
	if err != nil {
		fmt.Print("读取失败")
		log.Fatal(err)
	}

	defer f.Close()

	fi, err := f.Stat()

	fmt.Print(fi.Size())

	if err != nil {
		log.Fatal(err)
	}

	pdfDoc := documentloaders.NewPDF(f, fi.Size())

	doc, err := pdfDoc.Load(context.Background())

	fmt.Println(doc[0])
	if err != nil {
		fmt.Print("load失败")
		log.Fatal(err)
	}

	ollama := CreateOllama()

	chain := chains.LoadMapReduceSummarization(ollama)

	output, err := chain.Call(context.Background(), map[string]any{"input_documents": doc})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(output)

}
