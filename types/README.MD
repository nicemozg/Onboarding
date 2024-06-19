## В языке программирования `Go`
типы данных можно разделить на две категории:

###  Значимые типы (value types):

+ int
+ int8, int16, int32, int64
+ uint, uint8, uint16, uint32, uint64, uintptr
+ float32, float64
+ complex64, complex128
+ bool
+ string
+ struct (составные типы данных)
+ array (массивы)
+ interface (интерфейсы)

### Ссылочные типы (reference types):

+ slice (срезы)
+ map (карты)
+ chan (каналы)
+ pointer (указатели)
+ function (функции)

_Значимые типы данных хранят свои значения непосредственно в памяти, в то время как ссылочные типы
хранят ссылки на значения в памяти. Использование значимых типов приводит к копированию значения при присваивании,
в то время как использование ссылочных типов обычно означает работу с адресами памяти и обмен данными по ссылке.

```go
package main

import "fmt"

// Определение структуры
type Person struct {
    Name string
    Age  int
}

// Функция для демонстрации различных типов
func main() {
    // Встроенные типы
    var a int = 42
    var b float64 = 3.14
    var c bool = true
    var d string = "hello"

    // Массив
    var e [3]int = [3]int{1, 2, 3}

    // Срез
    var f []int = []int{1, 2, 3, 4, 5}

    // Карта
    var g map[string]int = map[string]int{"one": 1, "two": 2}

    // Указатель
    var h *int = &a

    // Структура
    var i Person = Person{Name: "Alice", Age: 30}

    // Канал
    var j chan int = make(chan int)

    fmt.Println(a, b, c, d)
    fmt.Println(e)
    fmt.Println(f)
    fmt.Println(g)
    fmt.Println(*h)
    fmt.Println(i)
    fmt.Println(j)
}
```