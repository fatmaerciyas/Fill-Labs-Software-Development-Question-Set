package main

// fmt package provides the function to print anything
import (
	"fmt"
)


func main() {

   // declaring the variable using the var keyword
   var flag = true
   slice := make([]string,1)

   for flag == true {
	var inputFromUser string
	fmt.Println("Please enter the elements to be added to the array. You can press the X key to exit the program.")
	
	// scanning the input by the user
	fmt.Scanln(&inputFromUser)

	if inputFromUser == "X" || inputFromUser == "x"{
		flag = false
	} else{

		slice = append(slice, inputFromUser)

		// sort.Strings(slice) can also be used but I will use the redix algorithm
		radixSort(slice)

		 fmt.Println(slice);
	   }
	}

}


// Function for counting sort
func countingSort(slice []string, index int) {
	n := len(slice)
	output := make([]string, n)
	count := make([]int, 256)
 
    // Store count of occurrences in count[]
	for i := 0; i < n; i++ {
	   char := CharAtIndex(slice[i], index)
	   count[char]++
	}
 
	// Change count[i] so that count[i] now contains actual position of this digit in output[]
	for i := 1; i < 256; i++ {
	   count[i] += count[i-1]
	}
 
	// Build the output array
	for i := n - 1; i >= 0; i-- {
	   char := CharAtIndex(slice[i], index)
	   output[count[char]-1] = slice[i]
	   count[char]--
	}
 
	// Sorted alphabetically. Now it will be sorted by character "a"

	 // Copy the output array to slice[],
    // so that slice[] now contains sorted
    // numbers according to current digit
	for i := 0; i < n; i++ {
		slice[i] = output[i]
	 }

}



func sortByCharacter(output []string) {
	var n = len(output)
	
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			countA := countLetter(output[j], 'a')
			countB := countLetter(output[j+1], 'a')
			
			//if there is no "a" letter in element sort by length
			if countA < countB || (countA == countB && len(output[j]) < len(output[j+1])) {
				output[j], output[j+1] = output[j+1], output[j]
			}
		}
	}
}



//Count letter "a"
func countLetter(word string, letter byte) int {
	count := 0
	for i := 0; i < len(word); i++ {
		if word[i] == letter {
			count++
		}
	}
	return count
}


// Get the ASCII value of the character at the specified index in the input string. 
func CharAtIndex(str string, index int) int {
	if index < len(str) {
	   return int(str[index])
	}
	return 0
 }

 
  // Sorts an array of strings using radix sort
 func radixSort(slice []string) {
	maxLen := MaxLen(slice)
 
	for i := maxLen - 1; i >= 0; i-- {
		sortByCharacter(slice)

	}
 }
 
// Finding the element with max length
func MaxLen(slice []string) int {
	maxLen := 0
	for i, str := range slice {
		countingSort(slice, i)
	   if len(str) > maxLen {
		  maxLen = len(str)
		  
	   }
	 
	}
	return maxLen
	
 }

/*This code can be written more easily and quickly using ready-made functions and methods. But I chose to use algorithms because the tests measure our ability to use algorithms.
Simpler code example:*/


//  input := []string{"aaaasd", "a", "aab", "aaabcd", "ef", "cssssssd", "fdz", "kf", "zc", "lklklklklklklklkl", "l"}
// 
/*we have a slice of strings and we want to sort them in alphabetical order. */

// sort.Slice(in, func(i, j int) bool {
//     s1, s2 := input[i], input[j]

// 	/*strings.Count(text that we want to search anything in it, the word or letter ) */

//     count1, count2 := strings.Count(s1, "a"), strings.Count(s2, "a")
//     if count1 != count2 {
//         return count1 > count2
//     }
// 
/* 	This code snippet checks if the count of the letter "a" is different between two strings. If the counts are different, it returns true if the count of 'a' in s1 is greater than the count of 'a' in s2, and false otherwise.
//  If the counts of 'a' are the same, it proceeds to the original logic and checks if the number of Unicode characters in s1 is greater than the number of Unicode characters in s2.*/

//     return utf8.RuneCountInString(s1) > utf8.RuneCountInString(s2)
// })

