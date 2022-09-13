package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

func Login() string {
	var result map[string]string
	api := Config() + `/user/login`

	// 获取 token
	response, err := http.PostForm(api, url.Values{
		"username": {User},
		"password": {Passwrod},
	})
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	// 请求体，用来获取 token
	body, err := ioutil.ReadAll(response.Body)
	if err == nil {
		json.Unmarshal([]byte(string(body)), &result)
	}

	// 获取 token
	token := result["token"]
	return token
}
