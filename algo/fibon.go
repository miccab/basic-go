package algo


func FibonRecur(n int) int {
	switch n {
	case 0:
		return 0
	case 1:
		return 1
	default:
		return FibonRecur(n-1) + FibonRecur(n-2)
	}
}


