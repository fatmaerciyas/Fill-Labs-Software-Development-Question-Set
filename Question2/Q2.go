package main

// fmt package provides the function to print anything
import (
	"fmt"
	"sort"
)

//I defined global slice. Because I don't want the slice to be created again every time I call it in the divideNumber() function.
var slice = make([]int,0)

func main() {

	var number int
	fmt.Println("Please enter a number")

	// scanning the input by the user
	fmt.Scanln(&number)

	fmt.Println(divideNumber(number))
}

//Find the even positive divisors of the given number except 0
func divideNumber(number int) []int{
	
	slice = append(slice, number)
	output := number / 2 

	//if the number is an even number other than 0, call function 
	if (output != 0 && output %2 == 0){
		slice = divideNumber(output)
	} 

	//order by ascending
	sort.Sort(sort.IntSlice(slice))

	//It can also be sorted like this
	//sort.Slice(slice, func(i, j int) bool {
	// 	return slice[i] < slice[j]
	// })

	return slice
}

