package datastructure

import "fmt"

type floatMap map[string]float64

func (m floatMap) output() {
	fmt.Println(m)
}

func Arrays() {
	prices := [10]float32{10.23, 12.45, 13.14}
	fmt.Println(prices)
	var interviewers [4]string = [4]string{"sharanyu", "pillai", "data", "engineer"}
	fmt.Println(interviewers)
	interviewers[1] = "nyuer"
	fmt.Println(interviewers)
	slicedInterviewers := interviewers[1:3]
	fmt.Println(slicedInterviewers, "size:", len(slicedInterviewers))
	slicedInterviewersTwo := interviewers[1:]
	fmt.Println(slicedInterviewersTwo)
	slicedInterviewers[1] = "test"
	fmt.Println(slicedInterviewers)
	fmt.Println(interviewers, len(slicedInterviewers), cap(slicedInterviewers))

	var reducingFactor [4]float64 = [4]float64{12.32}
	fmt.Println("reducingFactor:", reducingFactor, len(reducingFactor), cap(reducingFactor))

	var slicedReducingFactor = reducingFactor[1:2]
	fmt.Println("slicedReducingFactor:", slicedReducingFactor, len(slicedReducingFactor), cap(slicedReducingFactor))

	userNames := make([]string, 2)
	userNames = append(userNames, "Vignesh")
	userNames = append(userNames, "Tesla")
	fmt.Println(userNames)

}

func DynamicArrays() {
	var prices []float64 = []float64{10.99}
	prices = append(prices, 5.99)
	prices = prices[1:]
	prices[0] = 1
	fmt.Println(prices)

	discountPrices := []float64{101.99, 80.99, 20.59}
	prices = append(prices, discountPrices...)
	fmt.Println(prices)

	for index, value := range discountPrices {
		fmt.Println(index, value)
	}

}

func Maps() {
	websites := map[string]string{}
	websites["google"] = "https://google.com"
	websites["aws"] = "https://aws.com"
	fmt.Println(websites)
	fmt.Println(websites["azure"])

	delete(websites, "azure")
	fmt.Println(websites)

	courseRatings := make(floatMap, 3)
	courseRatings["english"] = 2.0
	courseRatings["mandarin"] = 7.8
	courseRatings.output()

	for key, value := range courseRatings {
		fmt.Println(key, value)
	}
}
