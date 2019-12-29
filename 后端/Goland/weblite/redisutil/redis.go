package redisutil

import (
	"fmt"
	"os"

	redis "github.com/gomodule/redigo/redis"
)

func Incr() {
	c, err := redis.Dial("tcp", "localhost:32771")
	errCheck(err)
	c.Send("auth", "123456")
	defer c.Close()

	_, setErr := c.Do("set", "url", "nihao ")
	errCheck(setErr)

	//使用redis的string类型获取set的k/v信息
	r, getErr := redis.String(c.Do("get", "url"))
	errCheck(getErr)
	fmt.Println(r)

}

func errCheck(err error) {
	if err != nil {
		fmt.Println("sorry,has some error:", err)
		os.Exit(-1)
	}
}
