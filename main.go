package main

import "log"

func main() {
	log.Println("Hello world !")
	log.Println("This is the new-print branch")
	PrintMyName("Sveinn")
	PrintMyName("6")
}

func PrintMyName(name string) {
	log.Println(name)
}
