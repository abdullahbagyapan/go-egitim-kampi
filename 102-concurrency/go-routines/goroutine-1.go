package main

func main() {

	go func() {
		println("Hello")
	}()
	println("World")

	// without time sleep print:	'World'

	// with time sleep print:		'Hello World' or 'World Hello'	(nondeterministic behaviour)
	//time.Sleep(time.Second)
}
