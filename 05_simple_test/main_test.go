package main

import "testing"

var primeTests = []struct {
	num      int
	expected bool
}{
	{0, false},
	{1, false},
	{2, true},
	{7, true},
	{8, false},
	{11, true},
	{12, false},
}

func Test_isPrime(t *testing.T) {

	for _, test := range primeTests {
		if output := isPrime(test.num); output != test.expected {
			t.Errorf("Output %t not equal to expected %t", output, test.expected)
		}
	}
}
