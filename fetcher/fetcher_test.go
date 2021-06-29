package fetcher

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestFetch(t *testing.T) {
	contents, err := Fetch("https://www.ruoai.cn/location")
	t.Log(string(contents), err)
}

func TestFor(t *testing.T) {
	s := []int{1, 2, 3, 4}
	var newSlice []int
	for _, el := range s {
		newSlice = append(newSlice, el)
	}

	for len(newSlice) > 0 {
		fmt.Printf("newSlice: %v \n", newSlice)
		currentEl := newSlice[0]
		newSlice = newSlice[1:]

		fmt.Printf("currentEle is: %v \n", currentEl)
		newSlice = append(newSlice, generateInt()...)
	}

	t.Log(newSlice)
}

func generateInt() (s []int) {
	s = make([]int,1)
	s[0] = rand.Intn(10)
	return
}
