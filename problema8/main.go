package main

import (
	"fmt"
	"time"
)

// Objetivo: Simular "futuros" en Go usando canales. Una función lanza trabajo asíncrono
// y retorna un canal de solo lectura con el resultado futuro.
// TODO: completa las funciones y experimenta con varios futuros a la vez.

func asyncCuadrado(x int) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		// TODO: simular trabajo
		time.Sleep(time.Millisecond * 100)
		ch <- x * x
	}()
	return ch
}

// Pista: crea una función fanIn que recibe múltiples <-chan int y retorna un único <-chan int
// que emita todos los valores. Requiere goroutines y cerrar el canal de salida cuando todas terminen
func fanIn(channels ...<-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, ch := range channels {
			for val := range ch {
				out <- val
			}
		}
	}()
	return out
}

func main() {
	// TODO: crea varios futuros y recolecta sus resultados: f1, f2, f3
	f1 := asyncCuadrado(8)
	f2 := asyncCuadrado(14)
	f3 := asyncCuadrado(42)

	// TODO: Opción 1: esperar cada futuro secuencialmente

	// TODO: Opción 2: fan-in (combinar múltiples canales)
	merged := fanIn(f1, f2, f3)
	for result := range merged {
		fmt.Println(result)
	}
}
