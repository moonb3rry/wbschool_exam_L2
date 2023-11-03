Что выведет программа? Объяснить вывод программы.

```go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func asChan(vs ...int) <-chan int {
	c := make(chan int)

	go func() {
		for _, v := range vs {
			c <- v
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}

		close(c)
	}()
	return c
}

func merge(a, b <-chan int) <-chan int {
	c := make(chan int)
	go func() {
		for {
			select {
			case v := <-a:
				c <- v
			case v := <-b:
				c <- v
			}
		}
	}()
	return c
}

func main() {

	a := asChan(1, 3, 5, 7)
	b := asChan(2, 4 ,6, 8)
	c := merge(a, b )
	for v := range c {
		fmt.Println(v)
	}
}
```

Ответ:
```
Вывод будет неопределенный: числа 1, 2, 3, 4, 5, 6, 7, 8 выведутся в случайном
порядке, а затем будет бесконечное количество 0 из-за дедлока. Почему:
1. В функции asChan задержка между добавлением в канал случайна, поэтому в каналы
a и b числа могут добавиться в другом порядке.
2. В функции merge цикл for бесконечный, и select будет читать дефолтные значения
из обоих каналов (для int -- 0) после того как закончатся реальные. Чтобы это
обойти, можно добавить в блок select проверки не только на значения, но и на то,
был ли канал закрыт.

```
