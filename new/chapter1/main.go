package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func generateRandomNumber() {
	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(5) + 1
	stars := strings.Repeat("*", r)
	fmt.Println((stars))
}

func patientDetails() {
	firstName := "Bob"
	lastName := "Smith"
	age := 34
	hasPeanutAllergy := false

	fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println(age)
	fmt.Println(hasPeanutAllergy)
}

func MultipleVariables() {
	Debug, LogLevel, startupTime := false, "info", time.Now()

	fmt.Println(Debug, LogLevel, startupTime)
}

func getConfig() (bool, string, time.Time) {
	return false, "info", time.Now()
}

func operators() {
	visits := 15
	fmt.Println("First visit :", visits == 1)
	fmt.Println("Retunr visit :", visits != 1)

	// check if there are a silver member
	fmt.Println("Silver member :", visits >= 10 && visits < 21)

	// gold member
	fmt.Println("Gold member :", visits > 20 && visits <= 30)

	// platnum member
	fmt.Println("Platinum member :", visits > 30)
}

func main() {
	operators()
}
