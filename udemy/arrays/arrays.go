package main

import "fmt"

func main() {
	//1
	var hobbys = [3]string{"programming", "eating", "sport"}
	fmt.Println(hobbys)
	//2
	firstEl := hobbys[0]
	fmt.Println(firstEl)
	restOfHobbys := hobbys[1:3]
	fmt.Println(restOfHobbys)	
	//3
	thirdSlice1 := hobbys[:2]
	thirdSlice2 := hobbys[0:2]
	fmt.Println(thirdSlice1, thirdSlice2)
	//4
	//cap, len	
	fourthSlice := thirdSlice1[1:3]
	fmt.Println(fourthSlice)	
	//5		
	var goals = []string {"learnGo", "earnGold"}
	fmt.Println(goals)		
	//6
	goals[1] = "earnDollars"
	goals = append(goals, "beHappy")
	fmt.Println(goals)	
	//7
	type Product struct {
	Title string 
	ID    int
	Price float64
	}

	products := []Product{
		{"hair", 0, 9.02},
		{"test", 1, 1.02},
	}
	fmt.Println(products)
	
	products = append(products, Product{"new", 2, 1.552})
	// prices = append(prices, discountPrices...)
	fmt.Println(products)

}