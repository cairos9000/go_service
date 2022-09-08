package redis

import "testing"

func TestGetDataFromRedis(t *testing.T) {
	var ind int
	res := make([]int, 100)
	_, err := GetDataFromRedis(100, res)

	if err == nil {
		t.Error("Expected error")
	}

	ind, err = GetDataFromRedis(5, res)
	if ind != 0 && err != nil {
		t.Error("Must get all values from Redis")
	}

	ind, err = GetDataFromRedis(6, res)
	if ind != 5 && err == nil {
		t.Error("Must get all values from Redis")
	}

}
