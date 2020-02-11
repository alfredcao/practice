package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

const googleSearchResult = `
{
	"responseData" : {
		"results": [
        {
            "GsearchResultClass": "GwebSearch",
            "cacheUrl": "http://www.google.com/search?q=cache:Ry2YgNocTkEJ:c.mi.com",
            "clicktrackUrl": "https://www.google.com/url?q=http://c.mi.com/in/space-uid-534194697.html&sa=U&ved=0ahUKEwi_6bvGlKTOAhUH_WMKHVA-AxIQFggEMAA&client=internal-uds-cse&usg=AFQjCNGE3VOe4N36-OZxOg-Cw-IXKUigAg",
            "content": "Rajesh <b>123</b> Profile ,Mi Community. ... Profile. ID: 534194697. Nickname: Rajesh <b>123</b>. Status: Offline. Gender: Secret. Day of birth: - ...",
            "contentNoFormatting": "Rajesh 123 Profile ,Mi Community. ... Profile. ID: 534194697. Nickname: Rajesh 123. Status: Offline. Gender: Secret. Day of birth: - ...",
            "formattedUrl": "c.mi.com/in/space-uid-534194697.html",
            "title": "Rajesh <b>123</b> Profile - Mi Community",
            "titleNoFormatting": "Rajesh 123 Profile - Mi Community",
            "unescapedUrl": "http://c.mi.com/in/space-uid-534194697.html",
            "url": "http://c.mi.com/in/space-uid-534194697.html",
            "visibleUrl": "c.mi.com"
        },
        {
            "GsearchResultClass": "GwebSearch",
            "cacheUrl": "http://www.google.com/search?q=cache:5s2uT8-PeLYJ:c.mi.com",
            "clicktrackUrl": "https://www.google.com/url?q=http://c.mi.com/in/space-uid-1544909598.html&sa=U&ved=0ahUKEwi_6bvGlKTOAhUH_WMKHVA-AxIQFggGMAE&client=internal-uds-cse&usg=AFQjCNFy5-U3EsLYAVTMLO-7iEXMSuHXjg",
            "content": "ID: 1544909598. Nickname: Arijit <b>Kundu123</b>. Status: Offline. Real Name: Arijit Kundu. Gender: Male. Education degree: Bachelor. Interests: Playing Pc games.",
            "contentNoFormatting": "ID: 1544909598. Nickname: Arijit Kundu123. Status: Offline. Real Name: Arijit Kundu. Gender: Male. Education degree: Bachelor. Interests: Playing Pc games.",
            "formattedUrl": "c.mi.com/in/space-uid-1544909598.html",
            "title": "Arijit <b>Kundu123</b> Profile - Mi Community",
            "titleNoFormatting": "Arijit Kundu123 Profile - Mi Community",
            "unescapedUrl": "http://c.mi.com/in/space-uid-1544909598.html",
            "url": "http://c.mi.com/in/space-uid-1544909598.html",
            "visibleUrl": "c.mi.com"
        }
    	]
	}
}
`

type GSearchResult struct {
	GsearchResultClass  string `json:"GsearchResultClass"`
	CacheUrl            string `json:"cacheUrl"`
	ClicktrackUrl       string `json:"clicktrackUrl"`
	Content             string `json:"content"`
	ContentNoFormatting string `json:"contentNoFormatting"`
	FormattedUrl        string `json:"formattedUrl"`
	Title               string `json:"title"`
	TitleNoFormatting   string `json:"titleNoFormatting"`
	UnescapedUrl        string `json:"unescapedUrl"`
	Url                 string `json:"url"`
	VisibleUrl          string `json:"visibleUrl"`
}

type GSearchResponseData struct {
	ResponseData struct {
		Results []GSearchResult `json:"results"`
	} `json:"responseData"`
}

func main() {
	//demo1()
	//demo2()
	demo3()
}

func demo1() {
	var gSearchResponseData GSearchResponseData
	decoder := json.NewDecoder(strings.NewReader(googleSearchResult))
	err := decoder.Decode(&gSearchResponseData)
	if err != nil {
		fmt.Println("decode failed, err :", err)
		return
	}

	fmt.Println(gSearchResponseData)
}

func demo2() {
	var gSearchResponseData GSearchResponseData
	err := json.Unmarshal([]byte(googleSearchResult), &gSearchResponseData)
	if err != nil {
		fmt.Println("decode failed, err :", err)
		return
	}

	fmt.Println(gSearchResponseData)
}

func demo3() {
	c := make(map[string]interface{})
	c["name"] = "Gopher"
	c["title"] = "programmer"
	c["contact"] = map[string]interface{}{
		"home": "415.333.3333",
		"cell": "415.555.5555",
	}

	data, err := json.MarshalIndent(c, "", "	")
	if err != nil {
		fmt.Println("encode failed, err :", err)
		return

	}
	fmt.Println(string(data))
}
