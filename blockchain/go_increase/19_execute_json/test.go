package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func poster_data(pageIndex int, statsType string) (body []byte) {
	url := "https://api.filscout.com/api/v1/minerrank/blocks"
	contentType := "application/json"
	//data := `{"continent": 0, "pageIndex": 2 ,"pageSize": 50, "sectorSize": "", "statsType": "24h"} `
	data := fmt.Sprintf("{\"continent\":0,\"pageIndex\":%d,\"pageSize\":50,\"sectorSize\":\"\",\"statsType\":\"%s\"}", pageIndex, statsType)
	fmt.Println(data)
	resp, err := http.Post(url, contentType, strings.NewReader(data))
	if err != nil {
		fmt.Printf("get failed, err:%v\n", err)
		return
	}
	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("read from resp.Body failed, err:%v\n", err)
		return
	}
	return body
}

func main() {
	body := poster_data(1, "24h")
	fmt.Printf("%#v", body)
}
