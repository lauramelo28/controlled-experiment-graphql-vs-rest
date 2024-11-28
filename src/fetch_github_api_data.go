package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

const (
	restEndpoint  = "https://api.github.com/search/repositories?q=stars:>0&sort=stars&order=desc&per_page=10"
	graphqlURL    = "https://api.github.com/graphql"
	queryGraphQL  = `{"query": "query($number_of_repos_per_request: Int!, $cursor: String) { search(query: \"stars:>0\", type: REPOSITORY, first: $number_of_repos_per_request, after: $cursor) { edges { node { ... on Repository { name createdAt url stargazers { totalCount } issues(states: CLOSED) { totalCount } pullRequests(states: [OPEN, CLOSED, MERGED]) { totalCount } releases { totalCount } primaryLanguage { name } closedIssues: issues(states: [CLOSED]) { totalCount } totalIssues: issues(states: [OPEN, CLOSED]) { totalCount } defaultBranchRef { name target { ... on Commit { committedDate } } } } } } pageInfo { hasNextPage endCursor } } }", "variables": {"number_of_repos_per_request": 10, "cursor": null}}`
	repetitions   = 5000
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

	// Salvar os resultados no arquivo CSV
	saveToCSV("resultados.csv", measurements)
	fmt.Println("Resultados salvos em 'resultados.csv'")
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

// Função para salvar os resultados em um arquivo CSV
func saveToCSV(filename string, measurements []Measurement) {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Erro ao criar o arquivo CSV:", err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.Write([]string{"API", "Tempo de Resposta (ms)", "Tamanho da Resposta (bytes)"})

	for _, m := range measurements {
		row := []string{
			m.APIType,
			fmt.Sprintf("%d", m.ResponseTime),
			fmt.Sprintf("%d", m.ResponseSize),
		}
		writer.Write(row)
	}
}
