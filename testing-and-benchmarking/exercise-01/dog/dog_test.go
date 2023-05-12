package dog_test

import (
	"fmt"
	"testing"

	"github.com/jrdelarosa8/go-notes/testing-and-benchmarking/exercise-01/dog"
)

type testCase struct {
	humanYears int
	dogYears   int
}

var testCases = []testCase{
	{0, 0},
	{1, 7},
	{7, 49},
	{-1, 0},
}

func TestYears(t *testing.T) {
	for _, c := range testCases {
		dogAge := dog.Years(c.humanYears)
		if dogAge != c.dogYears {
			t.Errorf("Got %v, wanted %v", dogAge, c.dogYears)
		}
	}
}

func TestYearsTwo(t *testing.T) {
	for _, c := range testCases {
		dogAge := dog.YearsTwo(c.humanYears)
		if dogAge != c.dogYears {
			t.Errorf("Got %v, wanted %v", dogAge, c.dogYears)
		}
	}
}

func ExampleYears() {
	fmt.Println(dog.Years(10))
	// Output:
	// 70
}

func ExampleYearsTwo() {
	fmt.Println(dog.YearsTwo(7))
	// Output:
	// 49
}
func BenchmarkYears(b *testing.B) {
	for i := 0; i < b.N; i++ {
		dog.Years(10)
	}
}

func BenchmarkYearsTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		dog.YearsTwo(10)
	}
}
