# 💻GraphQL vs REST


## **1. Desenho do Experimento**

O experimento foi elaborado visando assegurar a validade dos resultados e responder às questões de pesquisa propostas. Para isso, foram definidos hipóteses, variáveis dependentes e independentes.
 

 ####  A. Hipóteses Nula e Alternativa
* Hipótese Nula (H0): Não há diferença significativa entre GraphQL e REST em termos de tempo de resposta e tamanho das respostas para as consultas.

* Hipótese Alternativa (H1): GraphQL apresenta uma diferença significativa (tempo de resposta menor ou tamanho da resposta menor) quando comparado ao REST.

#### B. Variáveis Dependentes
* Tempo de Resposta: Medido em milissegundos, representa a rapidez com que a API responde a uma consulta.

* Tamanho da Resposta: Medido em bytes, representa a quantidade de dados trafegados em uma consulta, o que impacta a eficiência da comunicação de rede.

#### B. Variáveis Independentes
* Tipo de API: Será utilizado GraphQL e REST como comparativos.

* Tipo de Consulta: Variar entre consultas de dados específicos e  múltiplas entidades.

* Volume de Dados: O número de registros ou objetos retornados nas consultas será modificado para simular diferentes cargas de dados.

#### C. Tratamentos
Para cada variável independente (tipo de API e tipo de consulta), será aplicado um tratamento que envolve a realização de consultas similares em ambas as APIs. Exemplos de consultas incluem:

* Busca por um registro específico.
* Busca por múltiplos registros.


#### D. Objetos Experimentais
Neste experimento, serão utilizados dados reais de repositórios mais populares do GitHub para que possa ser consultado tanto pela API REST quanto pela GraphQL. Estes dados serão estruturados de maneira similares em ambas as APIs para assegurar a comparação entre elas.

##### E. Tipo de Projeto Experimental
Será realizado um experimento experimental com dois grupos: consultas GraphQL e consultas REST. Este modelo permite a comparação direta entre os dois grupos de tratamento, com base nos dados obtidos a partir de consultas similares.

O experimento será estruturado como um estudo comparativo controlado com delineamento de medidas repetidas. Com isso, cada tipo de API (GraphQL e REST) será testado usando um conjunto de consultas padronizadas, que variam em complexidade e volume de dados.

O delineamento de medidas repetidas será utilizado para que as mesmas consultas sejam executadas repetidamente em ambas as APIs, reduzindo o efeito de variáveis de confusão e controlando a variabilidade dos resultados em função de fatores externos. Assim, cada consulta é executada em ambas as APIs sob condições controladas e replicada diversas vezes para garantir a precisão e confiabilidade dos resultados.

As consultas serão realizadas em três diferentes cenários experimentais:

Consulta Simples: Retorno de um único registro ou item específico.
Consulta Complexa: Retorno de múltiplos registros ou itens agregados.

Para cada cenário, as medições serão feitas nos tempos de resposta e tamanhos das respostas para ambas as APIs, de forma a verificar se existe uma diferença significativa entre GraphQL e REST em termos de eficiência. 
Esse delineamento permite uma análise robusta do desempenho, o impacto do volume de dados e da complexidade de cada consulta.


#### F. Quantidade de Medições
Para assegurar validade e confiabilidade, cada tipo de consulta será repetido pelo menos 1000 vezes para cada API (GraphQL e REST), de forma que as variâncias possam ser controladas e dados atípicos, identificados.

#### G. Ameaças à Validade

* Ameaças à Validade Interna: a variabilidade de desempenho do servidor, variabilidade de rede ou cachê de dados. Será necessário garantir que todas as consultas sejam realizadas em condições de rede e de servidor similares.

* Ameaças à Validade Externa: A aplicação dos resultados deste experimento é limitada a API específica (do GitHub) e ao ambiente onde o experimento foi realizado. Assim, os resultados podem não se generalizar para outros tipos de APIs ou configurações de servidores.

* Controle de Influências Externas: Realizar o experimento em um ambiente controlado, com cargas de rede estáveis, minimizando o impacto de variáveis externas.

## 2. Metodologia

### 2.1 Criação do Dataset

Este experimento compara o desempenho entre as APIs REST e GraphQL do GitHub. Utilizaremos consultas similares àquelas do trabalho anterior, "Análise de repositórios populares no GitHub - [Lab01]", buscando dados de repositórios populares. Os dados serão coletados tanto via API REST quanto via GraphQL e estruturados de forma similar para garantir uma comparação justa. Para mitigar ameaças à validade, realizaremos 1000 requisições para cada API, permitindo uma análise estatística mais robusta e minimizando o impacto de variações momentâneas no desempenho.

### 2.2 Consultas

As consultas a seguir serão utilizadas para recuperar informações sobre os repositórios. Ambas as consultas são projetadas para retornar informações semelhantes, como nome, data de criação, URL, número de estrelas, issues, pull requests, releases e linguagem principal. A repetição de 1000 requisições para cada consulta nos permitirá medir o tempo médio de resposta e outras métricas de desempenho.

**GraphQL:**

```graphql
query($number_of_repos_per_request: Int!, $cursor: String){
  search(query: "stars:>0", type: REPOSITORY, first: $number_of_repos_per_request, after: $cursor) {
    edges {
      node {
        ... on Repository {
          name
          createdAt
          url
          stargazers { totalCount }
          issues(states: CLOSED) { totalCount }
          pullRequests(states: [OPEN, CLOSED, MERGED]) { totalCount }
          releases { totalCount }
          primaryLanguage { name }
          closedIssues: issues(states: CLOSED) { totalCount }
          totalIssues: issues(states: [OPEN, CLOSED]) { totalCount }
          defaultBranchRef {
            name
            target {
              ... on Commit {
                committedDate
              }
            }
          }
        }
      }
    }
    pageInfo {
      hasNextPage
      endCursor
    }
  }
}

```
Variáveis da consulta GraphQL:

{
  "number_of_repos_per_request": 10,
  "cursor": null
}


A consulta GraphQL busca 10 repositórios por requisição (number_of_repos_per_request: 10) e utiliza a paginação via cursor (after: $cursor) para obter os próximos resultados nas requisições subsequentes. O cursor será atualizado a cada requisição com o valor de endCursor retornado pela consulta anterior. Este processo será repetido 100 vezes para completar as 1000 requisições, totalizando a coleta de dados de 1000 repositórios.

REST:

https://api.github.com/search/repositories?q=stars:>0&sort=stars&order=desc&per_page=10

A consulta REST utiliza parâmetros de URL para buscar 10 repositórios por página (per_page=10), ordenados por número de estrelas decrescente (sort=stars&order=desc). Similarmente à consulta GraphQL, a paginação será utilizada para obter os 1000 repositórios através de múltiplas requisições. A paginação na API REST geralmente é feita através de um parâmetro page ou de um link de paginação fornecido no cabeçalho da resposta. Este processo também será repetido até completar as 1000 requisições.

Ao executar 1000 requisições para cada API, poderemos comparar o desempenho em cenários de carga mais realistas e obter resultados mais confiáveis.

## 5. Resultados

Os resultados obtidos após a execução de 1000 requisições para cada API demonstram uma diferença significativa no desempenho entre REST e GraphQL.

A API REST apresentou uma média de tempo de resposta de 209,841 ms, enquanto a GraphQL obteve uma média de 170,775 ms.  Isso indica que, neste experimento, as consultas GraphQL foram, em média, mais rápidas que as consultas REST.

Em relação ao tamanho da resposta, a API GraphQL apresentou uma média de 95,534 ms contra 98 ms da API REST. Isso sugere que as respostas GraphQL, em termos de tamanho, foram ligeiramente menores.


| Métrica                      | REST      | GraphQL   |
|-------------------------------|-----------|-----------|
| Tempo de Resposta (ms) | 209,841  | 170,775  |
| Tamanho de Resposta (ms)  | 98 | 95,534 |

Esses resultados sugerem que a API GraphQL oferece um desempenho superior em termos de tempo de resposta para este cenário específico, com uma pequena vantagem também no tamanho da resposta.  No entanto, é importante lembrar que este é apenas um experimento e os resultados podem variar dependendo de diversos fatores, como a complexidade da consulta, a carga do servidor e a qualidade da conexão de rede.




  
