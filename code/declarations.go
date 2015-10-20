// Variable
var x interface{} 	// x is nil and has static type interface{}
// Duck types variable
var x = interface{} 	// static type of x is inferred from the RHS

// Function
func F(name Type, …)(Type, …){ … }
func Printf(format string, params ...interface{}) // variadic function

// Type
type Sample struct { }

// Interface
type Sample interface { … }
