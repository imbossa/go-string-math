package go_string_math

import "testing"

func TestVerify(t *testing.T) {
	var tests = []struct {
		input    any
		expected string
	}{
		{1, "1"},
		{1.0, "1"},
		{-1.0, "-1"},
		{"343.32", "343.32"},
		{"1", "1"},
		{"-1.0", "-1"},
		{"1.000000", "1"},
		{"232142343435345452232131231231232342353453456455343432.323232", "232142343435345452232131231231232342353453456455343432.323232"},
	}

	for _, test := range tests {
		output, err := verify(test.input)

		if err != nil {
			t.Errorf("Test Failed: %v inputted, %v expected, received: %v", test.input, test.expected, err)
		}

		if output != test.expected {
			t.Errorf("Test Failed: %v inputted, %v expected, received: %v", test.input, test.expected, output)
		}
	}
}

func TestAdd(t *testing.T) {
	var tests = []struct {
		a        any
		b        any
		expected string
	}{
		{1, 1, "2"},
		{1.0, 1.0, "2"},
		{-1.0, 1.0, "0"},
		{"343.32", "343.32", "686.64"},
		{"1", "1", "2"},
		{"-1.0", "1.0", "0"},
		{"1.000000", "1.000000", "2"},
		{"232142343435345452232131231231232342353453456455343432.323232", "232142343435345452232131231231232342353453456455343432.323232", "464284686870690904464262462462464684706906912910686864.646464"},
	}

	for _, test := range tests {
		output, err := Add(test.a, test.b)

		if err != nil {
			t.Errorf("Test Failed: %v + %v inputted, %v expected, received: %v", test.a, test.b, test.expected, err)
		}

		if output != test.expected {
			t.Errorf("Test Failed: %v + %v inputted, %v expected, received: %v", test.a, test.b, test.expected, output)
		}
	}
}

func TestSubtract(t *testing.T) {
	var tests = []struct {
		a        any
		b        any
		expected string
	}{
		{1, 1, "0"},
		{1.0, 1.0, "0"},
		{-1.0, 1.0, "-2"},
		{"343.32", "343.32", "0"},
		{"1", "1", "0"},
		{"-1.0", "1.0", "-2"},
		{"1.000000", "1.000000", "0"},
		{"232142343435345452232131231231232342353453456455343432.323232", "232142343435345452232131231231232342353453456455343432.323232", "0"},
	}

	for _, test := range tests {
		output, err := Subtract(test.a, test.b)

		if err != nil {
			t.Errorf("Test Failed: %v - %v inputted, %v expected, received: %v", test.a, test.b, test.expected, err)
		}

		if output != test.expected {
			t.Errorf("Test Failed: %v - %v inputted, %v expected, received: %v", test.a, test.b, test.expected, output)
		}
	}
}
