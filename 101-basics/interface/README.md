# Interface

Interfaceler method imzalarıdır. Interface içinde, methodları implemente eden herhangi bir türü tutabilir.

```go
type User interface {
	Like()
}

func (u *user) Like() {

	like := new(Like)
	u.Likes = append(u.Likes, like)
}
```


## Boş Interface

Boş interfaceler (`interface{}`) bütün türler tarafından implemente edilir.
Bu yüzden boş interfaceler herhangi bir değeri temsil ederler.

```go
func printValue(i interface{}) {

    fmt.Println(i)
}
```

Yukarıda belirtilen `printValue()` fonksiyonu, boş interfaceden dolayı parametre olarak her türü alabilir.
