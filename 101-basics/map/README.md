# Map

Mapler `key:value` şeklinde değer tutan veri tipleridir.

## Tanımlama

```go
var map_1 map[int]string
```

Go'da *map*ler referans tutuculardır, ve `map_1` değişkeni ise daha referans atanmadığı için `nil`dir.
Eğer *map* tanımlaması yapılacaksa doğru olan *built in* ile gelen `make()` fonksiyonu kullanılmasıdır.

```go
map_2 = make(map[string]int)
```

> Daha fazla bilgiye [buradan](https://go.dev/blog/maps) ulaşabilirsiniz

## Map Üzerinden İşlemler

### Değer Atama

Maplerde değer atama işlemi aynı `array`lerde olduğu gibi değişken üzerinden yapılır.

```go
map_1 := make(map[string]int)

map_1["test"] = 1
```

### Değer Silme

Maplerde değer silme işlemi *built in* gelen `delete()` fonksiyonu ile yapılır.

```go
delete(map_1, 2)
```

### Değer Kontrol Etme

```go
fred, exists := names[0]    // exists: bool tipinde bir değişkendir
```