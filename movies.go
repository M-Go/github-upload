package main

import "fmt"

func main() {
	var name string = "Forrest Gump"
	var lengthMinutes = 136
	var director string; director = "Robert Zemeckis"
	var basedOnTrueStory = bool(false)
	releaseYear := 1994
	
	fmt.Printf("The first movie is %v, %v minutes length, directed by %v, and released in %v. Based on true story: %v", name, lengthMinutes, director, releaseYear, basedOnTrueStory)
}