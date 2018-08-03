package main

import (
	"fmt"
	"time"
	"github.com/go-redis/redis"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr: "0.0.0.0:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	defer client.Close()

	//var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		//wg.Add(1)
		go func() {
			//defer wg.Done()
			err := client.Incr("mykeykey22").Err()
			if err != nil {
				panic(err)
			}
		}()
	}
	//wg.Wait()
	time.Sleep(time.Second * 5)
	v, err := client.Get("mykeykey22").Result()
	if err == nil {
		fmt.Println(v) // not 100
	}
}