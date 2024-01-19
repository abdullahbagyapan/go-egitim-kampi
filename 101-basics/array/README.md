# Array

Diğer programlama dilleri gibi Go'da da `array` tanımlama benzer şekilde yapılır.

```go
var arr_3 [3]int
var arr_5 = [5]int{1, 2, 3, 4, 5}
```

## 'make()' Methodu ile Array Tanımlama

Fonksiyon içinde tanımlarken ise `make()` methodu kullanılır.

```go
func main() {
	arr_6 := make([]int, 6)
}
```



## Değer Atama

Değer atama işlemi, diğer programlama dilleri gibi değişken üzerinden yapılır.

```go
arr_6[0] = 10
```
