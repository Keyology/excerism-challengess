
package hamming
// Go Package

// imports go built in modules 
import(
	"errors"
)



// Function calculates the hamming Distance
func Distance(a, b string) (int, error) {

if len(a) != len(b){
// checks if parameter strings are equal length else throw error 
	return 0, errors.New("stings must be equal length")
}

// counter keep tracks of chacters that don't match
 counter := 0

 // loops over string a to get the index value
for index := range a {
	// check if strings dont't match then updates counter
	if a[index] != b[index]{
		counter++

	}
		

}
// return counter value
return counter, nil
	
}
