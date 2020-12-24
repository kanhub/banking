package main

import (
	"fmt"
	"log"

	"github.com/chakrakan/banking/accounts"
	"github.com/chakrakan/banking/dict"
)

func main() {
	// Bank stuff
	fmt.Println("Banking")
	account := accounts.NewAccount("Kanisk")
	account.Deposit(1000)
	err := account.Withdraw(150)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(account.String())
	fmt.Println("----------")

	// Dictionary stuff
	fmt.Println("Dictionary Stuff")
	dictionary := dict.Dictionary{"first": "First item"}
	dictionary["hello"] = "hello world"
	fmt.Println("Initial Dictionar: ", dictionary)

	// Searching item in dictionary
	definitionItemIn, errItemIn := dictionary.Search("first")
	// item not in dictionary
	definitionItemNotIn, errItemNotIn := dictionary.Search("last")
	if errItemIn != nil {
		fmt.Println(errItemIn)
	}
	if errItemNotIn != nil {
		fmt.Println(errItemNotIn)
	}

	fmt.Println(definitionItemIn)
	fmt.Println(definitionItemNotIn)

	// Adding
	errAdd := dictionary.Add("hello", "hola")
	if errAdd != nil {
		fmt.Println(errAdd)
	}
	fmt.Println("After Add: ", dictionary)

	// Update
	errUpdate := dictionary.Update("hello", "hola")
	if errUpdate != nil {
		fmt.Println(errUpdate)
	}
	fmt.Println("After Update: ", dictionary)

	// Delete
	dictionary.Delete("hello")
	fmt.Println("After Delete: ", dictionary)
}
