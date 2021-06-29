package parse

import (
	"fmt"
	"new_crawler/engine"
	"new_crawler/model"
	"regexp"
)

var (
	NameReg      = regexp.MustCompile(`<p class="nickname">([^<]+)</p>`)
	GenderReg    = regexp.MustCompile(`<span>([^<])士</span>`)
	BirthdayReg  = regexp.MustCompile(`<span>生日：([^<]+)</span>`)
	AgeReg       = regexp.MustCompile(`<span>年龄：([\d]+)岁</span>`)
	HeightReg    = regexp.MustCompile(`<span>体重：([^<]+)</span>`)
	WeightReg    = regexp.MustCompile(`<span>身高：([^<]+)</span>`)
	IncomeReg    = regexp.MustCompile(`<span>收入：([^<]+)</span>`)
	MarriageReg  = regexp.MustCompile(`<span>婚况：([^<]+)</span>`)
	EducationReg = regexp.MustCompile(`<span>学历：([^<]+)</span>`)
	HometownReg  = regexp.MustCompile(`<span>家乡：([^<]+)</span>`)
	HouseReg     = regexp.MustCompile(`<span>购房：([^<]+)</span>`)
)

func Profile(content []byte) engine.ParseResult {

	profile := model.Profile{}
	profile.Name = extractString(content, NameReg)
	profile.Gender = extractString(content, GenderReg)
	profile.Birthday = extractString(content, BirthdayReg)
	profile.Age = extractString(content, AgeReg)
	profile.Height = extractString(content, HeightReg)
	profile.Weight = extractString(content, WeightReg)
	profile.Income = extractString(content, IncomeReg)
	profile.Marriage = extractString(content, MarriageReg)
	profile.Education = extractString(content, EducationReg)
	profile.Hometown = extractString(content, HometownReg)
	profile.House = extractString(content, HouseReg)

	result := engine.ParseResult{
		Items: []interface{}{fmt.Sprintf("userInfo: %#v ", profile)},
	}

	return result
}

func extractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)
	if len(match) >= 2 {
		return string(match[1])
	}
	return ""
}
