package main

import (
	"admodev/pdfgener/pkg"
	"fmt"
)

func main() {
	fmt.Println("Initializing server...")

	if err := pkg.InitializeServer(); err != nil {
		fmt.Printf("Error initializing server: %v\n", err)
	}
}
