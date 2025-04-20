package main

import (
	"fmt"
	"log"
)

func handleError(source string, err error, fatal bool) {
	feedbackText := fmt.Sprintf("Error in %s: %v\n", source, err)
	if fatal {
		log.Fatal(feedbackText + "Error is fatal. Escaping program.")
	}
	fmt.Println(feedbackText + "Error not fatal. Continuing.")
}
