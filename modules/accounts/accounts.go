package accounts

import (
	"errors"
	"fmt"
)

// Account sturct
type Account struct {
	owner   string //private
	balance int    //private
}

var errNoMoney = errors.New("can't withdraw")

// NewAccount creates Account
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account //adress를 return :복사본을 만들고싶을때
}

// Deposit x amount on your account
func (a *Account) Deposit(amonunt int) {
	a.balance += amonunt
}

// Balance of your account
func (a Account) Balance() int {
	return a.balance
}

// Withdraw X amount from your account
func (a *Account) Withdraw(amonunt int) error {
	if a.balance < amonunt {
		return errNoMoney //errors.New("can't withdraw you are poor")
	}
	a.balance -= amonunt //go는 else필요x
	return nil           // js,파이썬의 null or none과 같음
}

// ChangeOwner of the account
func (a *Account) ChnageOwner(newOwner string) {
	a.owner = newOwner
}

// Owner of the account
func (a Account) Owner() string {
	return a.owner
}

func (a Account) String() string {
	return fmt.Sprint(a.Owner(), "'s account\nHas: ", a.Balance())
} //go가 내부적으로 호출하는 method를 사용하는 방법
