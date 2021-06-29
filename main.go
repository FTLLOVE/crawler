package main

import (
	"fmt"
	"new_crawler/engine"
	"new_crawler/parse"
	"time"
)

const URL = "https://www.ruoai.cn/location"

func main() {
	startTime := time.Now()
	e := &engine.SimpleEngine{}

	e.Run(engine.Request{
		Url:       URL,
		ParseFunc: parse.CityList,
	})

	// calculation time
	fmt.Printf("reducer time: %v", time.Since(startTime).Seconds())
}
