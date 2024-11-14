package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

const (
	restEndpoint = "https://api.github.com/repos/octocat/hello-world"
	graphqlURL   = "https://api.github.com/graphql"
	queryGraphQL = `{"query": "query { repository(owner: \"octocat\", name: \"hello-world\") { name owner { login } description } }"}`
	repetitions  = 10
)

type Measurement struct {
	APIType      string
	ResponseTime int64
	ResponseSize int
}

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Erro ao carregar o arquivo .env:", err)
		os.Exit(1)
	}

	token := os.Getenv("GITHUB_TOKEN")

	measurements := []Measurement{}

	// Faz as medições para a API REST
	for i := 0; i < repetitions; i++ {
		measurement := measureREST(token)
		measurements = append(measurements, measurement)
	}

	// Faz as medições para a API GraphQL
	for i := 0; i < repetitions; i++ {
		measurement := measureGraphQL(token)
		measurements = append(measurements, measurement)
	}

	fmt.Println("Resultados do Experimento:")
	for _, measurement := range measurements {
		fmt.Printf("API: %s | Tempo de Resposta: %dms | Tamanho da Resposta: %d bytes\n",
			measurement.APIType, measurement.ResponseTime, measurement.ResponseSize)
	}
}

// Função responsável por fazer consultas à API REST e medir tempo e tamanho da resposta
func measureREST(token string) Measurement {
	start := time.Now()

	req, _ := http.NewRequest("GET", restEndpoint, nil)
	req.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Erro na consulta REST:", err)
		return Measurement{}
	}
	defer resp.Body.Close()

	duration := time.Since(start).Milliseconds()
	body, _ := io.ReadAll(resp.Body)

	return Measurement{
		APIType:      "REST",
		ResponseTime: duration,
		ResponseSize: len(body),
	}
}

// Função responsável por fazer consultas à API GraphQL e medir tempo e tamanho da resposta
func measureGraphQL(token string) Measurement {
	start := time.Now()

	req, _ := http.NewRequest("POST", graphqlURL, bytes.NewBuffer([]byte(queryGraphQL)))
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Erro na consulta GraphQL:", err)
		return Measurement{}
	}
	defer resp.Body.Close()

	duration := time.Since(start).Milliseconds()
	body, _ := io.ReadAll(resp.Body)

	return Measurement{
		APIType:      "GraphQL",
		ResponseTime: duration,
		ResponseSize: len(body),
	}
}
