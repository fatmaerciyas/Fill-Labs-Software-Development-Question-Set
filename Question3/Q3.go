package main

// fmt package provides the function to print anything
import (
	"fmt"
)

func main() {
	//I created it as interface{} type to be able to store different types of data.
	var array []interface{}
	array = append(array, "bb", 4, "aa", 4, 2, 3.2, "aa", 2, "aa", 3)

	fmt.Println(getMostRepeatedIntData(array)) 
}

func getMostRepeatedIntData(array []interface{}) interface{} {

	// this will hold the values in the array and the corresponding counter value
	 counter := make(map[interface{}]int) // map[KeyType]ValueType

	 // Increase the counter value corresponding to the element in each loop
	 for _, value := range array {
		counter[value]++
	}

	// Find the highest value in the counter map 
	maxCount := 0
	for _, value := range counter {
		if value > maxCount {
			maxCount = value
		}
	}

	//most repeated elements can be more than 1
	mostRepeatedElements := make([]interface{}, 0)

	// using the MaxCount value, find the corresponding key
	for key, value := range counter {
		if value == maxCount {
			mostRepeatedElements = append(mostRepeatedElements, key)
		}
	}

	return mostRepeatedElements
}

