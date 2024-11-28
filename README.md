# controlled-experiment-graphql-vs-rest

#  Um experimento controlado GraphQL vs REST- [Lab05]

💻Laboratório de Experimentação de Software 

O objetivo deste laboratório é realizar um experimento controlado para avaliar quantitativamente os benefícios da adoção de uma API GraphQL. Especificamente, as seguintes perguntas devem ser respondidas: 

## **Questões de Pesquisa:**

**RQ 01**. Respostas às consultas GraphQL são mais rápidas que respostas às consultas REST?

**RQ 02.** Respostas às consultas GraphQL tem tamanho menor que respostas às consultas REST? 

## **Desenho do Experimento**
 
Para realizar o desenho do experimento, será definido cada um dos tópicos a seguir: 

* A. Hipóteses Nula e Alternativa
* B. Variáveis Dependentes
* B. Variáveis Independentes
* C. Tratamentos
* D. Objetos Experimentais
* E. Tipo de Projeto Experimental
* F. Quantidade de Medições
* G. Ameaças à Validade

## 👩🏻‍💻 Alunos:
* Bárbara Mattioly Andrade  
* Laura Enísia Rodrigues Melo
* Samuel Marques Sousa Leal 
 
## 👨‍🏫 Professor:
* João Paulo Carneiro Aramuni

## Instruções de Execução

### Pré-requisitos

1. **Instalar Go**: [Baixe e instale o Go](https://go.dev/dl/).
2. **Configurar o Token do GitHub**:
   - Crie um arquivo `.env` na raiz do projeto com o seguinte conteúdo:
     ```
     PERSONAL_ACCESS_TOKEN={token}
     ```
   - Substitua `{token}` pelo seu token de acesso do GitHub.

### Executando o Projeto

1. Caso não tenha o Go instalado em sua máquina, instale através do site oficial (https://golang.org/dl/), seguindo as instruçõesde instalação.
2. Verifique que o Go está devivamente instalado em seu sistema, executando o seguinte comando em seu terminal
      ```bash
      go version
3. Abra o terminal e navegue até o diretório do projeto `cd src/`.
4. Execute o comando:

   ```bash
   go run fetch_github_api_data.go
