package qqmfjh_com

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

// 大发快三
// https://qqmfjh.com/plan/api.do?code=dfk3&plan=0&size=4&planSize=1
type Res struct {
	code int
	data map[string]interface{}
}

func run() {
	resp, err := http.Get("https://qqmfjh.com/plan/api.do?code=dfk3&plan=0&size=4&planSize=1")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	// 初始化请求变量结构
	formData := make(map[string]interface{}, 2)

	err = json.Unmarshal(body, &formData)
	if err != nil {
		return
	}
	//fmt.Println(res)
	for k, v := range formData {
		switch vType := v.(type) {
		case int:
			fmt.Println("int", k, strconv.Itoa(vType))
		case string:
			fmt.Println("string", k, vType)

		case float32:
			fmt.Println("float32", k, strconv.FormatFloat(float64(vType), 'f', 2, 64))

		case float64:
			fmt.Println("float64", k, strconv.FormatFloat(vType, 'f', 2, 64))
		case interface{}:

		}

	}

}
