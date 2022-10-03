package main

import (
	"ch26/ex26.1/appB/exB1/bankaccount"
	"fmt"
)

func main() {
	account := bankaccount.NewAccount() // 1. 계좌 생성
	account.Deposit(1000)
	fmt.Println(account.Balance())
}
