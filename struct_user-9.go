package main

import "fmt"

type USalary struct {
	Basic, HRA, TA float64
}

type User struct {
	FirstName, LastName, Email string
	Age                        int
	MonthlySalary              []USalary
}

func (e User) UserInfo() string {
	fmt.Println(e.FirstName, e.LastName)
	fmt.Println(e.Age)
	fmt.Println(e.Email)
	for _, info := range e.MonthlySalary {
		fmt.Println("==")
		fmt.Println(info.Basic)
		fmt.Println(info.HRA)
		fmt.Println(info.TA)
	}
	return "--"
}

func main() {

	e := User{
		FirstName: "Samit",
		LastName:  "Verma",
		Email:     "samitverma279@gmail.com",
		Age:       23,
		MonthlySalary: []USalary{
			USalary{
				Basic: 15000.00,
				HRA:   5000.00,
				TA:    2000.00,
			},
			USalary{
				Basic: 16000.00,
				HRA:   5000.00,
				TA:    2100.00,
			},
			USalary{
				Basic: 17000.00,
				HRA:   5000.00,
				TA:    2200.00,
			},
		},
	}

	fmt.Println(e.UserInfo())
}
