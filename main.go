package main

import (
	"fmt"
	"github.com/google/uuid"
	"time"
)

const checkCount = 200000000

func main() {
	w := createWorker(func(nowPosition int) {
		const expectCount = 10
		const expectPercent int = 100 / expectCount

		curPercent := float64(nowPosition) / float64(checkCount) * 100
		if curPercent == float64(int(curPercent)) && int(curPercent)%expectPercent == 0 {
			fmt.Printf("%d%%\n", int(curPercent))
		}
	})

	sameUUID := make(map[string]int)

	mp := make(map[string]struct{})
	for i := 1; i <= checkCount; i++ {
		w <- i
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

	// wait goroutine "status" complete
	time.Sleep(time.Millisecond)

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
