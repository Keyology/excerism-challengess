package scrabble

import (
	"strings"
)

// store each group of letters in a array
//

// func Score(words string) int{

// 	capWords := strings.ToLower(words)

//  scrableWords := map[string]int{"a": 1, "e": 1, "i": 1, "o": 1, "u": 1, "l": 1, "n": 1, "r": 1, "s": 1, "t": 1, "d": 2, "g": 2, "b": 3, "c": 3, "m": 3, "p": 3, "f": 4, "h": 4, "v": 4, "w": 4, "y": 4, "k": 5, "j": 8, "x": 8, "q": 10, "z": 10}

// 	scoreTotal := 0
// 	pointer := &scoreTotal

// 	for i := 0; i < len(capWords); i++ {
		

// 		wordIndex := string([]rune(capWords)[i])

// 		for key, value := range onePoint {

// 			if key == wordIndex && value == 1 {
				
// 				*pointer++

// 			} else if key == wordIndex && value == 2 {
// 				*pointer += 2
// 			} else if key == wordIndex && value == 3 {
// 				*pointer += 3

// 			} else if key == wordIndex && value == 4 {
// 				*pointer += 4
// 			} else if key == wordIndex && value == 5 {
// 				*pointer += 5
// 			} else if key == wordIndex && value == 8 {
// 				*pointer += 8
// 			} else if key == wordIndex && value == 10 {
// 				*pointer += 10
// 			} else {
// 				*pointer = 0
// 			}

// 		}
		

// 	}
// 	return scoreTotal

// }


func Score(word string)int{

	// store the words and points as key value pairs 

scrableWords := map[rune]int{'a': 1,
'e': 1,
'i': 1,
'o': 1,
'u': 1,
'l': 1,
'n': 1,
'r': 1,
's': 1,
't': 1,
'd': 2,
'g': 2,
'b': 3,
'c': 3,
'm': 3,
'p': 3,
'f': 4,
'h': 4,
'v': 4,
'w': 4,
'y': 4,
'k': 5,
'j': 8,
'x': 8,
'q': 10,
'z': 10,
}

// variable to keep track of the score 
score := 0

//convert word to lower case
word = strings.ToLower(word)

// loop over words
for _,value := range word{
	
	score += scrableWords[value]

}

return score



}

