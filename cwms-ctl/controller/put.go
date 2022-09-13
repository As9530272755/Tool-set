package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

func Put(payload interface{}, api string) {

	// 转为 byte 格式
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		log.Panic(err)
	}

	// 转为 reader
	body := strings.NewReader(string(jsonPayload))

	// 提交数据
	req, _ := http.NewRequest("PUT", api, body)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("token", Login())

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()
}
