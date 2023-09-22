package main

import (
	"awesomeProject1/Assistant"
	"awesomeProject1/Booker"
	"awesomeProject1/Developer"
	"awesomeProject1/Operator"
	"awesomeProject1/Director"
	"fmt"
)

func main() {
	director := Director.NewDirector("Director", 260000, "ToleBi, 32")

	developer := Developer.NewDeveloper("Developer", 370000, "AbylayKhan, 43")

	booker := Booker.NewBooker("Booker", 280000, "Panfilova, 54")

	operator := Operator.NewOperator("Operator", 300000, "KazbekBi, 65")

	assistant := Assistant.NewAssistant("Assistant", 170000, "Kunayeva, 13")

	fmt.Println("Director Position:", director.GetPosition())
	director.SetPosition("Junior Director")
	fmt.Println("Director Position (Updated):", director.GetPosition())

	fmt.Println("Developer Salary:", developer.GetSalary())
	developer.SetSalary(385000)
	fmt.Println("Developer Salary (Updated):", developer.GetSalary())

	fmt.Println("Booker Position:", booker.GetPosition())

	fmt.Println("Operator Position:", operator.GetPosition())

	fmt.Println("Assistant Address:", assistant.GetAddress())

}
