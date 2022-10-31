package main

import "fmt"

type emp struct {
	name   string
	age    int
	gender string
}
type research_work struct {
	name       string
	title      string
	start_date string
	end_date   string
}

func main() {

	emp1 := emp{"ramkumar", 47, "M"}

	research1 := research_work{"ramkumar", "Test Bench", "01/01/2022", "31/12/2022"}

	fmt.Println("The values of the employee are ", emp1, "and his research work are ", research1.title, "start date is ", research1.start_date)

	new_emp := newEmp("zafar", 44)
	fmt.Println(new_emp)
	research2 := new(research_work)
	fmt.Println(research2)
}

func newEmp(name string, age int) emp {
	return emp{
		name:   name,
		age:    age,
		gender: "M",
	}
}
