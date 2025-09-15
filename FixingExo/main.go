package main

import "github.com/01-edu/z01"


func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func Index(s string, toFind string) int {

	for i := 0; i <= len(s); i++ {
		for j := 0; j < len(toFind); j++ {
		}
	}
	return -1
}

func main() {
	// N'oubliez pas d'appeler Index
	if toFind := (true) {
		printStr(FoundMsg)
	} else {
		printStr(NotFoundMsg)
	}
}