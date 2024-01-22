# Go Routines

Goroutine, Go runtime tarafından yönetilen *hafif* bir iş parçacığıdır(thread).

> Standart threadlerin boyutları yaklaşık 1MB alan kaplarken, goroutineler yaklaşık 2kB alan kaplar. 

Goroutineler diğer iş parçacıklarıyla beraber aynı adres uzayını kullanır, bir değişkene ulaşılmak istenildiğinde iş parçacıkları birbirleriyle senkronize şekilde çalışmalıdır.

## Go Routine Tanımlama

Go routine tanımlamak için fonksiyonun başına `go` keywordunu koymak yeterlidir.

```go
go func() {
    fmt.println("Hello World")
}()
```

veya:

```go
func print(s string) {
    fmt.Println(s)
}

func main() {
    go print()
}
```

şeklinde tanımlanabilir.

## Go Routine Davranışları

Ana(*main*) goroutine, diğer goroutinelerin tamamlanmasını beklemez.

```go
for i:=0; i <10; i++ {
    go func() {
        fmt.Println(i)
    }()
}

time.Sleep(time.Second*2)   // goroutinelerin tamamlanmasını beklemek için
```

kodun çıktısı:

```go
10
10
10
10
10
10
10
10
10
10
```

> Goroutineler çalışmaya başlayana kadar hafızadaki `i` değeri `10` oldu.

Eğer `i` değerini tutan başka bir değişkeni parametre olarak verirsek:

```go
for i:=0; i <10; i++ {
    j := i
    go func() {
        fmt.Println(j)
    }()
}

time.Sleep(time.Second*2)   // goroutinelerin tamamlanmasını beklemek için
```

kodun çıktısı:

```go
9
4
5
6
7
8
1
2
3
0
```

şeklinde *nondeterministic* bir çıktı alırız, çünkü goroutinelerin çalışmak için herhangi sıraya ihtiyaçları yoktur.
Peki bu durumdan nasıl kurtuluruz, cevap `WaitGroup`.

## WaitGroup
