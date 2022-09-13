package controller

import (
	"cwms-ctl/model"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strings"
)

// 解析 response 请求转换为 map[string]map[string]interface{} 类型
func parseResponse(response *http.Response) (map[string]map[string]interface{}, error) {
	var result map[string]map[string]interface{}

	body, err := ioutil.ReadAll(response.Body)
	if err == nil {
		err = json.Unmarshal(body, &result)
	}

	return result, err
}

func QueryWarn(rluse string) (ID string) {
	// 该变量用于获取检索内容
	var warnruleList map[string]interface{}

	// 告警规则 ID
	id := ""

	api := Config() + `/warnrule/list`

	// 通过告警规则名称进行查询
	payload := model.NewWarnrule(rluse)

	jsonPayload, _ := json.Marshal(payload)

	// 转为 reader 类型
	body := strings.NewReader(string(jsonPayload))

	// 提交请求
	req, err := http.NewRequest("POST", api, body)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("token", Login())

	// 判读没有错误那么就将这次的请求体传递给 parseResponse() 进行解析
	response, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	} else {
		returnMap, _ := parseResponse(response)
		data := returnMap["data"]["searchList"]

		// 类型断言并将对应的 tmp 赋值给 tmp2
		tmp, ok := data.([]interface{})
		if !ok {
			panic("失败")
		}
		for _, v := range tmp {
			warnruleList = v.(map[string]interface{})
		}
	}
	defer response.Body.Close()

	// 解析 response 数据
	data, _ := json.Marshal(warnruleList)
	if string(data) == "null" {
		log.Println("查询失败请检查告警规则名称是否输入正确!")
	}

	// 通过正则取出 id
	compileRegex := regexp.MustCompile("id\":(.*?),\"message_receiver")
	matchArr := compileRegex.FindStringSubmatch(string(data))

	id = matchArr[len(matchArr)-1]

	return id
}

// 用来判断当前新增的告警字段中是否有对应的 namespace 字段，如果有就报错
func QueryWarnRuleDetail(id string, alertPayload interface{}, WarnName string) {
	// 该变量用于获取检索内容
	var warnRuleDetail []interface{}

	// api
	api := Config() + `/warnRuleDetail/list`

	payload := model.Warnrule{IndexArray: []struct {
		ColName   string `json:"colName"`
		ColType   string `json:"colType"`
		IndexType string `json:"indexType"`
		Relation  string `json:"relation"`
		Value     string `json:"value"`
	}{
		{
			ColName:   "rule_id",
			IndexType: "1",
			Relation:  "1",
			Value:     id,
		},
	}, IsPage: false}

	jsonPayload, _ := json.Marshal(payload)

	body := strings.NewReader(string(jsonPayload))

	// 提交请求
	req, err := http.NewRequest("POST", api, body)
	if err != nil {
		log.Panic(err)
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("token", Login())

	// 判读没有错误那么就将这次的请求体传递给 parseResponse() 进行解析
	response, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	} else {
		returnMap, _ := parseResponse(response)
		data := returnMap["data"]["searchList"]
		tmp, ok := data.([]interface{})
		if ok {
			warnRuleDetail = tmp
		}
	}

	defer response.Body.Close()

	byteData, _ := json.Marshal(warnRuleDetail)
	data := string(byteData)

	// 用于判断告警字段是否相同,如果相同就直接退出程序并回显
	reg := regexp.MustCompile(`"warn_field":(.*?),"warn_value"`).FindAllString(data, -1)
	for i := 0; i < len(reg); i++ {
		if strings.Contains(reg[i], WarnName) {
			log.Println("新增告警字段重复请 web 页面查看！")
			return
		}
	}
	Put(alertPayload, Config()+`/warnRuleDetail/update`)
}
