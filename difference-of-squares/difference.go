package diffsquares

func SquareOfSum(number int)int{
	sum := 0
	
	for i := 0; i <= number; i++{
		sum += i

	}

	return sum * sum
}

func SumOfSquares(number int) (square int){
	for i := 0; i <= number; i++{
		square += i * i
		
	}
	return 
}


func Difference(number int)int {
	return SquareOfSum(number) - SumOfSquares(number)
}
