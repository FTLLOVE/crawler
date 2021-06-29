//
// simple engine
//
//
package engine

import (
	"fmt"
	"new_crawler/fetcher"
)

type SimpleEngine struct{}

func (*SimpleEngine) Run(seeds ...Request) {
	var requests []Request

	for _, r := range seeds {
		requests = append(requests, r)
	}

	for len(requests) > 0 {
		currentReq := requests[0]
		requests = requests[1:]

		fmt.Printf("Fetch Url is : %s \n", currentReq.Url)

		parseResult, err := Worker(currentReq)
		if err != nil {
			continue
		}

		requests = append(requests, parseResult.Requests...)

		//TODO got data and persistence
		for _, i := range parseResult.Items {
			fmt.Printf("got item: %s \n", i)
		}
	}
}

// send url to fetch and request convert parseResult
func Worker(req Request) (ParseResult, error) {
	contents, err := fetcher.Fetch(req.Url)
	if err != nil {
		return ParseResult{}, err
	}
	return req.ParseFunc(contents), nil
}
