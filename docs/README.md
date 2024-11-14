# 💻GraphQL vs REST


## **Desenho do Experimento**
 
Para realizar o desenho do experimento, será definido cada um dos tópicos a seguir: 

 ##  A. Hipóteses Nula e Alternativa
* Hipótese Nula (H0): Não há diferença significativa entre GraphQL e REST em termos de tempo de resposta e tamanho das respostas para as consultas.

* Hipótese Alternativa (H1): GraphQL apresenta uma diferença significativa (tempo de resposta menor ou tamanho da resposta menor) quando comparado ao REST.

## B. Variáveis Dependentes
* Tempo de Resposta: Medido em milissegundos, representa a rapidez com que a API responde a uma consulta.

* Tamanho da Resposta: Medido em bytes, representa a quantidade de dados trafegados em uma consulta, o que impacta a eficiência da comunicação de rede.

## B. Variáveis Independentes
* Tipo de API: Será utilizado GraphQL e REST como comparativos.

* Tipo de Consulta: Variar entre consultas de dados específicos e  múltiplas entidades.

* Volume de Dados: O número de registros ou objetos retornados nas consultas será modificado para simular diferentes cargas de dados.

## C. Tratamentos
Para cada variável independente (tipo de API e tipo de consulta), será aplicado um tratamento que envolve a realização de consultas similares em ambas as APIs. Exemplos de consultas incluem:

* Busca por um registro específico.
* Busca por múltiplos registros.


## D. Objetos Experimentais
Neste experimento, serão utilizados dados reais de repositórios mais populares do GitHub para que possa ser consultado tanto pela API REST quanto pela GraphQL. Estes dados serão estruturados de maneira similares em ambas as APIs para assegurar a comparação entre elas.

## E. Tipo de Projeto Experimental
Será realizado um experimento experimental com dois grupos: consultas GraphQL e consultas REST. Este modelo permite a comparação direta entre os dois grupos de tratamento, com base nos dados obtidos a partir de consultas similares.

O experimento será estruturado como um estudo comparativo controlado com delineamento de medidas repetidas. Com isso, cada tipo de API (GraphQL e REST) será testado usando um conjunto de consultas padronizadas, que variam em complexidade e volume de dados.

O delineamento de medidas repetidas será utilizado para que as mesmas consultas sejam executadas repetidamente em ambas as APIs, reduzindo o efeito de variáveis de confusão e controlando a variabilidade dos resultados em função de fatores externos. Assim, cada consulta é executada em ambas as APIs sob condições controladas e replicada diversas vezes para garantir a precisão e confiabilidade dos resultados.

As consultas serão realizadas em três diferentes cenários experimentais:

Consulta Simples: Retorno de um único registro ou item específico.
Consulta Complexa: Retorno de múltiplos registros ou itens agregados.

Para cada cenário, as medições serão feitas nos tempos de resposta e tamanhos das respostas para ambas as APIs, de forma a verificar se existe uma diferença significativa entre GraphQL e REST em termos de eficiência. 
Esse delineamento permite uma análise robusta do desempenho, o impacto do volume de dados e da complexidade de cada consulta.


## F. Quantidade de Medições
Para assegurar validade e confiabilidade, cada tipo de consulta será repetido pelo menos 10 vezes para cada API (GraphQL e REST), de forma que as variâncias possam ser controladas e dados atípicos, identificados.

## G. Ameaças à Validade

* Ameaças à Validade Interna: a variabilidade de desempenho do servidor, variabilidade de rede ou cachê de dados. Será necessário garantir que todas as consultas sejam realizadas em condições de rede e de servidor similares.

* Ameaças à Validade Externa: A aplicação dos resultados deste experimento é limitada a API específica (do GitHub) e ao ambiente onde o experimento foi realizado. Assim, os resultados podem não se generalizar para outros tipos de APIs ou configurações de servidores.

* Controle de Influências Externas: Realizar o experimento em um ambiente controlado, com servidor dedicado e cargas de rede estáveis, minimizando o impacto de variáveis externas.