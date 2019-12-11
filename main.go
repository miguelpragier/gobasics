// Package is the basic unity of code
// Packages are similar to libraries or modules in another languages
package main

// Import is the instruction to refer to another packages
import (
	"fmt"
	"github.com/miguelpragier/gobasics/mypack"
)

var globalVariable bool // this global-scoped variable is initialized with "false" value

// init function executes before main(), and is designed to
func init() {
	globalVariable = true

	// printLn is a built-in routine
	println("check stdout! this routine runs first. Boolean global variable is now ", globalVariable)
}

// variablesAndSimpleTypes demonstrates all the possible ways to declare and initialize
// variables and constants of simple ( not struct ) types
func variablesAndSimpleTypes() {
	// Constants should be declared and initialized at once

	// Here goes a <uint64> type coertion.
	// in this case, if you don't specify type, the compiler would define as <int>
	const answer uint64 = 42 // uint64 constant

	// Automatic inference
	const (
		pi       = 3.1415926535 // float64 constant
		tab      = '\t'         // use single quote to declare rune
		strConst = "string constant"
		// The lines above only executes from GO 1.13 ( and later )
		intVar             = 0b00001111 // int variable declared as binary literal
		integerHexadecimal = 0x9a
	)

	var strA = "string variable declared with 'var' keyword"
	var strB string // string variable declared and initialized with ""
	strC := strConst + "short form declaring and initializing by inference"

	// Declaring one or more variables in the same instruction block
	// The same can be done with "const" declarations
	var (
		strD                             string = `literal string containing "double quoted substring" and not subject to escape characters like \t or \n`
		intA                                    = 0                       // int by inference
		intB                             uint8                            // declared as eight bits unsigned int, initialized with zero
		fltA                                    = 12.5                    // float64 type is standard for floating point values
		fltB                                    = float32(9) / float32(2) // here we're forcing float32 type
		int9By2                                 = 9 / 2                   // Gotcha! Here the result is integer.
		runeA                                   = 'A'
		runeB                            rune   = 66
		octalInt                                = 0o660
		hexadecimalFloatingPoint                = 0x123p8
		laterVariableUsingAPreviousValue        = strD
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

	fmt.Printf("int9By2 is type %T and value %#v\n", int9By2, int9By2)

	fmt.Printf("integerHexadecimal is type %T and value %d\n", integerHexadecimal, integerHexadecimal)

	fmt.Printf("octalInt is type %T and value %d\n", octalInt, octalInt)

	fmt.Printf("hexadecimalFloatingPoint is type %T and value %d\n", hexadecimalFloatingPoint, hexadecimalFloatingPoint)

	fmt.Printf("laterVariableUsingAPreviousValue is type %T and value %#v\n", laterVariableUsingAPreviousValue, laterVariableUsingAPreviousValue)

	// printLn is a built-in routine
	println("Rune is an alias for type int32")

	fmt.Printf("runeA is type %T and value %c\n", runeA, runeA)

	fmt.Printf("runeB is type %T and value %c\n", runeB, runeB)
}

func goStructs() {
	// Type Aliasing. We are not creating our type, but renaming an existing one.
	type trueOrFalse bool

	// A simple struct using the custom type/alias above
	type nestedType struct {
		float64Field float64
		booleanField trueOrFalse
	}

	// A struct using the simple struct above
	type myStruct struct {
		intField    int
		stringField string
		nt          nestedType
	}

	// Aliasing our custom type
	type myStructTypeAlias myStruct

	// We CANNOT use structs for constants.
	// Structs initializers are not considerated valid constant values
	//const complexConst = nestedType{
	//	float64Field: 0,
	//	booleanField: false,
	//}

	var (
		// Declaring and initializing a simple aliased type
		tf = trueOrFalse(true)

		// declaring variable of aliased type nestedType, which is a valid type for variables
		nestedTypeVar nestedType

		myStructVar = myStruct{
			intField:    15,
			stringField: "show me the code",
			nt: nestedType{
				float64Field: 161718.1920,
				booleanField: tf,
			},
		}

		anotherMyStruct = myStructTypeAlias{intField: 15,
			stringField: "show me the code",
			nt: nestedType{
				float64Field: 161718.1920,
				booleanField: tf,
			},
		}
	)

	var newStruct struct {
		X byte
		Y uint8
	}

	newStruct.X = 255

	newStruct.Y = newStruct.X

	anotherStruct := struct {
		SingleField string
	}{"anonymous struct declared and initialized"}

	var structArray = []struct {
		s string
		i int
	}{
		{s: "abc", i: 1},
		{s: "def", i: 2},
		{s: "ghi", i: 3}, // See this comma? It's mandatory when the block/curly brackets/mustache closes on another line
	}

	type structA struct {
		x int
	}

	type structB struct {
		y structA
	}

	var nestedStructs = struct {
		z structB
	}{z: structB{y: structA{0}}} // Notice that here we entirely ignored commas, once the entire initializator in inline

	nestedStructs.z.y.x = 15

	fmt.Printf("tf is type %T and value %#v\n", tf, tf)

	fmt.Printf("nestedTypeVar is type %T and value %#v\n", nestedTypeVar, nestedTypeVar)

	fmt.Printf("myStructVar is type %T and value %#v\n", myStructVar, myStructVar)

	fmt.Printf("anotherMyStruct is type %T and value %#v\n", anotherMyStruct, anotherMyStruct)

	fmt.Printf("anotherStruct is type %T and value %#v\n", anotherStruct, anotherStruct)

	fmt.Printf("structArray is type %T and value %#v\n", structArray, structArray)

	fmt.Printf("nestedStructs is type %T and value %#v\n", nestedStructs, nestedStructs)
}

func deferredRoutine() {
	fmt.Println("this is a deferred routine. Was scheduled by GO compiler to run after the last synchronous instruction end")
	fmt.Println("it's not the same as OOP destructor, but it's often used the same way of, to close/dispose allocated resources")
}

func closureRoutines() {
	// Anonymous functions are allowed in go, and you can attrib their bodies to a variable or execute directly when declaring
	closure := func(msg string) { fmt.Println(msg) }

	// defer instruction schedule a function to execute just before the current routine exits
	defer closure("This is the deferred anonymous function, and you have to see me by last")

	// this closure is declared and immediately used; the key are the trailing parenthesis
	intResult := func() int { return 42 }()

	fmt.Printf("This closure returned %d\n", intResult)
}

func packagedRoutine() {
	// Now we're using our own package
	mypack.ExportedRoutine()
}

func main() {
	variablesAndSimpleTypes()

	goStructs()

	defer deferredRoutine()

	closureRoutines()

	packagedRoutine()
}
