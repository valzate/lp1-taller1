package main

import (
	"fmt"
	"sync"
	"time"
)

// Objetivo: Implementar un pool de workers que procesa trabajos y retorna resultados.
// Usa un canal de trabajos y otro de resultados. Cierra canales correctamente.
// TODO: completa las funciones y la orquestación en main().

type trabajo struct {
	ID int
	X  int
}

type resultado struct {
	ID        int
	X         int
	Procesado int
}

func worker(id int, jobs <-chan trabajo, results chan<- resultado, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := range jobs {
		// TODO: procesar j (simular trabajo con Sleep)

		r := resultado{
			ID:        j.ID,
			X:         j.X,
			Procesado: j.X * 2, // ejemplo de procesamiento
		}
		fmt.Printf("[worker %d] procesa trabajo %d -> %d\n", id, j.ID, r.Procesado)
		results <- r
	}
	time.Sleep(100 * time.Millisecond) // simula tiempo de procesamiento
	fmt.Printf("[worker %d] no hay más trabajos\n", id)
}

func main() {
	nTrabajos := 12
	nWorkers := 3

	jobs := make(chan trabajo)
	results := make(chan resultado)

	var wg sync.WaitGroup

	// TODO: lanzar nWorkers workers
	wg.Add(nWorkers)
	for i := 1; i <= nWorkers; i++ {
		go worker(i, jobs, results, &wg)

	}

	// TODO: productor de trabajos
	go func() {
		for i := 1; i <= nTrabajos; i++ {
			j := trabajo{ID: i, X: i * 10}
			fmt.Printf("[main] produce trabajo %d\n", j.ID)
			jobs <- j
			time.Sleep(50 * time.Millisecond) // simula tiempo entre producciones

		}
		close(jobs) // importante: cerrar para que los workers terminen
	}()

	// TODO: recolectar resultados en otra goroutine y cerrar results al final
	var wgCollect sync.WaitGroup
	wgCollect.Add(1)
	go func() {
		defer wgCollect.Done()
		wg.Wait()      // esperar que todos los workers terminen
		close(results) // entonces cerrar resultados
	}()

	// Consumidor principal de resultados
	for r := range results {
		fmt.Printf("[main] resultado: trabajo %d -> %d\n", r.ID, r.Procesado)
	}

	wgCollect.Wait()
	fmt.Println("Pool finalizado.")
}
