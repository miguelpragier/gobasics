package mypack

import "fmt"

var observeThisVariable = "INITIAL VALUE SET DURING DECLARATION"

// Init gonna be executed once this package is imported
func init() {
	fmt.Println("Executing mypack.init() routine")

	notExportedRoutine()

	observeThisVariable = "SECOND VALUE SET DURING init() routine"
}

// notExportedRoutine can't be seen from outside this package (mypack)
func notExportedRoutine() {
	fmt.Printf("here's the current value of observeThisVariable: %s\n", observeThisVariable)
}

// ExportedRoutine can be executed by other packages like this: mypack.ExportedRoutine()
func ExportedRoutine() {
	fmt.Printf("here's how is observeThisVariable: %s\n", observeThisVariable)
}
