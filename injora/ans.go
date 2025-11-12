// Assignment: Conditional Statements in Go (Golang)
// Topics Covered: if, else if, else, nested if, short if, switch, fallthrough, switch without expression

// -----------------------------------------------------------
// Section A: Easy (Fundamentals)
// -----------------------------------------------------------

// 1. Write a program to check if a number is positive, negative, or zero using if statements.

// =========================================================
package main

import "fmt"

func posnevze(a int) string {
	if a > 0 {
		return "positive"
	} else if a < 0 {
		return "negative"
	} else {
		return "zero"
	}
}
func main() {
	var a int
	fmt.Scan(&a)
	fmt.Print(posnevze(a))
}

// =========================================================

// 2. Write a program to determine whether a person is eligible to vote based on age.

// =========================================================
package main

import "fmt"

func vote (a int)string{
    if a>17 {
        return "eligible"
    }  else {
        return "not eligible"
    }
}
func main(){
    var a int
    fmt.Scan(&a)
    fmt.Print(vote(a))
}
// =========================================================

// 3. Write a program that checks whether a given number is even or odd.

// =========================================================
package main

import "fmt"

func parity (a int)string{
    if a%2==0 {
        return "even"
    }  else {
        return "odd"
    }
}
func main(){
    var a int
    fmt.Scan(&a)
    fmt.Print(parity(a))
}
// =========================================================

// 4. Write a program that takes a student’s marks and prints “Pass” or “Fail” using an if-else condition.

// =========================================================
package main

import "fmt"

func passfail (a int)string{
    if a>60 {
        return "pass"
    }  else {
        return "fail"
    }
}
func main(){
    var a int
    fmt.Scan(&a)
    fmt.Print(passfail(a))
}
// =========================================================

// 5. Write a program to check whether a given year is a leap year or not.

// =========================================================
package main

import "fmt"

func leapyear (a int)string{
    if a%4 == 0 {
        return "Yes"
    }  else {
        return "No"
    }
}
func main(){
    var a int
    fmt.Scan(&a)
    fmt.Print(leapyear(a))
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
package main

import "fmt"

func leapyear (a int)string{
    if a>=90 {
        return "A"
    }  else if 75<=a && a<=89{
        return "B"
    } else if 60<=a && a<=74{
        return "C"
    } else {
        return "F"
    }
}
func main(){
    var a int
    fmt.Scan(&a)
    fmt.Print(leapyear(a))
}
// =========================================================

// 2. Write a program that takes two numbers and prints which one is greater using if-else.

// =========================================================
package main

import "fmt"

func greater (a ,b int)int{
    if a>b {
        return a
    } else {
        return b
    }
}
func main(){
    var a int
    var b int
    fmt.Scan(&a)
    fmt.Scan(&b)
    fmt.Print(greater(a,b)," is greater")
}
// =========================================================

// 3. Write a program that checks if a character is a vowel or consonant using a switch statement.

// =========================================================
package main

import "fmt"

func main(){
    var a string
    fmt.Scan(&a)
    switch a {
        case "a" :
        fmt.Println("Vowel")
        case "e" :
        fmt.Println("Vowel")
        case "i" :
        fmt.Println("Vowel")
        case "o" :
        fmt.Println("Vowel")
        case "u" :
        fmt.Println("Vowel")
        default:
        fmt.Println("Consonent")
    }
}
// =========================================================

// 4. Write a program that takes a day name (like Monday, Tuesday, etc.) and prints whether it’s a weekday or weekend using switch.

// =========================================================
package main

import "fmt"

func main(){
    var day string
    fmt.Scan(&day)
	switch day {
	case "Monday":
		fmt.Println("Start of the week")
	case "Friday":
		fmt.Println("Weekend loading...")
	case "Saturday", "Sunday":
		fmt.Println("It's the weekend!")
	default:
		fmt.Println("Midweek grind")
	}
}
// =========================================================

// 5. Write a program that uses a short if statement to check if a number is divisible by 3 and 5 at the same time.

// =========================================================
package main

import "fmt"

func div (a int)string{
    if a%3 ==0 && a%5==0 {
        return "Yes"
    } else {
        return "No"
    }
}
func main(){
    var a int
    fmt.Scan(&a)
    fmt.Println(div(a))
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
package main

import "fmt"

func auth (a,b,user,pass string)string{
    if  a==user && b==pass{
        return "Login Successful"
    } else if a==user && b!=pass{
        return "Invalid password"
    } else {
        return "User not found"
    }
}
func main(){
    var a string
    var b string
    user := "Injora"
    pass := "dev_club_president"
    fmt.Scan(&a)
    fmt.Scan(&b)
    fmt.Println(auth(a,b,user,pass))
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
package main

import "fmt"

func tri (a,b,c int)string{
    if  a==b && b==c && c==a{
        return "Equilateral"
    } else if a==b || b==c || c==a{
        return "Isoceles"
    } else {
        return "Scalene"
    }
}
func main(){
    var a int
    var b int
    var c int
    fmt.Scan(&a)
    fmt.Scan(&b)
    fmt.Scan(&c)
    fmt.Println(tri(a,b,c))
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
package main

import "fmt"

func main(){
    var a int
    fmt.Scan(&a)
    switch a {
        case 1 :
        fmt.Println("Start Game")
        case 2 :
        fmt.Println("Load Game")
        case 3 :
        fmt.Println("Exit")
        default :
        fmt.Println("Invalid Option")
    }
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
package main

import "fmt"

func main(){
    var a int
    fmt.Scan(&a)
    switch {
        case a<0 :
        fmt.Println("Freezing")
        case 0<=a && a<=15 :
        fmt.Println("Cold")
        case 16<=a && a<=30 :
        fmt.Println("Warm")
        default :
        fmt.Println("Hot")
    }
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
package main

import "fmt"

func elig (a int,b bool)string{
    if  a>=60 && b==true{
        return "Eligible for admission"
    } else if a<60 && b==true{
        return "Not eligible, less total percentage"
    } else if a>=60 && b==false {
        return "Not eligible, not passed in both matha nd science"
    } else {
        return "Not eligible, no conditions met"
    }
}
func main(){
    var a int
    var b bool
    fmt.Scan(&a)
    fmt.Scan(&b)

    fmt.Println(elig(a,b))
}
// =========================================================

// -----------------------------------------------------------
// Bonus (Challenge Zone)
// -----------------------------------------------------------

// 1. Write a program that takes three numbers as input and prints the largest using only conditional statements.

// Hint: Use nested if or multiple if-else conditions.

// =========================================================
package main

import "fmt"

func large (a,b,c int)int{
    if  a>b && a>c{
        return a
    } else if b>a && b>c{
        return b
    } else{
        return c
    }
}
func main(){
    var a int
    var b int
    var c int
    fmt.Scan(&a)
    fmt.Scan(&b)
    fmt.Scan(&c)
    fmt.Println(large(a,b,c))
}
// =========================================================

// 2. Create a small calculator using a switch statement:
//    - Ask the user to input two numbers and an operator (+, -, *, /).
//    - Perform the operation based on the chosen operator.
//    - Display the result.
//    - If the operator is invalid, print “Invalid operator”.

// =========================================================
package main

import "fmt"

func calc (a int,b int,c string)int{
    if  c=="+" {
        return a+b
    } else if c == "-" {
        return a-b
    } else if c=="/"{
        return a/b
    } else{
        return a*b
    }
}
func main(){
    var a int
    var b int
    var c string
    fmt.Scan(&a)
    fmt.Scan(&b)
    fmt.Scan(&c)
    fmt.Println(calc(a,b,c))
}
// =========================================================

// 3. Write a program that determines whether a given year is a century leap year or not.
//    - Ask the user to enter a year.
//    - A year is a century leap year if it is divisible by 400.
//    - A year is a normal leap year if it is divisible by 4 but not by 100.
//    - Otherwise, it is not a leap year.

// =========================================================
package main

import "fmt"

func cent (a int)string{
    if  a%400==0 {
        return "Century Leap"
    } else if  a%4==0 && a%100!= 0 {
        return "Leap Year"
    }else{
        return "None"
    }
}
func main(){
    var a int
    fmt.Scan(&a)
    fmt.Println(cent(a))
}
// =========================================================

// 4. Write a program that uses fallthrough in a switch to demonstrate cumulative conditions:
//    - Example: Enter a number and print messages for all cases up to that number using fallthrough.
//    - For example, if number = 2, output should print case 1 and case 2 messages.

// =========================================================
package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	switch n {
	case 1:
		fmt.Println("Case 1")
		fallthrough
	case 2:
		fmt.Println("Case 2")
		fallthrough
	case 3:
		fmt.Println("Case 3")
		fallthrough
	case 4:
		fmt.Println("Case 4")
		fallthrough
	case 5:
		fmt.Println("Case 5")
	default:
		fmt.Println("out of range")
	}
}

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
package main

import "fmt"

func main() {
	var mark int
	var attend int
	var back bool

	fmt.Scan(&mark)
	fmt.Scan(&attend)
	fmt.Scan(&back)

	if mark >= 85 && attend > 90 && back == false {
		fmt.Println("Scholarship Approved")
	} else {
		fmt.Println("Scholarship Denied")
	}
}

// =========================================================
