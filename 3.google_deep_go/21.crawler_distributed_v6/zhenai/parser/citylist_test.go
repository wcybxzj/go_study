package parser

import (
	"io/ioutil"
	"testing"
)

func TestParserCityList(t *testing.T) {
	//contents, err := fetcher.Fetch("http://city.zhenai.com")
	contents, err := ioutil.ReadFile("./citylist_test_data.html")

	if err != nil {
		panic(err)
	}

	//debug
	//fmt.Printf("%s\n",contents)

	result := ParserCityList(contents, "")

	const resultSize = 470

	expectedUrls := []string{
		"http://city.zhenai.com/aba",
		"http://city.zhenai.com/akesu",
		"http://city.zhenai.com/alashanmeng",
	}

	if len(result.Requests) != resultSize {
		t.Errorf("need %d ,but %d",
			resultSize, len(result.Requests))
	}

	for i, url := range expectedUrls {
		if result.Requests[i].Url != url {
			t.Errorf("expected url #%d: %s but was %s",
				i, url, result.Requests[i].Url)
		}
	}

}
