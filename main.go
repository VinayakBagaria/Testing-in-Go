package main

import "fmt"

const englishHeroPrefix = "Hello "

func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return englishHeroPrefix + name
}

func main() {
	fmt.Println(Hello("world"))
}
