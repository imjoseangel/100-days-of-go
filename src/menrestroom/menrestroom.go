///bin/true; exec /usr/bin/env go run "" "$@"
package main

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
	"strings"
	"time"

	// run: `go get github.com/imjoseangel/functions-go@v1.0.5`
	"github.com/imjoseangel/functions-go"
)

var (
	untaken    []int
	taken      []int
	timePeeing time.Duration
	newStall   int
	left       []int
	right      []int
	stall      int
	stallPrint []string
)

const (
	stalls        int           = 10
	stallFreq     time.Duration = 1
	mintimepeeing int           = 1
	maxtimepeeing int           = 10
	emoEmpty      string        = "\U0001F6BD"
	emoTaken      string        = "\U0001F6B6"
	emoDoor       string        = "\U0001F6AA"
)

// func makeRange(min int, max int) []int {
// 	NewList := make([]int, max-min+1)
// 	for Item := range NewList {
// 		NewList[Item] = min + Item
// 	}
// 	return NewList
// }

// func sumArray(array []int) int {
// 	result := 0
// 	for _, numb := range array {
// 		result += numb
// 	}
// 	return result
// }

// func sliceArray(array []int, start int) []int {

// 	result := []int{}
// 	for i := start; i < len(array); i += 2 {
// 		result = append(result, array[i])
// 	}
// 	return result
// }

// func index(array []int, item int) int {
// 	for result := range array {
// 		if array[result] == item {
// 			return result
// 		}
// 	}
// 	return -1
// }

func init() {

	rand.Seed(time.Now().UnixNano())
	timePeeing = time.Duration(rand.Intn(maxtimepeeing-mintimepeeing+1) + mintimepeeing)
	untaken = functions.Range{
		MinList: 1,
		MaxList: stalls,
	}.RangeArray()

	// untaken = makeRange(1, stalls)

	newStall = int(math.Ceil(float64(functions.Array(untaken).SumArray()) / float64(len(untaken))))
	if stalls%2 == 0 {
		left = functions.SliceArray(untaken[0:newStall-1], 1)
	} else {
		left = functions.SliceArray(untaken[0:newStall-1], 0)
	}
	right = functions.SliceArray(untaken[newStall:], 1)
	stallPrint = strings.SplitN(strings.Repeat(emoEmpty, stalls)+emoDoor, "", stalls+1)

}

func takeStall() []string {

	if len(untaken) > 0 {
		if len(taken) == 0 {
			stall = newStall
		} else {
			if len(left) > 0 {
				randomIndex := rand.Intn(len(left))
				stall = left[randomIndex]
				left = append(left[:randomIndex], left[randomIndex+1:]...)
			} else if len(right) > 0 {
				randomIndex := rand.Intn(len(right))
				stall = right[randomIndex]
				right = append(right[:randomIndex], right[randomIndex+1:]...)

			} else {
				randomIndex := rand.Intn(len(untaken))
				stall = untaken[randomIndex]
			}
		}
		stallIndex := functions.IndexArray(untaken, stall)
		untaken = append(untaken[:stallIndex], untaken[stallIndex+1:]...)
	}

	taken = append(taken, stall)
	stallPrint[taken[len(taken)-1]-1] = emoTaken

	return stallPrint
}

func leaveStall() {

	if len(taken) > 0 {
		oldStall := taken[0]
		taken = append(taken[:0], taken[1:]...)
		untaken = append(untaken, oldStall)
		sort.Ints(untaken)
		stallPrint[oldStall-1] = emoEmpty
		rand.Seed(time.Now().UnixNano())
		timePeeing = time.Duration(rand.Intn(maxtimepeeing-mintimepeeing+1) + mintimepeeing)
	}
}

func main() {

	takeTicker := time.NewTicker(stallFreq * time.Second)
	leaveTicker := time.NewTicker(timePeeing * time.Second)

	for len(untaken) > 0 {
		fmt.Println("\033[H\033[2J")
		for _, item := range stallPrint {
			fmt.Print(item + " ")
		}

		select {
		case <-takeTicker.C:
			takeStall()
		case <-leaveTicker.C:
			leaveStall()
			leaveTicker.Reset(timePeeing * time.Second)
		}
	}

	fmt.Println("\033[H\033[2J")
	for _, item := range stallPrint {
		fmt.Print(item + " ")
	}

	fmt.Println()

}
