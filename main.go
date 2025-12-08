package main

import (
	"fmt"
	"log"
	"os"
)

const (
	readmeFile        = "README.md"
	careerITStartYear = 2009
)

func main() {
	fmt.Println("Updater XP")
	fmt.Printf("Career Start: %d\n", careerITStartYear)
	fmt.Println("================================")

	filePath := readmeFile
	if len(os.Args) > 1 {
		filePath = os.Args[1]
	}

	if err := updateReadme(filePath); err != nil {
		log.Printf("could not update the readme: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Update completed successfully")
	fmt.Println("================================")
}
