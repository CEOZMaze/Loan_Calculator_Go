package main

import (
	"flag"
	"fmt"
	"math"
	"os"
)

var (
	payment   = flag.Float64("payment", 0, "The payment amount")
	principal = flag.Float64("principal", 0, "The loan principal value")
	periods   = flag.Float64("periods", 0, "The number of months")
	interest  = flag.Float64("interest", 0, "The interest rate")
	fType     = flag.String("type", "", "The type of payment")
)

func validInput() {
	flagCount := 0
	flag.Visit(func(f *flag.Flag) {
		flagCount++
	})
	if *fType != "diff" && *fType != "annuity" || *fType == "" || flagCount < 4 || *interest <= 0 || *payment < 0 || *principal < 0 || *periods < 0 {
		fmt.Println("Incorrect parameters.")
		os.Exit(0)
	} else if *fType == "diff" && *payment > 0 {
		fmt.Println("Incorrect parameters.")
		os.Exit(0)

	}
}

func diffCalculator() {
	overPayment := 0.0
	for d := 1.0; d <= *periods; d++ {
		monthlyInterest := *interest / (12 * 100)
		cPayment := (*principal / *periods) + monthlyInterest*(*principal-(*principal*(d-1) / *periods))
		cPayment = math.Ceil(cPayment)
		overPayment += cPayment
		fmt.Printf("Month %.0f: payment is %.f\n", d, cPayment)

	}
	overPayment = overPayment - *principal
	fmt.Printf("\nOverpayment = %.f", overPayment)
}

func annCalculator(principal, payment, interest, periods float64) {
	interest = interest / (12 * 100)
	growthFactor := math.Pow(1+interest, periods)
	adjustedGrowth := growthFactor - 1
	scalingFactor := interest * growthFactor

	if principal <= 0 {
		resultPrincipal := payment * adjustedGrowth / scalingFactor
		resultPrincipal = math.Floor(resultPrincipal)
		totalPayment := payment * periods
		overpayment := totalPayment - resultPrincipal
		fmt.Printf("Your loan principal = %d!\nOverpayment = %d", int(math.Round(resultPrincipal)), int(math.Round(overpayment)))
	} else if principal > 0 {

		annuityPayment := principal * scalingFactor / adjustedGrowth
		annuityPayment = math.Ceil(annuityPayment)
		totalPayment := annuityPayment * periods
		totalPayment = math.Ceil(totalPayment)
		overpayment := totalPayment - principal
		overpayment = math.Ceil(overpayment)
		fmt.Printf("Your loan annuity = %d!\nOverpayment = %d", int(annuityPayment), int(overpayment))
	}
}

func calculateMonthly(principal, interest, periods float64) {
	interest = interest / (12 * 100)
	resultPayment := principal * (interest * math.Pow(1+interest, periods)) / (math.Pow(1+interest, periods) - 1)
	resultPayment = math.Ceil(resultPayment)
	fmt.Printf("Your monthly payment = %d!", int(resultPayment))
	//return int(math.Ceil(resultPayment))

}

func loanPrincipal(payment, interest, periods float64) {

	interest = interest / (12 * 100)
	growthFactor := math.Pow(1+interest, periods)
	adjustedGrowth := growthFactor - 1
	scalingFactor := interest * growthFactor

	resultPrincipal := payment * adjustedGrowth / scalingFactor
	fmt.Printf("Your loan principal = %d!", int(math.Round(resultPrincipal)))

}

func numberPayments(principal, payment, interest float64) {
	interest = interest / (12 * 100)

	resultPeriods := math.Log(payment/(payment-interest*principal)) / math.Log(1+interest)

	resultPeriods = math.Ceil(resultPeriods)

	years := int(math.Floor(resultPeriods / 12))
	months := int(resultPeriods) % 12

	overpayment := payment*resultPeriods - principal
	overpayment = math.Ceil(overpayment)

	if years > 0 && months > 0 {
		fmt.Printf("It will take %d years and %d months to repay this loan!\nOverpayment = %.f", years, months, overpayment)
	} else if years > 0 && months == 0 {
		fmt.Printf("It will take %d years to repay this loan!\nOverpayment = %.f", years, overpayment)
	} else if years == 0 && months > 0 {
		fmt.Printf("It will take %d months to repay this loan!\nOverpayment = %.f", months, overpayment)
	}
}

func whichOne() {
	if *fType == "annuity" && *periods > 0 {
		annCalculator(*principal, *payment, *interest, *periods)

	} else if *fType == "annuity" && *periods == 0 {
		numberPayments(*principal, *payment, *interest)
	} else if *fType == "diff" {
		diffCalculator()
	} else if *payment == 0.0 {
		calculateMonthly(*principal, *interest, *periods)
	} else if *principal == 0.0 {
		loanPrincipal(*payment, *interest, *periods)
	} else if *periods == 0.0 {
		numberPayments(*payment, *interest, *principal)
	}
}

func main() {

	flag.Parse()
	validInput()
	whichOne()
}
