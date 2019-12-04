package main

import (
	"fmt"
	"os"
	"strings"
	"net/http"
	"net/url"
	"io/ioutil"
	"encoding/json"
)

const gt = "https://translate.googleapis.com/translate_a/single?client=gtx&sl=auto&tl=%s&dt=t&q=%s"


func main() {
	argsWithoutProg := os.Args[2:]
	sentence := strings.Join(argsWithoutProg," ")

	url := fmt.Sprintf(gt, "es", url.PathEscape(sentence))

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