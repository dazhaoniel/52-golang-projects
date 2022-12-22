package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
)

const gt = "https://translate.googleapis.com/translate_a/single?client=gtx&sl=auto&tl=%s&dt=t&q=%s"

func main() {
	argLang := os.Args[1]

	supported := map[string]string{
		"chinese":  "zh-CN",
		"english":  "en",
		"japanese": "ja",
		"spanish":  "es",
		"arabic":   "ar",
		"french":   "fr",
	}

	if _, ok := supported[argLang]; !ok {
		panic("Target language not supported!")
	}

	argsWithoutProg := os.Args[3:]
	sentence := strings.Join(argsWithoutProg, " ")

	url := fmt.Sprintf(gt, supported[argLang], url.PathEscape(sentence))

	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	var resp []interface{}

	err = json.Unmarshal(responseData, &resp)
	if err != nil {
		panic(err)
	}

	translatedText := ""
	for _, obj := range resp[0].([]interface{}) {

		if len(obj.([]interface{})) == 0 {
			break
		}
		translatedText += obj.([]interface{})[0].(string)
	}

	fmt.Println(translatedText)
}
