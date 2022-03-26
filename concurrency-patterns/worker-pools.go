package main

import "fmt"

func main() {
	//fibonacci
	tarefas := make(chan int, 45)
	resultados := make(chan int, 45)
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)

	for i := 0; i < 45; i++ {
		tarefas <- i
	}
	close(tarefas)

	for i := 0; i < 45; i++ {
		resultado := <-resultados
		fmt.Println(resultado)
	}

}

func worker(tarefas <-chan int, resultados chan<- int) {
	for number := range tarefas {
		resultados <- fibo(number)
	}
}

func fibo(position int) int {
	if position <= 1 {
		return position
	}

	return fibo(position-2) + fibo(position-1)
}
