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

## Question 4: User Management System
User management system, which is the last question of the Quiz to join the Fill Labs team. The system basically works based on CRUD operations. When the program runs, the person running it can see all users in the database and add new users. Can change the information of an existing user and delete the user.This program has been developed within one week, with Go on the backend and TypeScript with React & Next.js on the frontend. The goal was to be able to create a REST API in a programming language unfamiliar to me within this one-week timeframe.
### Install and Run 
You can easily download the project from GitHub. After downloading the project, open two terminals on your computer, one for the client and one for the server. Then, to run the server side, navigate to the server folder in your terminal and execute the command ``` go run mainGin.go ``` This command will run the mainGin.go file. SQLite database is used in the project, and if you want to see the changes in the data on the server side, you can also open the file with the .db extension under the database folder.
To run the client, navigate to the user-management project in your terminal. Since the React project is created with Vite, you need to execute the command ``` npm run dev ``` After running the command, click on the link displayed in the terminal and enjoy the project :)


If you're curious about how the project looks before running it, you can check it out here. --> https://drive.google.com/drive/folders/19ML3PwGrZSk5FhKAOwV0OoR-Yy6D2rfS?usp=sharing 


![download](https://github.com/fatmaerciyas/Fill-Labs-Software-Development-Question-Set/assets/55760365/a5bdf3b8-b520-46d8-b6f8-df069244fe3b)
