### Объяснение примера

#### Использование sync.Cond и sync.Mutex в Go

В данном примере демонстрируется использование условной переменной (`sync.Cond`) вместе с мьютексом (`sync.Mutex`) для организации ожидания сигнала от одной горутины другой.

#### Основные компоненты

1. **Инициализация мьютекса и условной переменной:**
   - Создается мьютекс `mu` с помощью `sync.Mutex`, который будет использоваться для защиты доступа к общему состоянию.
   - Создается условная переменная `cond` с помощью `sync.NewCond(&mu)`, которая будет использоваться для ожидания и уведомления горутин.

2. **Горутина, ожидающая сигнала:**
   - Запускается анонимная функция в новой горутине (`go func()`), которая блокирует мьютекс (`mu.Lock()`).
   - В цикле проверяется переменная `ready`, пока она не станет `true`.
   - Для ожидания изменения состояния используется `cond.Wait()`, который автоматически освобождает мьютекс и блокирует текущую горутину до получения сигнала от другой горутины.
   - После получения сигнала выводится сообщение `"Goroutine notified"`, и мьютекс разблокируется (`mu.Unlock()`).

3. **Основная горутина:**
   - Блокирует мьютекс `mu.Lock()` перед изменением общего состояния.
   - Устанавливает переменную `ready` в `true`, что означает, что готово к выполнению действий, которые должны быть синхронизированы.
   - Вызывает `cond.Signal()`, чтобы отправить сигнал одной из ожидающих горутин, что общее состояние изменилось.
   - Разблокирует мьютекс `mu.Unlock()`.

#### Как это работает

- При запуске программы создается две горутины: одна ожидает сигнала изменения состояния, а другая изменяет состояние и отправляет сигнал.
- Использование мьютекса `mu` и условной переменной `cond` позволяет безопасно ожидать изменения состояния и уведомлять о нем другие горутины.
- `sync.Cond` автоматически управляет мьютексом внутри себя, что делает использование условных переменных удобным и безопасным для конкурентной среды Go.

#### Заключение

Этот пример демонстрирует применение условных переменных и мьютексов в Go для синхронизации работы между горутинами. Эти механизмы являются важными инструментами для разработки безопасных и эффективных многопоточных приложений на языке Go.

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var mu sync.Mutex
	cond := sync.NewCond(&mu)

	go func() {
		time.Sleep(2 * time.Second)
		mu.Lock()
		defer mu.Unlock()
		cond.Signal() // сигнализируем об изменении состояния
	}()

	mu.Lock()
	defer mu.Unlock()
	cond.Wait() // ждем сигнала об изменении состояния
	fmt.Println("Received signal")
}
```

```go
package main

import (
   "fmt"
   "sync"
   "time"
)

func main() {
   var mu sync.Mutex
   cond := sync.NewCond(&mu)

   // Первая горутина
   go func() {
      mu.Lock()
      fmt.Println("Goroutine 1: Waiting for signal")
      cond.Wait() // Ждем сигнала
      fmt.Println("Goroutine 1: Received signal")
      mu.Unlock()
   }()

   // Вторая горутина
   go func() {
      mu.Lock()
      fmt.Println("Goroutine 2: Waiting for signal")
      cond.Wait() // Ждем сигнала
      fmt.Println("Goroutine 2: Received signal")
      mu.Unlock()
   }()

   // Отправляем сигнал после некоторой задержки
   time.Sleep(2 * time.Second)
   mu.Lock()
   fmt.Println("Sending broadcast signal to wake up all goroutines")
   cond.Broadcast() // Отправляем широковещательный сигнал
   mu.Unlock()

   // Ждем завершения работы горутин
   time.Sleep(1 * time.Second)
   fmt.Println("Main goroutine: Done")
}
```

Если одна горутина была разбужена с помощью `Signal()`, 
остальные горутины, которые ожидают на `sync.Cond`, останутся заблокированными 
до тех пор, пока не будет отправлен новый сигнал снова через `Signal()` или `Broadcast()`.

Концепция `sync.Cond` в Go предполагает, что каждый вызов `Signal()` разбудит ровно одну 
горутину из ожидающих. Если вам нужно разбудить несколько или все ожидающие горутины, 
используйте `Broadcast()`, который разбудит все ожидающие горутины одновременно.
