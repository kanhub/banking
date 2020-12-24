package main

import (
	"fmt"
	"log"

	"github.com/chakrakan/banking/accounts"
)

func main() {
	account := accounts.NewAccount("Kanisk")
	account.Deposit(1000)
	err := account.Withdraw(150)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(account.String())
}
