package parse

import (
	"new_crawler/engine"
	"regexp"
)

var cityListReg = regexp.MustCompile(`<a href="(https://www.ruoai.cn/[0-9]+/)">([^<]+)</a>`)

func CityList(contents []byte) (res engine.ParseResult) {
	match := cityListReg.FindAllSubmatch(contents, -1)

	for _, v := range match {
		res.Items = append(res.Items, "City="+string(v[2]))
		res.Requests = append(res.Requests, engine.Request{
			Url:       string(v[1]),
			ParseFunc: City,
		})
	}

	return
}
