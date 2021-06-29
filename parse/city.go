package parse

import (
	"new_crawler/engine"
	"regexp"
)

const cityReg = `<a class="name" href="(https://www.ruoai.cn/info/[0-9]+)">([^<]+)</a>`

func City(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(cityReg)
	match := re.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}

	for _, v := range match {
		result.Items = append(result.Items, "User: "+string(v[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url:       string(v[1]),
			ParseFunc: Profile,
		})

	}

	return result
}
