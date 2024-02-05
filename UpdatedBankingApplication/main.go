package main

import (
	"UpdatedBankingApplication/pack7"
	"UpdatedBankingApplication/pack8"
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var FirstName, LastName, Email string
	var AccountNumber, AccountBalance, accno, depositmoney, withdrawmoney, TransferAmount, empid int
	cust1 := pack7.Customer{Firstname: "Krishna", Lastname: "Mahesh", Accountnumber: 1234, Accountbalance: 100, Email: "krishnamaheshecea@gmail.com", IsOpen: true}
	cust2 := pack7.Customer{Firstname: "Manikanta", Lastname: "Venkata", Accountnumber: 5678, Accountbalance: 200, Email: "manikantakothuri@gmail.com", IsOpen: true}
	cust3 := pack7.Customer{Firstname: "Rambabu", Lastname: "Simhadri", Accountnumber: 1486, Accountbalance: 300, Email: "rambabusimhadri@gmail.com", IsOpen: true}
	var cust4 *pack7.Customer
	emp := pack8.InitialiseEmployee()
	for i := 0; i < 5; i++ {
		fmt.Println("Choose an option:")
		fmt.Println("1. Customer")
		fmt.Println("2. Employee")
		fmt.Print("Enter option (1 or 2): ")
		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			fmt.Println("Please Enter Account number to login")
			fmt.Scanf("%d\n", &accno)
			if accno == cust1.Accountnumber || accno == cust2.Accountnumber || accno == cust3.Accountnumber {
				fmt.Println("logged as a customer")
				for i := 0; i < 6; i++ {
					fmt.Println("Choose an option:")
					fmt.Println("1. Create New Customer Account")
					fmt.Println("2. Deposit Money")
					fmt.Println("3. Withdraw Money ")
					fmt.Println("4. View AccountBalance of all Customers")
					fmt.Println("5. Transfer Money ")
					fmt.Println("6. Close Account")
					fmt.Print("Enter option (1 or 2 or 3 or 4 or 5 or 6): ")
					scanner.Scan()
					option := scanner.Text()
					switch option {
					case "1":
						fmt.Println("Enter Firstname")
						fmt.Scanf("%s\n", &FirstName)
						fmt.Println("Enter Lastname")
						fmt.Scanf("%s\n", &LastName)
						fmt.Println("Enter Account Number")
						fmt.Scanf("%d\n", &AccountNumber)
						fmt.Println("Enter Account Balance")
						fmt.Scanf("%d\n", &AccountBalance)
						fmt.Println("Enter Email")
						fmt.Scanf("%s\n", &Email)

						cust4 = pack7.InitialiseCustomer(FirstName, LastName, AccountNumber, AccountBalance, Email)
						fmt.Printf("New Account created successfully to  %s \n", cust4.Firstname)
					case "2":
						fmt.Println("Enter AccountNumber to deposit")
						fmt.Scanf("%d\n", &accno)
						fmt.Println("Enter Money to deposit")
						fmt.Scanf("%d\n", &depositmoney)
						if cust1.Accountnumber == accno {
							cust1.AddMoney(accno, depositmoney)
							fmt.Printf("The total amount after deposit is %d to the Account Number %d\n", cust1.Accountbalance, cust1.Accountnumber)
						}

						if cust2.Accountnumber == accno {
							cust2.AddMoney(accno, depositmoney)
							fmt.Printf("The total amount after deposit is %d to the Account Number %d\n", cust2.Accountbalance, cust2.Accountnumber)
						}
						if cust3.Accountnumber == accno {
							cust3.AddMoney(accno, depositmoney)
							fmt.Printf("The total amount after deposit is %d to the Account Number %d\n", cust3.Accountbalance, cust3.Accountnumber)
						}
						if cust4.Accountnumber == accno {
							cust3.AddMoney(accno, depositmoney)
							fmt.Printf("The total amount after deposit is %d to the Account Number %d\n", cust4.Accountbalance, cust4.Accountnumber)
						}

					case "3":
						fmt.Println("Enter AccountNumber to withdraw")
						fmt.Scanf("%d\n", &accno)
						fmt.Println("Enter Money to withdraw")
						fmt.Scanf("%d\n", &withdrawmoney)
						if cust1.Accountnumber == accno {
							cust1.WithdrawMoney(accno, withdrawmoney)
							fmt.Printf("The total amount after withdraw is %d from the Account Number %d\n", cust1.Accountbalance, cust1.Accountnumber)
						}
						if cust2.Accountnumber == accno {
							cust2.WithdrawMoney(accno, withdrawmoney)
							fmt.Printf("The total amount after withdraw is %d from the Account Number %d\n", cust2.Accountbalance, cust2.Accountnumber)
						}
						if cust3.Accountnumber == accno {
							cust3.WithdrawMoney(accno, withdrawmoney)
							fmt.Printf("The total amount after withdraw is %d from the Account Number %d\n", cust3.Accountbalance, cust3.Accountnumber)
						}
						if cust4.Accountnumber == accno {
							cust4.WithdrawMoney(accno, withdrawmoney)
							fmt.Printf("The total amount after withdraw is %d from the Account Number %d\n", cust4.Accountbalance, cust4.Accountnumber)
						}

					case "4":
						fmt.Println("The following are the Account Balance of all the customers")
						fmt.Printf("The account Balance of account number %d is %d\n", cust1.Accountnumber, cust1.Accountbalance)
						fmt.Printf("The account Balance of account number %d is %d\n", cust2.Accountnumber, cust2.Accountbalance)
						fmt.Printf("The account Balance of account number %d is %d\n", cust3.Accountnumber, cust3.Accountbalance)
						fmt.Printf("The account Balance of account number %d is %d\n", cust4.Accountnumber, cust4.Accountbalance)

					case "5":
						fmt.Println("Enter the amount to transfer")
						fmt.Scanf("%d\n", &TransferAmount)
						pack7.TransferMoney(TransferAmount, &cust3, cust4)
						fmt.Printf("The total balance in Account number %d after transfer of %d rupees to Account Number %d  is %d\n", cust3.Accountnumber, TransferAmount, cust4.Accountnumber, cust3.Accountbalance)
						fmt.Printf("The total balance in Account number %d after transfer of %d rupees from Account Number %d is %d\n", cust4.Accountnumber, TransferAmount, cust3.Accountnumber, cust4.Accountbalance)
					case "6":

						cust4.CloseAccount()
					default:
						fmt.Println("Enter valid number")

					}
				}
			} else {
				fmt.Println("login failed. Please try again")
			}
		case "2":
			fmt.Println("Please Enter Employee id to login")
			fmt.Scanf("%d\n", &empid)
			err := pack8.LoginAsEmployee(emp, empid)
			if err != nil {
				fmt.Println(err)
			} else {

				fmt.Println("The following are the employee details")
				for _, details := range emp {
					fmt.Println(details)
				}
			}
		default:
			fmt.Println("Please enter a valid input ")
		}
	}

}
