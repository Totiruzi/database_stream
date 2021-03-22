package main

import (
	"odbcstream/odbcstream"
	// "fmt"
)

func main() {
	odbcstream.InitialiseDBConnection("psql",	"france")
	// fmt.Println("Hello Go")
}
