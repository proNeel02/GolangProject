package main

import "fmt"

func main() {

	// var colors map[string]string

	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#00ff00",
		"blue":  "#0000ff",
	}

	//colors := make(map[int]string)

	print(colors)
	fmt.Println("============")

	addKey(colors, "white", "#fff")
	print(colors)
}

func addKey(colors map[string]string, key string, value string) {
	colors[key] = value
}

func print(c map[string]string) {

	for key, val := range c {
		fmt.Println(key + ":" + val)
	}
}
