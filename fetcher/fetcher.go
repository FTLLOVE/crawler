//
// the package has only one method,its function is to get the url and send http request,
// and return the obtained byte slice data
//
package fetcher

import (
	"errors"
	"io/ioutil"
	"net/http"
)

const (
	UserAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_4) AppleWebKit/537.31 (KHTML, like Gecko) Chrome/71.0.3578.87 Safari/537.21"
)

func Fetch(url string) ([]byte, error) {

	newReq, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	newReq.Header.Add("User-Agent", UserAgent)

	client := http.Client{CheckRedirect: func(req *http.Request, via []*http.Request) error {
		return nil
	}}

	resp, err := client.Do(newReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		return ioutil.ReadAll(resp.Body)
	}

	return nil, errors.New("fetch occur error")
}
