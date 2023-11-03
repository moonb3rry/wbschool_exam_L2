Что выведет программа? Объяснить вывод программы.

```go
package main

type customError struct {
	msg string
}

func (e *customError) Error() string {
	return e.msg
}

func test() *customError {
	{
		// do something
	}
	return nil
}

func main() {
	var err error
	err = test()
	if err != nil {
		println("error")
		return
	}
	println("ok")
}
```

Ответ:
```
ok

Здесь мы объявляем структуру customError, которая реализует интерфейс error.
Функция test возвращает указатель на customError, и в этом конкретном случае
возвращает nil, таким образом, в main'e не зайдем в тело if'a, и выведется ok.
```
