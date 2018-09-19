package main

const prefix = "hello,"

func sayHello(name string) string {
	return prefix + name
}

func main() {
	println(sayHello("world"))
}
