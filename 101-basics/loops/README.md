# Loops

Diğer programlama dillerinden farklı olarak Go'da sadece `for` döngüsü vardır.

## Tanımlama

```go
slc_1 := []int{10, 20, 30, 40, 50, 60, 70, 80, 90}

for i, value := range slc_1 {
    fmt.Printf("index: %d, value: %d \n", i, value)
}
```
