package pointers

import "fmt"

func LearnPointers() {
	age := 23
	var agePointer *int = &age
	fmt.Println(getAdultYears(agePointer))
	getTeenageYears(&age)
	fmt.Println("age:", age)
	fmt.Printf("age:%d\n", age)
}

func getAdultYears(age *int) int {
	return *age - 18
}

func getTeenageYears(age *int) {
	*age = *age - 13
}
