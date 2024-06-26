### Инициализация каналов и WaitGroup:

- Создаются два канала:
    - `in` для передачи данных от основной горутины к воркерам.
    - `out` для получения результата от воркеров.
- Создается WaitGroup (`wg`), чтобы дождаться завершения всех воркеров.

### Создание и запуск воркеров:

- В цикле `for i := 0; i < numWorkers; i++` создаются и запускаются воркеры с помощью функции `worker(in, out, &wg)`.
- Каждый воркер начинает ждать данных из канала `in` (`<-chan int`) и отправляет результаты в канал `out` (`chan<- int`).

### Запись данных в канал `in`:

- В отдельной горутине с помощью анонимной функции `go func() {...}` заполняется канал `in` значениями от 1 до 5 с помощью цикла `for i := 1; i <= 5; i++`.
- После этого канал `in` закрывается (`close(in)`), чтобы воркеры могли завершить свою работу после обработки всех данных.

### Ожидание завершения воркеров:

- Еще одна анонимная горутина запущена для ожидания завершения всех воркеров с помощью `wg.Wait()`.
- После завершения всех воркеров канал `out` также закрывается (`close(out)`).

### Чтение результатов из канала `out` и вывод:

- В основной горутине происходит чтение результатов из канала `out` в цикле `for result := range out`.
- Как только данные поступают в канал `out`, они выводятся на экран с помощью `fmt.Println("Result:", result)`.

Таким образом, данные параллельно отправляются из основной горутины в воркеры через канал `in`, каждый воркер обрабатывает данные (возводит в квадрат) и отправляет результаты обратно через канал `out`. В конце основная горутина читает и выводит результаты из канала `out`.
