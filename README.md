# 🏃‍♂️ Processamento Paralelo de Dados em Go com Goroutines e Sincronização

Este script Go demonstra como utilizar goroutines, wait groups e mutexes para processar grandes volumes de dados de forma eficiente e segura.

## 🎯 Objetivo

O objetivo principal é dividir o processamento de um conjunto de dados em tarefas menores, distribuindo-as entre múltiplas goroutines para execução simultânea. Isso pode acelerar significativamente o tempo de processamento, especialmente em máquinas com múltiplos núcleos.

## ⚙️ Funcionamento

1. **Divisão dos Dados**: O conjunto de dados é dividido em partes menores (chunks), cada uma destinada a ser processada por uma goroutine.
2. **Criação de Goroutines**: Uma goroutine é criada para cada chunk de dados. Cada goroutine executa a função `processarDados`, que simula uma operação demorada em cada elemento do seu chunk.
3. **WaitGroup**: Um `sync.WaitGroup` é utilizado para coordenar a execução das goroutines. Ele garante que o programa principal aguarde a conclusão de todas as goroutines antes de finalizar.
4. **Mutex**: Um `sync.Mutex` é usado para proteger a seção crítica do código, onde os resultados são impressos. Isso evita que múltiplas goroutines tentem acessar e modificar a saída ao mesmo tempo, o que poderia levar a resultados incorretos.

## 🚀 Como usar

1. Clone o repositório:
    ```sh
    git clone git@github.com:br4tech/go-with-goroutines.git
    ```
2. Navegue até o diretório do projeto:
    ```sh
    cd go-with-goroutines
    ```
3. Instale as dependências:
    ```sh
    go mod download
    ```
4. Execute o script:
    ```sh
    go run main.go
    ```

## 🛠️ Requisitos

- Go 1.21 ou superior
- Git

## 📂 Estrutura do Projeto

```plaintext
go-with-goroutines/
├── main.go
├── go.mod
└── README.md
