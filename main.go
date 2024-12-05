package main

import (
	"fmt"
	"GoTraining/cache"
)

func main() {
	newCache := cache.New()

	newCache.Set("userId1", 123)
	newCache.Set("userId2", "456")
	newCache.Set("userId3", "789")

	fmt.Println(newCache.Get("userId3"))

	err := newCache.Delete("userId1")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(newCache)
}
