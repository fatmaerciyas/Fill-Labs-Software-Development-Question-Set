# Software Development Questions 

This repository consists of three questions and their answers asked to learn the Go programming language from scratch. To run the programs, you need to navigate to the folder of each question in your terminal and then execute your ``` .go  ``` file. After positioning, you can run your file as ``` go run filename.go  ```. For example, for the first question, you need to run the ``` Q1.go  ``` file.

## Question 1:
Write a function that sorts a bunch of words by the number of character "a"s within the word (decreasing order). If some words contain the same amount of character "a"s, then you need to sort those words by their lengths.
Answer is:
After running the program, enter each element you want to add to the array one by one. Each time you input an element, the program will provide you with the sorted output. In this program, I used the radix sort algorithm for sorting. To exit the program, simply press the "x" key.

## Question 2:
Write a recursive function that takes one integer parameter. Please bear in mind that finding the algorithm needed to generate the output below is the main point of the question. If the input is 9, the output is 2.4.9.
Answer is:
I assumed the algorithm for this question that the input number is halved until the end is 0, and I used a recursive function to send the output back to the function at each stage.

## Question 3:
Write a function that takes one parameter as an array/list. Find the most repeated data within a given array.
Answer is:
In this question, I used the interface {} data structure to hold an array consisting of different data types. Created a loop inside the array, and using a counter. Found which data is repeated the most. If two pieces of data have the same maximum repetition count, printed both.

![download](https://github.com/fatmaerciyas/Fill-Labs-Software-Development-Question-Set/assets/55760365/a5bdf3b8-b520-46d8-b6f8-df069244fe3b)
