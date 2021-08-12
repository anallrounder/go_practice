package accounts

import "fmt"

// Account sturct
type Account struct {
	owner   string //private
	balance int    //private
}

// NewAccount creates Account
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account //adress를 return :복사본을 만들고싶을때
}

// Deposit x amount on your account
func (a *Account) Deposit(amonunt int) {
	fmt.Println("Gonna deposit", amonunt)
	a.balance += amonunt
}

// Balance of your account
func (a Account) Balance() int {
	return a.balance
}

// Withdraw X amount from your account
func (a *Account) Withdraw(amonunt int) {
	fmt.Println("Gonna withdraw", amonunt)
	a.balance -= amonunt
}
