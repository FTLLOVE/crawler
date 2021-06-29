//
// The package is a structural model for processing request and response
//

package engine

//
// request model include Url and function ParseFunc
// the url parameter for send a fetch request
// The parseResult method accepts a request byte slice,return parseResult
//
type Request struct {
	Url       string
	ParseFunc func([]byte) ParseResult
}

//
// parseResult model include items and send request slice
// item is the data obtained
// request slice is the request model that needs to send data next request
//
type ParseResult struct {
	Items    []interface{}
	Requests []Request
}
