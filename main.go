package main

import (
	"github.com/joho/godotenv"
	example "langchainexample.com/m/expample"
)

func main() {
	godotenv.Load()
	example.ApiTool()
	// VectorStorage()
}
