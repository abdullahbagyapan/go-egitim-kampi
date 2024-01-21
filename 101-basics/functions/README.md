# Function

Go'da 2 türlü fonksiyon tanımı vardır.

1. Normal fonksiyonlar
2. Receiver fonksiyonlar

## Normal Fonksiyonlar

Diğer programlama dillerinde olduğu gibi tanımlanan fonksiyondur.

```go
func AddTwoNumber(x , y int) int {
    return x + y
}
```

## Receiver Fonksiyonlar

Structa özel olarak tanımlanmış fonksiyonlardır. Belirtilen struct typeı dışında erişilemez.

```go
func (u *User) SetName(name string) {
    u.Name = name
}
```