Код демонстрирует использование каналов и горутин в языке программирования `Go` для передачи и обработки сообщений асинхронно.

### Создание и инициализация каналов:

Создаются два канала `c1` и `c2` типа `chan string` для передачи строковых сообщений.
Запуск горутин:

В первой анонимной функции, запущенной как горутина (`go func()`), происходит задержка на 1 секунду 
(`time.Sleep(1 * time.Second)`) и после этого в канал `c1` отправляется строка "one" (`c1 <- "one"`).

Во второй анонимной функции, также запущенной как горутина, происходит задержка на 2 секунды (`time.Sleep(2 * time.Second)`) и после этого в канал c2 отправляется строка "two" (`c2 <- "two"`).
Использование select для чтения из каналов:

В основной функции `main` используется цикл `for` для ожидания получения сообщений из каналов. 
Цикл выполняется два раза (`for i := 0; i < 2; i++`), так как мы ожидаем получить два сообщения.
`select` для выбора первого доступного сообщения:

В блоке select проверяются оба канала `c1` и `c2`.

* Если сообщение получено из канала c1 (`msg1 := <-c1`), выводится сообщение "`Received one`".
* Если сообщение получено из канала c2 (`msg2 := <-c2`), выводится сообщение "`Received two`".

Результат работы:

После запуска программы через некоторое время (1 и 2 секунды соответственно) 
будут получены и выведены сообщения из каналов `c1` и `c2`, демонстрируя параллельное выполнение 
горутин и асинхронную передачу данных через каналы.
Этот пример иллюстрирует базовое использование горутин и 
каналов в `Go` для создания конкурентных программ, где выполнение задач может
происходить параллельно и без блокировки основной горутины.