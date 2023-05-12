// Package dog allows us to more fully understand dogs.
package dog

// Years converts human years to dog years.
func Years(n int) int {
	if n < 0 {
		return 0
	}
	return n * 7
}

// YearsTwo converts human years to dog years.
func YearsTwo(n int) int {
	if n < 0 {
		return 0
	}

	count := 0
	for i := 0; i < n; i++ {
		count += 7
	}
	return count
}
