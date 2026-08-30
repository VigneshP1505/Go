package structs

import (
	"fmt"
	"time"

	"example.com/first-app/structs/occupation"
)

type user struct {
	firstName string
	lastName  string
	birthDate string
	createdAt time.Time
}

type Employees struct {
	Count     int64
	startDate time.Time
	contracts int64
}

func StructsLearn() {
	companies := Employees{1000, time.Date(1971, time.March, 5, 23, 12, 19, 1200, time.UTC), 1200}
	fmt.Printf("%v", companies)
}

func Structs() {
	var appUser user
	appUser = user{
		firstName: "vignesh",
		lastName:  "Pugazhendhi",
		birthDate: "15-05-1997",
		createdAt: time.Now(),
	}
	outputPersonalData(appUser)
	printOutputData(&appUser)
	appUser.printOutput()
	var dob = appUser.getDOB()
	fmt.Printf("dob:%v", dob)
	var newUser *user = appUser.getClone()
	fmt.Print("\nprinting new cloned user:")
	newUser.printOutput()

	//occupation
	var userOccupation occupation.Occupation = *occupation.CreateOccupation()
	userOccupation.Title = "tesla"
	newOccupation, err := occupation.SpawnOccupation("Swe", "$10000", true)
	if err != nil {
		fmt.Print(err)
	} else {
		newOccupation.PrintOccupation()
	}

}

func outputPersonalData(u user) {
	fmt.Println(u.firstName, u.birthDate, u.createdAt)
}

func printOutputData(u *user) {
	fmt.Println((*u).firstName + " " + (*u).lastName)
}

func (u user) printOutput() {
	fmt.Println(u.firstName + " " + u.lastName)
}

func (u user) getDOB() string {
	return u.birthDate
}

func (u user) getName() string {
	return u.firstName + " " + u.lastName
}

func (u user) getClone() *user {
	return &user{
		firstName: u.firstName,
		lastName:  u.lastName,
		birthDate: u.birthDate,
		createdAt: time.Now(),
	}
}

// struct is a Go's way to create a data type of multiple fields
// You can attach methods to a struct

type Order struct {
	ID         int
	CustomerID int
	Amount     float64
	Status     string
}

type Customer struct {
	ID   int
	Name string
}

type OnlineOrder struct {
	ID       int
	Customer Customer
	Amount   float64
	Status   string
}

func (o Order) IsPending() bool {
	return o.Status == "PENDING"
}

func (o Order) UpdateByValue(status string) {
	o.Status = status
}

func (o *Order) UpdateOrder(status string) {
	o.Status = status
}

func CreateOrders() {
	order := Order{
		ID:         101,
		CustomerID: 112,
		Amount:     29.99,
		Status:     "PENDING",
	}
	fmt.Println(order.IsPending())
	order.UpdateOrder("PICKED_UP")
	fmt.Println(order.IsPending())
	order.Status = "PENDING"
	order.UpdateByValue("PICKED_UP")
	fmt.Println(order.IsPending())
}

//interfaces
//Any type that has all interface methods satisfies that interface

type PaymentProcessor interface {
	Process(amount float64) error
}

type StripeProcessor struct{}
type PayPalProcessor struct{}

func (s StripeProcessor) Process(amount float64) error {
	fmt.Println("stripe processor", amount)
	return nil
}

func (p PayPalProcessor) Process(amount float64) error {
	fmt.Println("Paypal processor", amount)
	return nil
}
