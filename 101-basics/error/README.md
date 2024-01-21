# Error

Go'da beklenmedik durumlar için `error` değerleri kullanılır.

Örnek olarak *Open* fonksiyonu dosyayı açamadığı durumda `error` dönüyor.

```go
func Open(name string) (file *File, err error)
```

Fonksiyonların kullanılmasından sonra ise `error check` yapılır.

```go
f, err := os.Open("filename.ext")
if err != nil {
    log.Fatal(err)
}
```

## Error Type

*Error* bir `interface`dir ve interfaceler herhangi bir değeri temsil edebilir.

```go
type error interface {
    Error() string
}
```

## Kullanım

```go
error := errors.New("Error ...")
```

> Error hakkında detaylı bilgi için [buraya](https://go.dev/blog/error-handling-and-go) tıklayın.