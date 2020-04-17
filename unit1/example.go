// This is a single line comment
/*
This comment
is over
multiple lines
*/
package main

// listing any packages that we need to import
import (
	// We are using fmt so import it
	"fmt"
	/*
	   make sure we only import things we use
	   otherwise we will get a compiler error
	   You can try by uncommenting the following and trying to build
	*///"io/ioutil"
)

// The main function of the package does not take any values in our return any
func main() {
	// This will print "I am alive"
	fmt.Println("I am alive")
}
