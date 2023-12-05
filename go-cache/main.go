package main

import (
	"fmt"
	"go-cache/simple/cache"
)

func main() {
	hash := cache.New()

	hash.Set("userId", 1)

	userId := hash.Get("userId")
	fmt.Println(userId)

	hash.Delete("userId")

	fmt.Println(hash["userId"])
}
