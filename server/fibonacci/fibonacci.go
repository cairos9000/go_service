package fibonacci

import (
	"errors"
	"github.com/cairos9000/go_service/constants"
	"github.com/cairos9000/go_service/redis"
	"log"
	"strconv"
	"strings"
)

func ParseArgs(text, proto string) (int, int, error) {
	var start int
	if proto == constants.Http {
		start = 1
	} else if proto == constants.Grpc {
		start = 0
	} else {
		return 0, 0, errors.New("incorrect protocol was sent")
	}

	if text[start] != '[' || text[len(text)-1] != ']' {
		return 0, 0, errors.New("incorrect view of interval it must be in format [x;y]")
	}

	args := strings.Split(text[start+1:len(text)-1], ";")
	if len(args) != 2 {
		return 0, 0, errors.New("incorrect interval. There must be two values")
	}
	args[0] = strings.Trim(args[0], " ")
	args[1] = strings.Trim(args[1], " ")

	x, valueError1 := strconv.Atoi(args[0])
	if valueError1 != nil {
		return 0, 0, errors.New("incorrect x value. Must be integer")
	}

	y, valueError2 := strconv.Atoi(args[1])
	if valueError2 != nil {
		return 0, 0, errors.New("incorrect y value. Must be integer")
	}

	return x, y, nil
}

func Fibo(x, y int) ([]int, error) {
	if x < 1 || y < 1 {
		return nil, errors.New("value of interval can't be negative or equal to zero")
	}

	if x > y {
		return nil, errors.New("value x is bigger than y")
	}

	var (
		fib1, fib2 int
		ind        int
		redisError error
		res        = make([]int, y)
	)

	if redis.Connection != nil {
		ind, redisError = redis.GetDataFromRedis(y, res)
		if redisError == nil {
			return res[x-1 : y], nil
		}
	}

	if ind == 1 || ind == 0 {
		fib1, fib2 = 1, 1
		ind = 2
		res[0] = 1
		res[1] = 1

	} else {
		fib1 = res[ind-2]
		fib2 = res[ind-1]
	}

	for ; ind < y; ind++ {
		res[ind] = fib1 + fib2
		fib1 = fib2
		fib2 = res[ind]
	}

	if redis.Connection != nil {
		err := redis.SetDataToRedis(res)
		if err != nil {
			log.Println(err)
		}
	}

	return res[x-1 : y], nil
}
