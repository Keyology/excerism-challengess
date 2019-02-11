// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package leap should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package leap



// IsLeapYear should have a comment documenting it.
func IsLeapYear(x int) bool {
	// Write some code here to pass the test suite.
	// Then remove all the stock comments.
	// They're here to help you get started but they only clutter a finished solution.
	// If you leave them in, reviewers may protest!

	// if leap year is divisble by 4 return true
	//else if leap year is divisble by 4 and 100 return false 
	// else if leap year divisble by 400 
	// else return false

	if ((x % 4 == 0) && !(x % 100 == 0)){
		return true 
	}else if  (( x % 100  == 0) && !(x % 400 == 0)){
		return false 
	} else if (x % 400  == 0){
		return true
	}else{
		return false

	}

	
	
}





