package main

import (
	"fmt"
	"sync"
)

// worker выполняет задачи, полученные из канала jobs, и отправляет результаты в канал results
func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()       // Уменьшает счетчик WaitGroup после завершения работы
	for j := range jobs { // Чтение задач из канала jobs до его закрытия
		fmt.Printf("worker %d started job %d\n", id, j)
		results <- j * 2 // Выполнение задачи: умножение числа на 2
	}
}

func main() {
	const numJobs = 5                  // Определение количества задач
	jobs := make(chan int, numJobs)    // Создание канала для задач с буфером размером numJobs
	results := make(chan int, numJobs) // Создание канала для результатов с буфером размером numJobs
	var wg sync.WaitGroup              // Создание WaitGroup для синхронизации горутин

	// Запуск 3 рабочих горутин
	for w := 1; w <= 3; w++ {
		wg.Add(1)                        // Увеличение счетчика WaitGroup
		go worker(w, jobs, results, &wg) // Запуск горутины worker
	}

	// Отправка задач в канал jobs
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs) // Закрытие канала jobs, чтобы сообщить горутинам о завершении отправки задач

	wg.Wait()      // Ожидание завершения всех рабочих горутин
	close(results) // Закрытие канала results

	// Чтение и вывод результатов из канала results
	for result := range results {
		fmt.Println("Result:", result)
	}
}
