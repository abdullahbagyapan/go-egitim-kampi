# Pointers

Pointer bir değerin hafızadaki adresini tutar.

## Tanımlama

Bir değişkenin pointer tipinde olmasını istiyorsak değişken türünün önüne `*` koymak yeterli olacaktır.

```go
var pstr *string
```

## Kullanım

Pointerlar, bir değişkenin adresini tutabilmeleri için o değişkenin adresini bilmeleri gerekir, `&` işareti ile bir değişkenin hafızadaki adresi öğrenilir.

```go
var pstr *string

str := "Hello World"

pstr = &str
```

> Not: Eğer tutulan adresteki değer değişirse pointerinde değeri değişmiş olur

```go
var pstr *string

str := "Hello World"
pstr = &str     // *pstr: "Hello World"

str = "Golang"  // *pstr: "Golang"
```