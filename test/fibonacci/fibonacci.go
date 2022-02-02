package fibonacci

import (
	"errors"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"strconv"
	"strings"
	"test/constants"
)

var RedisConnection redis.Conn

func DisconnectRedis(c redis.Conn) error {
	err := c.Close()
	return err
}

func ConnectToRedis(address string)(redis.Conn, error){
	c, err := redis.Dial(constants.Tcp, address)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func ParseArgs(text, proto string)(int, int, error){
	var start int
	if proto == constants.Http{
		start = 1
	} else if proto == constants.Grpc {
		start = 0
	} else{
		return 0, 0, errors.New("incorrect protocol was sent")
	}

	if text[start] != '[' || text[len(text) - 1] != ']'{
		return 0, 0, errors.New("incorrect view of interval it must be in format [x;y]")
	}

	args := strings.Split(text[start + 1: len(text) - 1], ";")
	if len(args) != 2{
		return 0, 0, errors.New("incorrect interval. There must be two values")
	}
	args[0] = strings.Trim(args[0], " ")
	args[1] = strings.Trim(args[1], " ")

	x, valueError1 := strconv.Atoi(args[0])
	if valueError1 != nil{
		return 0, 0, errors.New("incorrect x value. Must be integer")
	}

	y, valueError2 := strconv.Atoi(args[1])
	if valueError2 != nil{
		return 0, 0, errors.New("incorrect y value. Must be integer")
	}

	return x, y, nil
}

func GetDataFromRedis(y int, res []int) (int, error) {
	ind := 0
	var err error

	for ; ind < y; ind++{
		res[ind], err = redis.Int(RedisConnection.Do(constants.RedisGetMethod, ind))
		if err != nil{
			return ind, err
		}
	}

	return 0, nil

}

func SetDataToRedis(fibonacciNumbers []int) error {
	var err error

	for num, elem := range fibonacciNumbers{
		_, err = RedisConnection.Do(constants.RedisSetMethod, num, elem)
		if err != nil {
			return err
		}
	}
	return nil
}

func Fibo(x, y int) ([]int, error) {
	if x < 1 || y < 1{
		return nil, errors.New("value of interval can't be negative or equal to zero")
	}

	if x > y{
		return nil, errors.New("value x is bigger than y")
	}

	var (
		fib1, fib2 int
		ind int
		redisError error
		res = make([]int, y)
	)

	if RedisConnection != nil{
		ind, redisError = GetDataFromRedis(y, res)
		if redisError == nil{
			return res[x - 1: y], nil
		}
	}


	if ind == 1 || ind == 0{
		fib1, fib2 = 1, 1
		ind = 2
		res[0] = 1
		res[1] = 1

	} else {
		fib1 = res[ind - 2]
		fib2 = res[ind - 1]
	}

	for ; ind < y; ind++{
		res[ind] = fib1 + fib2
		fib1 = fib2
		fib2 = res[ind]
	}

	if RedisConnection != nil{
		err := SetDataToRedis(res)
		if err != nil{
			fmt.Println(err)
		}
	}

	return res[x - 1: y], nil
}
