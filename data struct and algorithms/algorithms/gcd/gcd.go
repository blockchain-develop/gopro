package gcd

func gcd(a int, b int) int {
	if a == 0 {
		return b
	}
	r := b % a
	return gcd(r, a)
}

func GCD(a int, b int) int {
	if a < b {
		return gcd(a, b)
	} else {
		return gcd(b, a)
	}
}
