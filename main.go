package main

import (
	"fmt"
	"github.com/google/uuid"
)

const checkCount = 200000000

func main() {
	sameUUID := make(map[string]int)

	mp := make(map[string]struct{})
	for i := 0; i < checkCount; i++ {
		curUUID := uuid.New().String()

		if _, ok := mp[curUUID]; ok {
			fmt.Printf("fined same uuid: %s", curUUID)

			recordSameUUID := func() {
				if v, ok := sameUUID[curUUID]; ok {
					sameUUID[curUUID] = v + 1
				} else {
					sameUUID[curUUID] = 2
				}
			}
			recordSameUUID()
		} else {
			mp[curUUID] = struct{}{}
		}
	}

	if len(sameUUID) == 0 {
		fmt.Printf("test uuid of %d ok!!!!!\n", checkCount)
	} else {
		fmt.Printf("test uuid of %d not ok.......\n", checkCount)
		fmt.Println("we have same uuids: ")
		fmt.Println("uuid \t count")
		for k, v := range sameUUID {
			fmt.Printf("%s %d\n", k, v)
		}
	}
}
