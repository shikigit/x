package parser

import (
	"encoding/json"
	"fmt"
	"github.com/crawlerclub/x/downloader"
	"github.com/crawlerclub/x/types"
	"io/ioutil"
	"testing"
)

func TestParse(t *testing.T) {
	conf, _ := ioutil.ReadFile("../cmd/www.newsmth.net.json")

	var crawlerConf types.CrawlerConf
	err := json.Unmarshal(conf, &crawlerConf)
	if err != nil {
		t.Fatal(err)
	}
	urlConf := crawlerConf.ParseConfs["article"]
	//pageUrl := urlConf.ExampleUrl
	pageUrl := "http://www.newsmth.net/nForum/article/Orienteering/59230"
	requestInfo := &types.HttpRequest{
		Url:      pageUrl,
		Method:   "GET",
		UseProxy: false,
		Platform: "pc",
	}

	fmt.Println(pageUrl)
	responseInfo := downloader.Download(requestInfo)
	if responseInfo.Error != nil {
		t.Fatal(responseInfo.Error)
	}
	fmt.Println(responseInfo.Encoding)

	retUrls, retItems, err := Parse(responseInfo.Content, pageUrl, &urlConf)
	if err != nil {
		t.Fatal(err)
	}
	jsonUrls, _ := json.Marshal(retUrls)
	jsonItems, _ := json.Marshal(retItems)
	t.Log("retUrls json: ", string(jsonUrls))
	t.Log("retItems json: ", string(jsonItems))
}
