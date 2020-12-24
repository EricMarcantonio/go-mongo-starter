package main

import "log"

// CheckErr - Panic's the application, allows trace to err where it was invoked
func CheckErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
