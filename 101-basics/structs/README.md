# Struct

Structlar, `class`lara benzeselerde aynısı değillerdir, yeni yapılar oluşturmak için kullanılır. Verileri gruplamada yardımcı olur.

## Tanımlama

Go'da struct tanımı, `type *Struct'ın Adı* struct {}` şeklinde yapılır.

```go
type User struct {

}
```

## Değişken Tanımlama

Structlar farklı değişkenler tutabilir.

```go
type User struct {
    Id int
    Name string
    Age int
}
```

## Nesne Oluşturma

Tanımlanan struct nesnesini oluşturmanın birden fazla yolu vardır.

```go
user_1 := User{
    Id:1,
    Name:"user_1"
    Age:18
}
```

```go
user_2 := New(User)

user_2.Id = 2
Name:"user_2"
Age:20
```