package main

import (
	"fmt"
	"github.com/google/uuid"
)

const checkCount = 200000000

func main() {
	mp := make(map[string]struct{})
	for i := 0; i < checkCount; i++ {
		curUUID := uuid.New().String()

		if _, ok := mp[curUUID]; ok {
			fmt.Printf("fined same uuid: %s", curUUID)
		} else {
			mp[curUUID] = struct{}{}
		}
	}
	for k, _ := range mp {
		fmt.Println(k)
	}
}
