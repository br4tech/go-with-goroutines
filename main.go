package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func processarDados(dados []int, wg *sync.WaitGroup, mutex *sync.Mutex) {
	defer wg.Done()

	for _, valor := range dados {
		// Simulação de uma operação demorada
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)

		mutex.Lock() // Bloqueia o acesso à seção crítica
		fmt.Printf("Processado: %d\n", valor)
		mutex.Unlock() // Libera o acesso à seção crítica
	}
}

func main() {
	dados := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	numGoroutines := 4

	var wg sync.WaitGroup
	wg.Add(numGoroutines)

	var mutex sync.Mutex

	chunkSize := len(dados) / numGoroutines
	for i := 0; i < numGoroutines; i++ {
		inicio := i * chunkSize
		fim := (i + 1) * chunkSize
		if i == numGoroutines-1 {
			fim = len(dados) // Garante que o último chunk pegue os elementos restantes
		}
		go processarDados(dados[inicio:fim], &wg, &mutex)
	}

	wg.Wait() // Aguarda todas as goroutines terminarem
	fmt.Println("Processamento concluído!")
}
