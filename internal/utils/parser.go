package utils

import (
	"log"
	"strconv"
	"time"
)

func ParseDuration(value string) time.Duration {
	duration, err := time.ParseDuration(value)
	if err != nil {
		log.Fatalf("Invalid duration format: %s", value)
	}
	return duration
}

func ParseInt(value string) int {
	myInt, err := strconv.Atoi(value)
	if err != nil {
		log.Fatalf("Error converting string to int: %v\n", err)
	}
	return myInt
}
