// Assignment: Conditional Statements in Go (Golang)
// Topics Covered: if, else if, else, nested if, short if, switch, fallthrough, switch without expression

// -----------------------------------------------------------
// Section A: Easy (Fundamentals)
// -----------------------------------------------------------
package main

import "fmt"
// 1. Write a program to check if a number is positive, negative, or zero using if statements.

// =========================================================
func main() {
	var a int
	fmt.Scan(&a)
	if(a>0){
		fmt.Println("Positive")
	}else if(a<0){
		fmt.Println("Negative")
	}else{
		fmt.Println("Zero")
	}




// =========================================================

// 2. Write a program to determine whether a person is eligible to vote based on age.

// =========================================================
	var age int
	fmt.Scan(&age)
	if(age>=18){
		fmt.Println("Eligible")
	}else{
		fmt.Println("Not Eligible")
	}



// =========================================================


// 3. Write a program that checks whether a given number is even or odd.

// =========================================================
	var num int
	fmt.Scan(&num)
	if(num%2==0){
		fmt.Println("Even")
	}else{
		fmt.Println("Odd")
	}



// =========================================================


// 4. Write a program that takes a student’s marks and prints “Pass” or “Fail” using an if-else condition.

// =========================================================
	var marks int
	fmt.Scan(&marks)
	if(marks>=40){
		fmt.Println("Pass")
	}else{
		fmt.Println("Fail")
	}



// =========================================================

// 5. Write a program to check whether a given year is a leap year or not.

// =========================================================
	var year int
	fmt.Scan(&year)
	if (year%4==0 && year%100!=0) || (year%400==0) {
		fmt.Println("Leap Year")
	}else{
		fmt.Println("Not a leap year")
	}



// =========================================================


// -----------------------------------------------------------
// Section B: Medium (Logic Building)
// -----------------------------------------------------------

// 1. Write a program to calculate grade based on marks using if-else if conditions.
//    90+ → Grade A
//    75–89 → Grade B
//    60–74 → Grade C
//    Below 60 → Grade F

// =========================================================
    var score int
	fmt.Scan(&score)
	if(score>=90){
		fmt.Println("Grade A")
	}else if(score >= 75){
		fmt.Println("Grade B")
	}else if(score >= 60){
		fmt.Println("Grade C")

	}else{
		fmt.Println("Grade F")
	}



// =========================================================

// 2. Write a program that takes two numbers and prints which one is greater using if-else.

// =========================================================
	var c,d int
	fmt.Scan(&c,&d)
	if(c>d){
		fmt.Printf("%d is greater",c)
	}else{
		fmt.Printf("%d is greater",d)
	}




// =========================================================

// 3. Write a program that checks if a character is a vowel or consonant using a switch statement.

// =========================================================
   var character string
   fmt.Scan(&character)
   if((character == "a" )|| (character=="e") || (character=="i") || (character=="o") || (character=="u")){
			fmt.Println("Vowel")
   }else{
	fmt.Println("Consonant")
   }




// =========================================================

// 4. Write a program that takes a day name (like Monday, Tuesday, etc.) and prints whether it’s a weekday or weekend using switch.

// =========================================================
   var day string
   fmt.Scan(&day)
   switch day{
   case "Monday":
	fmt.Println("Weekday")
   case "Tuesday":
	fmt.Println("Weekday")
   case "Wednesday":
	fmt.Println("Weekday")
   case "Thursday":
	fmt.Println("Weekday")
   case "Friday":
	fmt.Println("Weekday")
   case "Saturday":
	fmt.Println("Weekend")
   case "Sunday":
	fmt.Println("Weekend")
   default:
	fmt.Println("Bro Wrong Input")
   
   }





// =========================================================

// 5. Write a program that uses a short if statement to check if a number is divisible by 3 and 5 at the same time.

// =========================================================
var digit int
fmt.Scan(&digit)
if(digit%3==0 && digit%5==0){
	fmt.Println("Yessss !!!!")
}else{
	fmt.Println("Nooooo !!!!")
}



// =========================================================


// // -----------------------------------------------------------
// Section C: Hard (Real-World Logic & Nesting)
// -----------------------------------------------------------

// Note: Use the following syntax to take input from the user in Go:
//
// var variableName datatype
// fmt.Print("Enter something: ")
// fmt.Scan(&variableName)
//
// Example:
// var age int
// fmt.Print("Enter your age: ")
// fmt.Scan(&age)


// =========================================================

// 1. Write a program to simulate a login system:
//    - Ask the user to enter a username and password.
//    - If both match predefined values, print “Login successful”.
//    - If username is correct but password is wrong, print “Invalid password”.
//    - Otherwise, print “User not found”.

// Hint: You can compare input strings using the == operator.


// =========================================================
 var user,pass string
 fmt.Scan(&user,&pass)
 if(user=="ThorOdinson" && pass=="Point Break"){
	fmt.Println("Login Successful")
 }else if(user=="ThorOdinson"){
	fmt.Println("Invalid Password")
 }else{
	fmt.Println("User not found")
 }



// =========================================================

// 2. Write a program that checks whether a given triangle is:
//    - Equilateral (all sides equal)
//    - Isosceles (two sides equal)
//    - Scalene (all sides different)
//
//    - Ask the user to enter the lengths of all three sides.
//    - Use nested if statements to determine the triangle type.

// =========================================================
var s1,s2,s3 int
fmt.Scan(&s1,&s2,&s3)
if((s1==s2) && (s2==s3) && (s3==s1)){
	fmt.Println("Equilateral Triangle")
}else if((s1==s2) || (s2==s3) || (s3==s1)){
	fmt.Println("Isosceles Triangle")
}else{
	fmt.Println("Scalene")
}



// =========================================================

// 3. Write a program that simulates a menu system using a switch statement:
//    - Display a simple menu:
//         1 → “Start Game”
//         2 → “Load Game”
//         3 → “Exit”
//    - Ask the user to enter their choice (1, 2, or 3).
//    - Use a switch statement to print the corresponding message.
//    - If input doesn’t match any option, print “Invalid option”.

// =========================================================
fmt.Println("1 -> Start Game")
fmt.Println("2 -> Load Game")
fmt.Println("3 -> Exit")
var choice int
fmt.Scan(choice)
switch choice{
case 1:
	fmt.Println("Start Game")
case 2:
	fmt.Println("Load Game")
case 3:
	fmt.Println("Exit")
}




// =========================================================

// 4. Write a program that uses a switch without an expression to classify temperature:
//    - Ask the user to input the current temperature (integer).
//    - Based on the value, print one of the following:
//         Below 0 → “Freezing”
//         0–15 → “Cold”
//         16–30 → “Warm”
//         Above 30 → “Hot”

// Hint: Use logical operators (>, <, <=, >=) inside switch cases.

// =========================================================
var temp int
fmt.Scan(&temp)
if(temp<0){
	fmt.Println("Freezing")
}else if(temp<=15){
	fmt.Println("Cold")
}else if(temp<=30){
	fmt.Println("Warm")
}else{
	fmt.Println("Hot")
}



// =========================================================

// 5. Write a program to check admission eligibility for a student based on nested conditions:
//    - Ask the user to input total marks (in percentage).
//    - Ask if the student passed Math and Science (true/false).
//    - Conditions:
//         Minimum marks: 60%
//         Must have passed both Math and Science
//    - If both conditions are true, print “Eligible for admission”.
//    - Otherwise, print the specific reason for rejection.

// =========================================================
var tot int

	fmt.Scan(&tot)
	fmt.Println("Math:")
	var Math bool
	fmt.Scan(&Math)
	fmt.Println("Science:")
	var Science bool
	fmt.Scan(&Science)
	if tot >= 60 {
		if Math == true && Science == true {
			fmt.Println("Eligible")
		} else if Math == false && Science == true {
			fmt.Println("Ineligible due to Math")
		} else if Math == true && Science == false {
			fmt.Println("Ineligible due to Science")
		} else {
			fmt.Println("Ineligible due to Math and Science")
		} 
		}


// =========================================================


// -----------------------------------------------------------
// Bonus (Challenge Zone)
// -----------------------------------------------------------

// 1. Write a program that takes three numbers as input and prints the largest using only conditional statements.

// Hint: Use nested if or multiple if-else conditions.


// =========================================================
var num1,num2,num3 int
fmt.Scan(&num1,&num2,&num3)
if(num1>num2 && num1>num3){
	fmt.Printf("%d bigger\n",num1)
}else if(num2>num1 && num2>num3){
	fmt.Printf("%d bigger\n",num2)
}else{
	fmt.Printf("%d bigger\n",num3)
}



// =========================================================

// 2. Create a small calculator using a switch statement:
//    - Ask the user to input two numbers and an operator (+, -, *, /).
//    - Perform the operation based on the chosen operator.
//    - Display the result.
//    - If the operator is invalid, print “Invalid operator”.

// =========================================================




// =========================================================

// 3. Write a program that determines whether a given year is a century leap year or not.
//    - Ask the user to enter a year.
//    - A year is a century leap year if it is divisible by 400.
//    - A year is a normal leap year if it is divisible by 4 but not by 100.
//    - Otherwise, it is not a leap year.

// =========================================================




// =========================================================

// 4. Write a program that uses fallthrough in a switch to demonstrate cumulative conditions:
//    - Example: Enter a number and print messages for all cases up to that number using fallthrough.
//    - For example, if number = 2, output should print case 1 and case 2 messages.

// =========================================================




// =========================================================

// 5. Write a program that checks if a student qualifies for a scholarship:
//    - Ask for the student’s marks (percentage).
//    - Ask for attendance percentage.
//    - Ask if the student has any backlogs (true/false).
//    - Conditions:
//         Must have 85% or higher marks
//         Attendance above 90%
//         No backlog
//    - If all conditions are true, print “Scholarship Approved”.
//    - Otherwise, print “Scholarship Denied”.

// =========================================================




// =========================================================
}