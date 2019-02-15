package raindrops 

import "strconv"

func Convert(x int) string{
	rainDrops := ""


	if (x%3 ==0){
		rainDrops += "Pling"
	}

	if (x% 5 == 0){
		rainDrops += "Plang"
	}

	if ( x % 7 == 0){

		rainDrops += "Plong"

	}

	if rainDrops == ""{
		rainDrops += strconv.Itoa(x)
		
	}

	return rainDrops
}