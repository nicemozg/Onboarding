package main

import (
	"fmt"
	"time"
)

func main() {
	// Создаем два канала для передачи строк
	c1 := make(chan string)
	c2 := make(chan string)

	// Запускаем первую горутину
	go func() {
		time.Sleep(1 * time.Second) // Ждем 1 секунду
		c1 <- "one"                 // Отправляем сообщение в канал c1
	}()

	// Запускаем вторую горутину
	go func() {
		time.Sleep(2 * time.Second) // Ждем 2 секунды
		c2 <- "two"                 // Отправляем сообщение в канал c2
	}()

	// Используем цикл для ожидания получения сообщений из каналов
	for i := 0; i < 2; i++ {
		select {
		// Если получено сообщение из канала c1, выводим его
		case msg1 := <-c1:
			fmt.Println("Received", msg1)
		// Если получено сообщение из канала c2, выводим его
		case msg2 := <-c2:
			fmt.Println("Received", msg2)
		}
	}
}
