#### Инициализация переменных:

```go
counter int
rwMutex sync.RWMutex
```

`counter int`: Это переменная, которую будут изменять обе горутины.

`rwMutex sync.RWMutex`: Это мьютекс для синхронизации доступа к переменной counter. Он поддерживает два режима блокировки: чтение (`RLock и RUnlock`) и запись (`Lock и Unlock`).

Основная функция main:

Когда один из методов захватывает mutex то другой метод блокируется и ждет освобождения.

К примеру когда срабатыват `Lock` то метод `RLock` не сможет захватить mutex пока `Unlock` его не разблокирует. И точно так же работает в обартную сторону. Если `RLock` захватил mutex то `Lock` будет ждать пока `RUnlock` не разблокирует mutex. 