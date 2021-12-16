package main

import (
	"log"
)

func main() {
	// Template request
	res, err := TemplateEmail()
	if err != nil {
		log.Printf("error: %v", err)
	} else {
		log.Print(res)
	}

	// Raw email request
	res, err = RawEmail()
	if err != nil {
		log.Printf("error: %v", err)
	} else {
		log.Print(res)
	}

	// Standard email
	res, err = StdEmail()
	if err != nil {
		log.Printf("error: %v", err)
	} else {
		log.Print(res)
	}
}
