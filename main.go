package main

import (
	"fmt"
	"go_test/modules/accounts"
)

func main() {
	account := accounts.NewAccount("sun")
	account.Deposit(10)
	err := account.Withdraw(20) //balance가 -가 되지 않도록 예외처리를 해야한다.
	if err != nil {
		fmt.Println(err)
		//log.Fatalln(err) //-> Println(err)호출 뒤 os.Exit(1)불러 프로그램 종료
	}
	fmt.Println(account.Balance(), account.Owner())
}
