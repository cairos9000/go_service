package fibonacci

import "testing"

func TestParseArgs(t *testing.T) {
	_, _, err := ParseArgs("/[1;5]", "http")
	if err != nil {
		t.Error("Expected correct parsing")
	}
	_, _, err = ParseArgs("/[1;         5]", "http")
	if err != nil {
		t.Error("Expected correct parsing")
	}

	_, _, err = ParseArgs("/[1;5;123]", "http")
	if err == nil {
		t.Error("Expected error")
	}

	_, _, err = ParseArgs("/[1;5", "http")
	if err == nil {
		t.Error("Expected error")
	}

	_, _, err = ParseArgs("/1;5]", "http")
	if err == nil {
		t.Error("Expected error")
	}

	_, _, err = ParseArgs("/[1 5]", "http")
	if err == nil {
		t.Error("Expected error")
	}

}

func TestFibo(t *testing.T) {
	var (
		val []int
		sum int
	)
	_, err := Fibo(0, 5)
	if err == nil {
		t.Error("Expected error")
	}

	_, err = Fibo(5, 1)
	if err == nil {
		t.Error("Expected error")
	}

	val, err = Fibo(1, 5)
	if err != nil {
		t.Error("Expected correct work")
	}
	for _, i := range val {
		sum += i
	}

	if sum != 12 {
		t.Error("Incorrect fibonacci values")
	}

}
