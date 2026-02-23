// Неудачные попытки вызвать сбой
package main

import "testing"

func TestSummedRUneCodes(t *testing.T) {
	tests := []struct{
		name string
		input string
		expected int16
	}{
		{"test 1", "I am trying things", 1729},
		{"test 2", "doing my best to find a way to break this!", 3772},
		{"test 3", "even adding emojis or unicode doesn't break it, I'm sure I'm fine", 6065},
	}

	t.Run("my rune sum tests", func(t *testing.T) {
		for k := range tests {
			test := tests[k]
			t.Run(test.name, func(t *testing.T) {
				if got := summedRuneCodes(test.input); got != test.expected {
					t.Errorf("expected %d, got %d", test.expected, got)
				}
			})
		}
	})
}

// Выявление ошибок с помощью фаззинг-тестирования 
func FuzzSummedRuneCodes(f *testing.F) {
	tests := []string{
		"I am trying things",
		"doing my best to find a way to break this!",
		"even adding emojis or unicode doesn't break it, I'm sure I'm fine",
	}

	for t := range tests {
		f.Add(tests[t])
	}

	f.Fuzz(func(t *testing.T, seed string) {
		got := summedRuneCodes(seed)
		if got < 0 {
			t.Errorf("how did this happen ? somehow we got %d from string %s", got, seed)
		}
	})
}