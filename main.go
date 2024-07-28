package main 
// This is a special package name called main, that tells the compiler to look for the entry point function in this file.

import (
	"fmt"
	"unicode/utf8"
)
// Importing the fmt package, which is the standard package for formatting and printing.

func main(){
	fmt.Println("Hello World!")
	var intNum int = 32767

	// Integers
	intNum = intNum + 1
	fmt.Println(intNum)
	// Declaring a variable named intNum of type int. Other types include int8, int16, int32, int64, which specify the amount of bits used to store the number.
	// Int will default to 32 or 64 bits depending on your system capability
	var unIntNum uint = 32767
	fmt.Println(unIntNum)
	// Also have access to unsigned integers, which can store the same bit sizes as ints, unit8, unit16, unit32, and unit64 but can store twice as large values as ints (0,255) where as ints (-128,127).

	// Floats
	var floatNum float64 = 32767.9
	// Go doesn't have a built in float type, but it does have float32 and float64 types. Float64 is the most precise. 
	fmt.Println(floatNum)

	// Go has similar arthmetic operations like other languages.
	var floatNum32 float32 = 10.1
	var intNum32 int32 = 2
	// var result float32 = floatNum32 + intNum32 
	// Cannot add different types together. Would need to cast it
	var result float32 = floatNum32 + float32(intNum32)
	fmt.Println(result)

	var intNum1 int = 3
	var intNum2 int = 2
	fmt.Println(intNum1/intNum2)
	// This will return in a rounded down value.
	fmt.Println(intNum1%intNum2)
	//This will return the remainder of the division.

	// Strings
	var myString string = "Hello World"
	var myString2 string = `Hello
	World`
	var myString3 string = "Hello" + "  " + "World"
	fmt.Println(myString)
	fmt.Println(myString2)
	fmt.Println(myString3)
	fmt.Println(len("test"))
	// This doesn't return the length of the string, but actually the number of bytes in the string.
	// Anything outside of the ASCII range will return more than one byte as Go uses UTF-8 encoding.
	fmt.Println(len("A"))
	// Above gives you 1
	fmt.Println(len("😀"))
	// Above gives you 4	
	fmt.Println(utf8.RuneCountInString("😀"))
	// Above gives you 1

	// Runes
	// Runes are another type in go, which represent a single unicode character.
	var myRune rune = 'a'
	fmt.Println(myRune)
	// Returns the unicode value, which is 97 in this case.

	// Booleans
	var myBoolean bool = true
	fmt.Println(myBoolean)

	// Default values
	// The default value for a variable int, int32, int64, uint, uint32, uint64, float32, float64 is 0.
	// The default value for a string is an empty string.
	// The default value for a boolean is false.

	var myVar = "text" 
	// Type is inferred
	myVar2 := "text"
	// Shorthand for declaring a variable
	fmt.Println(myVar)
	fmt.Println(myVar2)

	var1, var2 := 1, 2
	// Fancy way to declare multiple variables at once
	fmt.Println(var1)
	fmt.Println(var2)

	// Better to declare type when unknown
	// E.g., myVar := foo()

	// Constants
	// Above rules apply to constants but values can't be changed
	// They need to declared with a value
	const myConst string = "const value"
	fmt.Println(myConst)



}