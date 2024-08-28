package main

import (
	"fmt"
	"learngowithtests/helloworld"
	"learngowithtests/integers"
	"learngowithtests/iteration"
	"learngowithtests/pointers"
	racer "learngowithtests/select"
	"learngowithtests/sum"
)

func main() {
	greetWillInEnglish := helloworld.Hello("Will", "English", "!")
	addedIntegers := integers.Add(1, 2)
	wallet := &pointers.Wallet{}

	wallet.Deposit(pointers.Bitcoin(100))
	fmt.Println(greetWillInEnglish)
	fmt.Println("1 + 2 =", addedIntegers)
	fmt.Println("Repeat \"a\" five times:", iteration.Repeat("a", 5))
	fmt.Println("Sum of [1, 2, 3, 4, 5]:", sum.Sum([]int{1, 2, 3, 4, 5}))
	fmt.Println("Wallet balance:", wallet.Balance())
	fmt.Println("Withdraw 10 BTC from wallet...")
	wallet.Withdraw(10)
	fmt.Println("Wallet balance:", wallet.Balance())
	wallet.Withdraw(pointers.Bitcoin(20))
	winner, err := racer.Racer("http://www.facebook.com", "http://www.google.com")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Speed test winner:", winner)
	}
}
