package fibonacci

import "testing"

func TestParseArgs(t *testing.T) {
	_, _, err := ParseArgs("/[1;5]", "http")
	if err != nil{
		t.Error("Expected correct parsing")
	}
	_, _, err = ParseArgs("/[1;         5]", "http")
	if err != nil{
		t.Error("Expected correct parsing")
	}

	_, _, err = ParseArgs("/[1;5;123]", "http")
	if err == nil{
		t.Error("Expected error")
	}

	_, _, err = ParseArgs("/[1;5", "http")
	if err == nil{
		t.Error("Expected error")
	}

	_, _, err = ParseArgs("/1;5]", "http")
	if err == nil{
		t.Error("Expected error")
	}

	_, _, err = ParseArgs("/[1 5]", "http")
	if err == nil{
		t.Error("Expected error")
	}

}


func TestFibo(t *testing.T) {
	var (
		val []int
		sum int
	)
	_, err := Fibo(0, 5)
	if err == nil{
		t.Error("Expected error")
	}

	_, err = Fibo(5, 1)
	if err == nil{
		t.Error("Expected error")
	}

	val, err = Fibo(1, 5)
	if err != nil{
		t.Error("Expected correct work")
	}
	for _, i := range val{
		sum += i
	}

	if sum != 12{
		t.Error("Incorrect fibonacci values")
	}

}

func TestGetDataFromRedis(t *testing.T) {
	var ind int
	res := make([]int, 100)
	_, err := GetDataFromRedis(100, res)

	if err == nil{
		t.Error("Expected error")
	}

	ind, err = GetDataFromRedis(5, res)
	if ind != 0 && err != nil{
		t.Error("Must get all values from Redis")
	}

	ind, err = GetDataFromRedis(6, res)
	if ind != 5 && err == nil{
		t.Error("Must get all values from Redis")
	}


}
