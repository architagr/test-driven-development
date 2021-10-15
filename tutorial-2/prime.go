package tutorial2

// CheckIfPrime this function helps in checking if the input number is prime or not
// we will be using method of contradiction i.e. first we will consider the input
// number is prime and try to validate if we are wrong
func CheckIfPrime(num int) bool {
	result := true

	if num < 2 {
		return false
	}
	for i := 2; i < num; i++ {
		if num%i == 0 {
			result = false
			break
		}
	}
	return result
}
