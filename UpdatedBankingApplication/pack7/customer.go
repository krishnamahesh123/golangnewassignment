package pack7

import "fmt"

type Bank interface {
	AddMoney(accno int, depositmoney int)
	WithdrawMoney(accno int, withdrawmoney int)
	CloseAccount()
}

type Customer struct {
	Firstname      string
	Lastname       string
	Accountnumber  int
	Accountbalance int
	Email          string
	IsOpen         bool
}

func InitialiseCustomer(Fname string, Lname string, Accno int, Accb int, email string) *Customer {
	cust := Customer{
		Firstname:      Fname,
		Lastname:       Lname,
		Accountnumber:  Accno,
		Accountbalance: Accb,
		Email:          email,
		IsOpen:         true,
	}
	return &cust
}

func (c *Customer) AddMoney(accno int, depositmoney int) {
	if c.Accountnumber == accno {
		balance := c.Accountbalance
		fmt.Printf("The amount want to deposit is %d to the AccountNumber %d\n", depositmoney, c.Accountnumber)
		updated_balance := balance + depositmoney
		c.Accountbalance = updated_balance
	} else {
		fmt.Println("The given account number does not exist")
	}

}
func (c *Customer) WithdrawMoney(accno int, withdrawmoney int) {
	balance := c.Accountbalance
	fmt.Printf("The amount want to withdraw is %d from the AccountNumber %d\n", withdrawmoney, c.Accountnumber)
	updated_balance := CheckBalance(balance, withdrawmoney)
	c.Accountbalance = updated_balance

}
func CheckBalance(balance int, withdraw int) int {
	if balance > withdraw {
		updated_balance := balance - withdraw
		return updated_balance
	} else {
		fmt.Println("Unable to withdraw/Transfer money because of low balance")
		return 0
	}

}

func TransferMoney(TransferAmount int, cust3 *Customer, cust4 *Customer) {
	fmt.Printf("Transfer Money of %d rupees from Account Number %d to the Account Number %d\n", TransferAmount, cust3.Accountnumber, cust4.Accountnumber)
	balance := cust3.Accountbalance
	updated_balance := CheckBalance(balance, TransferAmount)
	cust3.Accountbalance = updated_balance
	bal := cust4.Accountbalance
	cust4.Accountbalance = bal + TransferAmount
}
func (c *Customer) CloseAccount() {
	check := c.IsOpen
	if !check {
		fmt.Println("Account is already closed")
	} else {
		check = false
		c.IsOpen = check
		fmt.Printf("Account Number %d is closed Successfully \n", c.Accountnumber)
	}

}
