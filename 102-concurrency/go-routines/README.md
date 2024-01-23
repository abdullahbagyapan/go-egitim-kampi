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

WaitGroup bir blocklama mekanizmasıdır, çalışan goroutine kalmayıncaya kadar bloklamaya devam eder.


### WaitGroup Tanımlama

WaitGroup `sync` paketi altındadır,

```go
wg := sync.WaitGroup{}
```

şeklinde tanımlanır.

### WaitGroup Methodları

WaitGroup altında <b>3</b> tane method vardır, bunlar:

1. Add()
2. Wait()
3. Done()

#### Add() Methodu

`Add()` methodu ile beklenecek goroutine sayısı tanımlanır.

```go
wg := sync.WaitGroup{}

wg.Add(1)
```

#### Wait() Methodu

WaitGroup sayacı 0 olana kadar blocklama devam eder.

#### Done() Methodu

WaitGroup sayacını 1 azaltır.

```go
wg := sync.WaitGroup{}

wg.Add(1)

go func() {
    time.Sleep(time.Second)
    wg.Done()
}()

wg.Wait()
```

> WaitGroup ile alakalı daha fazla bilgiye [buradan](https://pkg.go.dev/sync#WaitGroup) ulaşabilirsiniz.

## Race Condition

### Race Condition Nedir ?

Çoklu threadler, hafızadaki değere aynı anda ulaşmak ve değiştirmek durumunda kaldıklarında ortaya race condition durumu çıkar.

Peki race condition sorununu nasıl çözeriz, cevap: `Mutex`

## Mutex

Mutex, çoklu threadler bir veriyi okumak yada yazmak istediğinde verinin tutarlılığını sağlamak için veriyi kilitleyen bir mekanizmadır.

### Mutex Tanımlama

Mutex, `sync` paketinin altındadır.

```go
lock := sync.Mutex{}
```

şeklinde tanımlanır.

### Mutex Methodları

Mutex altında <b>3</b> tane method vardır, bunlar:

1. Lock()
2. TryLock()
3. Unlock()

#### Lock() Methodu

Go runtimenında hafızadaki değerlerin tek bir goroutine tarafından okuma/yazma yapılmasını sağlar, kilitmeyi sağlar.

```go
lock.Lock()
```

şeklinde kullanılır.

Kilitlenen değişkenler sıra(queue) yapısına göre işlenir (FIFO).

#### TryLock() Methodu

Eğer *deadlock* oluşturmayacak ise kilitler ve sonucunu döner.

```go
isLock := lock.TryLock()
```

şeklinde kullanılır.

#### Unlock() Methodu

Kilitlenen değişkeni serbest bırakır.

```go
lock.Unlock()
```