# Variables

Go dilinde farklı şekillerde değişken tanımı yapılabilir.

## Using 'Var' Keyword

```go
var name string

func main() {
    var age = 18
    var foo = "bar"
}
```

## Using 'Const' Keyword

```go
const PI = 3.14

func main() {
    PI = 3  // değiştirilemez (immutable)
}
```

## Using ':=' Operator

Herhangi bir _tip_ tanımlaması yapmak zorunda kalmadan değişken tanımlanır.

> Not: Fonksiyonların dışında <b>kullanılamaz</b>.

```go
func main() {
    name := "GO TURKIYE"
}
```

## Multiple Definition

Birden fazla değişken tanımlanmak istendiği zaman kullanılır.

```go
var(
    id  int
    name string
)

const(
    PI      float = 3.14
    GRAVITY float = 9.8
)
```
