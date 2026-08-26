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
