# Channels

Channellar, değer alıp, gönderme işlemi yapan bir kanaldır, `<-` işareti ile işlemler yapılır.
Yapısı gereği channel'lar, diğer taraf hazır olana kadar atama ve okuma işlemlerini blocklar.

> Verinin akış yönü *ok*un yönüne göredir.

## Channel Türleri

2 tür channel vardır,

1. Unbuffered
2. Buffered

### Unbuffered Channel

Unbuffered Channel'larda, alıcı hazır olana kadar atama işlemi yapılmaz, blocklanır.
Doğrudan bir iletim vardır, senkron olarak veri iletilir.

#### Kullanım

Maps ve Slicelar gibi channellarda kullanılmadan önce oluşturulmalıdır.

```go
ch := make(chan int)    // channel oluşturma

ch <- 4                 // channel'a değer atama

value := <-ch           // channel'dan değer okuma

close(ch)               // channel kapatma
```

### Buffered Channel

Buffered Channel'lar unbuffered channel'lardan farklı olarak birden fazla girdi tutabilir, bu işlemde herhangi bir blocklama yaşanmaz.

#### Kullanım

```go
ch := make(chan int,2)    // channel oluşturma

ch <- 4                 // channel'a değer atama
ch <- 5                 // channel'a değer atama

value1 := <-ch           // channel'dan değer okuma
value2 := <-ch           // channel'dan değer okuma

close(ch)               // channel kapatma
```