package works

import (
	"fmt"
	"strings"
	"time"
)

func Dates() {
	now := time.Now()
	year := now.Year()
	fmt.Println(year)
}

func ReplaceLetters() {
	broken := "Bob's enormous octopus observed the ocean's horizon, hoping to spot a dolphin."
	// replacer := strings.NewReplacer("o", "#")
	// fixed := replacer.Replace(broken)
	fixed := strings.NewReplacer("o", "3").Replace(broken)

	fmt.Println(fixed)
}

func pass_fail()
