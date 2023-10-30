/*

Account
	attributes:
		account number
		account holder's name
		balance
	methods:
		deposit
		withdraw
		checkbalance

Saving account
	inherits from account

	attributes:
		interest rate

	methods:
		calculate the interest from a certain period and add it to the account balance

Current Account:
	inherits from account

	attributs:
		overdraft limit

	methods:
		check if withdrawal can be make without exceeding the overdraft limit

*/

package main

import (
	"fmt"
)

type Person struct {
	Name    string
	Age     int
	Account *Account
}

type Account struct {
	AccountNumber int
	Balance       int
}

type SavingAccount struct {
	Account
	InterestRate float64
}

func main() {

	me := createPerson()

	me.setPerson("Sanzhar", 19)
	me.setAccount(123, 0)

	myAccount := me.Account
	

	myAccount.deposit(123)
	myAccount.withdraw(23)
	fmt.Println(me.Account.checkbalance())

	fmt.Println("Lets deposit balance to saving account and get 5% every month")

	mySavingAcc := &SavingAccount{
		Account: *myAccount,
		InterestRate: 5,
	}

	fmt.Println("My balance", mySavingAcc.checkbalance())
	mySavingAcc.calculate(5)
	fmt.Println("After 5 months")
	fmt.Println("My balance", mySavingAcc.checkbalance())
}


func (s *SavingAccount) calculate(months int) {
	interest := (float64(s.Balance) * s.InterestRate / 100) * float64(months)
	s.deposit(int(interest))
}

func (a *Account) checkbalance() int {
	return a.Balance
} 

func (a *Account) withdraw(amount int) {
	a.Balance -= amount
} 

func (a *Account) deposit(amount int) {
	a.Balance += amount
} 

func (p *Person) setAccount(number, balance int) {
	p.Account.AccountNumber = number
	p.Account.Balance = balance
}

func (p *Person) setPerson(name string, age int) {
	p.Name = name
	p.Age = age
}

func createPerson() *Person {
	return &Person{
		Account: &Account{},
	}
}
