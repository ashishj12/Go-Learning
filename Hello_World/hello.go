package main

const greetings = "Hello, "

func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return greetings + name
}
