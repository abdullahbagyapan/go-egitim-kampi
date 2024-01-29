# Reverse Proxy

## Reverse Proxy Nedir ?

Reverse proxy, web sunucularının önünde bulunan ve istemcinin (örn. web tarayıcısı) isteklerini bu web sunucularına ileten bir sunucudur.
Reverse proxy'ler genellikle güvenliği, performansı ve güvenilirliği artırmaya yardımcı olmak için uygulanır.

## Hızlı Başlangıç

Yeni bir tane go module oluşturalım.

```bash
go mod init <module-name>
```

Daha sonra, gerekli bağımlılıkları indirelim.

```bash
go get
```

Ve son olarak, proxy serverımızı çalıştıralım.

```bash
go run main.go
```

## Kullanım

`localhost:3000/proxy` adresine GET isteği attığımız zaman, daha önce belirlediğimiz proxylere istek atıyoruz.

```go
var proxies = []ProxyInterface{
	Proxy{Destination: "127.0.0.2"},
	Proxy{Destination: "127.0.0.3"},
	Proxy{Destination: "127.0.0.4"},
	Proxy{Destination: "127.0.0.5"},
	Proxy{Destination: "127.0.0.6"},
}
```