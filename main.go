// Package is the basic unity of code
// Packages are similar to libraries or modules in another languages
package main

// Import is the instruction to refer to another packages
import "fmt"

var globalVariable bool // this global-scoped variable is initialized with "false" value

// init function executes before main(), and is designed to
func init(){
globalVariable=true

println("check stdout! this routine runs first. Boolean global variable is now ",globalVariable)
}


func main() {
	const strConst = "string constant"

	const (
		pi            = 3.1415926535 // float64 constant
		answer uint64 = 42           // uint64 constant
		tab           = '\t' // use single quote to declare rune
		intVar   = 0b00001111 // int variable declared as binary literal
	)

	var strA = "string variable declared with 'var' keyword"
	var strB string // string variable declared and initialized with ""
	strC := "short form declaring and initializing by inference as string"

	// Declaring one or more variables in the same instruction block
	// The same can be done with "const" declarations
	var (
		strD string = `literal string containing "double quoted substring" and not subject to escape characters like \t or \n`
		intA  = 0                             // int by inference
		intB  uint8                           // declared as eight bits unsigned int, initialized with zero
		fltA        = 12.5                    // float64 type is standard for floating point values
		fltB        = float32(9) / float32(2) // here we're forcing float32 type
		int9By2        = 9 / 2 // Gotcha! Here the result is integer.
		runeA       = 'A'
		runeB rune  = 66
	)

	// This is the standard/embedded function printLn(), to print content to stdout
	println("There goes the variables to stdout, using fmt.Printf function ( Printf from fmt imported package )")

	fmt.Printf("pi is type %T and value %#v\n", pi, pi)

	fmt.Printf("answer is type %T and value %#v\n", answer, answer)

	fmt.Printf("tab is type %T and value %#v\n", tab, tab)

	fmt.Printf("intVar is type %T and value %#v\n", intVar, intVar)

	fmt.Printf("strA is type %T and value %#v and got %d bytes lenght\n", strA, strA, len(strA))

	fmt.Printf("strB is type %T and value %#v\n", strB, strB)

	fmt.Printf("strC is type %T and value %#v\n", strC, strC)

	fmt.Printf("strC is type %T and value %#v\n", strD, strD)

	fmt.Printf("intA is type %T and value %#v\n", intA, intA)

	fmt.Printf("intB is type %T and value %#v\n", intB, intB)

	fmt.Printf("fltA is type %T and value %#v\n", fltA, fltA)

	fmt.Printf("fltB is type %T and value %#v\n", fltB, fltB)

	fmt.Printf("int9By2 is type %T and value %#v\n", int9By2,int9By2)

	// printLn is a built-in routine
	println("Rune is an alias for type int32")

	fmt.Printf("runeA is type %T and value %c\n", runeA, runeA)

	fmt.Printf("runeB is type %T and value %c\n", runeB, runeB)
	
	// Anonymous functions are allowed in go, and you can attrib their bodies to a variable or execute directly when declaring
	closure:=func(msg string){fmt.Println(msg)}
	
	// defer instruction schedule a function to execute just before the current routine exits
	defer closure("i'm a deffered closure, and you have to see me by last")
	
	intResult:=func()int{return 42})()
}
