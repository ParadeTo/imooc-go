package main

import (
	"time"
	"github.com/gomodule/redigo/redis"
)

func main() {
	pool := &redis.Pool{
		MaxIdle:3,
		IdleTimeout: 240 * time.Second,
		Dial:        func() (redis.Conn, error) { return redis.Dial("tcp", "localhost:6379") },
	}
	conn := pool.Get()
	defer conn.Close()
	for i := 0; i < 100; i++  {
		go func() {
			conn.Send("INCR", "mykey4")
			conn.Flush()
			//var v int
			//if vR == nil {
			//	v = 0
			//} else {
			//	fmt.Printf("%+v", vR.([]uint8)[0])
			//	v = int(vR.([]uint8)[0])
			//}
			//fmt.Println(int(v))
			//conn.Send("SET", "key", int(v) + 1)
			//conn.Send("GET", "key")
			//conn.Flush()
			//vR, err = conn.Receive()
			//fmt.Println(vR)
		}()
	}


	//var a = "1"
	//fmt.Println(uint8(a))
	time.Sleep(10 * time.Second)
}
