package main

import (
	"fmt"
	"go_test/modules/accounts"
)

func main() {
	account := accounts.NewAccount("sun")
	account.Deposit(10)
	fmt.Println(account)
	fmt.Println(account.Balance())
	account.Withdraw(20)
	fmt.Println(account.Balance())
}
