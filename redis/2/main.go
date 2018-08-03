package main

import (
	"github.com/go-redis/redis"
	"fmt"
	"sync"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	defer client.Close()

	var wg sync.WaitGroup
	for i := 0; i < 100; i++  {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			client.Incr("mykeyyou22")
			v, _:= client.Get("mykeyyou22").Result()
			fmt.Printf("%d: %s\n", i, v)
		}(i)
	}
	wg.Wait()

	v, err := client.Get("mykeyyou22").Result()
	if err == nil {
		fmt.Println(v) // not 100
	}
}
