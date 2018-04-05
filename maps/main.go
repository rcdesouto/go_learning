package main

import "fmt"

func main() {
	// One difference from an object is that the left side has be to all the same type
	/** In this case its strings on the left and strings on the right
	* All the items on the left that are set to strings have to be set as strings
	* and if the right is set to string then all the items have to be string. No changing it up
	 */
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
	}
	// Maps also very simliar to objects in javascript
	// In this case adding additional items similiar
	// On how I would do it in javascript
	colors["white"] = "#ffffff"

	// Code to delete a key from a map
	// delete(colors, "red")
	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println(color, hex)
	}
}
