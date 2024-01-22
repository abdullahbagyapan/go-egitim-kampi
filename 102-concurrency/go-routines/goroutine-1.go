package main

func main() {

	go func() {
		println("Hello")
	}()
	println("World")

	// without time sleep print:	'World'

	// with time sleep print:		'Hello World' or 'World Hello'	(nondeterministic behavior)
	//time.Sleep(time.Second)
}
