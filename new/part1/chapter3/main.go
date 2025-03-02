package main

import (
	"errors"
	"fmt"
)

const (
	goodScore      = 450
	lowScoreRatio  = 10
	goodScoreRatio = 20
)

var (
	ErrCreditScore = errors.New("Invalid credit score")
	ErrIncome      = errors.New("Income invalid")
	ErrLoanAmount  = errors.New("loan amount invalid")
	ErrLoanTerm    = errors.New("loan term is mot a multiple of 12")
)

func checkLoan(creditScore int, income float64, loanAmount float64, loanTerm float64) error {
	// Good credit svore gets a better rate
	interest := 20.0
	if creditScore >= goodScore {
		interest = 15.0
	}

	// validate score
	if creditScore < 1 {
		return ErrCreditScore
	}

	// validate income
	if income < 1 {
		return ErrIncome
	}

	// validate loan amount
	if loanAmount < 1 {
		return ErrLoanAmount
	}

	// validate term
	if loanTerm < 1 || int(loanTerm)%12 != 0 {
		return ErrLoanTerm
	}

	rate := interest / 100
	payment := ((loanAmount * rate) + loanAmount) / loanTerm
	totalInterest := (payment * loanTerm) - loanAmount

	approved := false
	if income > payment {
		ratio := (payment / income) * 100
		if creditScore >= goodScore && ratio < goodScoreRatio {
			approved = true
		} else if ratio < lowScoreRatio {
			approved = false
		}
	}

	fmt.Println("Credit Score :", creditScore)
	fmt.Println("Income :", income)
	fmt.Println("Loan Amount :", loanAmount)
	fmt.Println("Loan Term :", loanTerm)
	fmt.Println("Monthly Payment :", payment)
	fmt.Println("Rate :", rate)
	fmt.Println("Total Cost :", totalInterest)
	fmt.Println("Approved :", approved)

	return nil
}

func main() {
	// Approved
	fmt.Println("Applicant 1")
	fmt.Println("-----------")
	err := checkLoan(500, 1000, 1000, 24)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Denied
	fmt.Println("Applicant 2")
	fmt.Println("-----------")
	err = checkLoan(350, 1000, 10000, 24)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
}
