Что выведет программа? Объяснить вывод программы. Объяснить как работают defer’ы и их порядок вызовов.

```go
package main

import (
	"fmt"
)


func test() (x int) {
	defer func() {
		x++
	}()
	x = 1
	return
}


func anotherTest() int {
	var x int
	defer func() {
		x++
	}()
	x = 1
	return x
}


func main() {
	fmt.Println(test())
	fmt.Println(anotherTest())
}
```

Ответ:
```
2
1

Отложенные функции выполняются после завершения текущей функции, в которой 
они были объявлены, но перед возвратом значения из этой функции. Если в 
функции используется несколько defer-ов, они будут выполняться в обратном 
порядке, то есть последний объявленный defer будет выполнен первым, и т.д.

Порядок выполнения test():
1. Переменная x объявляется в регистре, который используется для возврата 
 значений из функций.
2. Затем, с помощью defer, объявляется анонимная функция, которая 
 увеличивает значение x на 1. Эта анонимная функция не выполняется немедленно.
3. Далее, значение x устанавливается равным 1.
4. После этого, оператор return выполняется. Значение x (которое равно 1)
 извлекается из регистра и возвращается из функции.
5. После возврата значения из функции, отложенная функция с defer 
 выполняется, и она увеличивает значение x до 2.
Итак, значение x сохраняется в регистре для возврата значений и возвращается
из функции после выполнения return, а не кладется на стек.


Порядок выполнения anotherTest():
1. Переменная x объявляется и выделяется место для нее на стеке.
2. Затем, с помощью defer, объявляется анонимная функция, которая увеличивает 
 значение x на 1. Эта анонимная функция не выполняется немедленно, но сохраняет
 ссылку на переменную x. 
3. Значение x устанавливается равным 1.
4. После этого, оператор return x выполняется. Значение x (которое равно 1)
 извлекается из стека и возвращается из функции.
5. После возврата значения из функции, отложенная функция с defer выполняется, 
 и она увеличивает значение x до 2. Это изменение значения x происходит на стеке,
 но оно не влияет на результат, который уже был возвращен оператором return.
Итак, в данной функции переменные и значения управляются стеком следующим образом: сначала объявляется и устанавливается значение x, затем выполняется return, и только после этого отложенная функция с defer выполняется.

```
