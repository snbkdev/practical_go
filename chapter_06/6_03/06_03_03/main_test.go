// Неудачные попытки вызвать сбой
package main

import "testing"

func TestFizzBuzzCodes(t *testing.T) {
	tests := []struct {
		name string
		input int64
		expected string
	}{
		{"fizz buzz test1", 37, "1, 2, Fizz, 4, Buzz, Fizz, 7, 8, Fizz, Buzz, 11, Fizz, 13, 14, Fizz Buzz, 16, 17, Fizz, 19, Buzz, Fizz, 22, 23, FIzz, Buzz, 26, Fizz, 28, 29, Fizz Buzz, 31, 32, Fizz, 34, Buzz, Fizz, 37"},
		{"fizz buzz test2", 5, "1, 2, Fizz, 4, Buzz"},
		{"fizz buzz test3", 12, "1, 2, Fizz, 4, Buzz, Fizz, 7, 8, Fizz, Buzz, 11, Fizz"},
	}

	for i := range tests {
		test := tests[i]
		t.Run(test.name, func(t *testing.T) {
			res := fizzbuzz(test.input)
			if res != test.expected {
				t.Fatalf("\ngot \n%s \nexpected \n%s", res, test.expected)
			}
		})
	}
}