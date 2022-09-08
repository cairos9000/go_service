package redis

import (
	"github.com/cairos9000/go_service/constants"
	"github.com/garyburd/redigo/redis"
)

var Connection redis.Conn

func DisconnectRedis(c redis.Conn) error {
	err := c.Close()
	return err
}

func ConnectToRedis(address string) (redis.Conn, error) {
	c, err := redis.Dial(constants.Tcp, address)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func GetDataFromRedis(y int, res []int) (int, error) {
	ind := 0
	var err error

	for ; ind < y; ind++ {
		res[ind], err = redis.Int(Connection.Do(constants.RedisGetMethod, ind))
		if err != nil {
			return ind, err
		}
	}

	return 0, nil

}

func SetDataToRedis(fibonacciNumbers []int) error {
	var err error

	for num, elem := range fibonacciNumbers {
		_, err = Connection.Do(constants.RedisSetMethod, num, elem)
		if err != nil {
			return err
		}
	}
	return nil
}
