## Примитивы пакета sync в языке Go

### Mutex

- **Описание**: `sync.Mutex` представляет собой примитив синхронизации, который предотвращает одновременный доступ к общему ресурсу из нескольких горутин. Мьютексы поддерживают блокировки, которые гарантируют, что только одна горутина имеет доступ к критической секции в определенный момент времени.
- **Зачем нужен**: Используется для защиты общих данных от гонок данных и несинхронизированного доступа.

### RWMutex

- **Описание**: `sync.RWMutex` также представляет собой мьютекс, но с поддержкой режима чтения-записи. Он позволяет нескольким горутинам читать данные одновременно, но блокирует доступ для записи во время чтения.
- **Зачем нужен**: Полезен в случаях, когда данные часто читаются, но редко изменяются, для увеличения параллелизма операций чтения.

### WaitGroup

- **Описание**: `sync.WaitGroup` предоставляет механизм ожидания завершения выполнения горутин. Она позволяет основной горутине дождаться завершения выполнения всех запущенных горутин.
- **Зачем нужен**: Используется для синхронизации завершения нескольких горутин перед продолжением выполнения основной программы.

### Cond

- **Описание**: `sync.Cond` представляет условную переменную, которая сигнализирует одной или нескольким горутинам о событии или изменении состояния.
- **Зачем нужен**: Полезен в сценариях, где горутины должны ожидать определенного условия перед продолжением работы.

### Once

- **Описание**: `sync.Once` обеспечивает выполнение функции ровно один раз, независимо от количества горутин, которые пытаются выполнить ее.
- **Зачем нужен**: Используется для инициализации ресурсов или глобальных переменных, которые должны быть установлены только один раз во время выполнения программы.

### Pool

- **Описание**: `sync.Pool` представляет пул объектов, который позволяет эффективно управлять временно неиспользуемыми объектами, чтобы избежать частых выделений и освобождений памяти.
- **Зачем нужен**: Используется для улучшения производительности путем повторного использования объектов вместо создания новых.

### Map

- **Описание**: `sync.Map` представляет собой потокобезопасную реализацию хеш-таблицы для пар ключ-значение.
- **Зачем нужен**: Подходит для ситуаций, когда необходим доступ к общей структуре данных из нескольких горутин без необходимости вручную синхронизировать доступ к отдельным элементам.

### Once

- **Описание**: `sync.Once` обеспечивает выполнение функции ровно один раз, независимо от количества горутин, которые пытаются выполнить ее.
- **Зачем нужен**: Используется для инициализации ресурсов или глобальных переменных, которые должны быть установлены только один раз во время выполнения программы.

### Описание

Каждый из этих примитивов предназначен для решения конкретных задач синхронизации в параллельном программировании на языке Go, обеспечивая безопасный и эффективный доступ к общим ресурсам и данных из нескольких горутин.