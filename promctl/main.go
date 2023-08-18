package main

import (
	"context"
	"fmt"
	api "github.com/prometheus/client_golang/api"
	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
	"github.com/prometheus/common/model"
	"log"
	"os"
	"time"
)

//
func st_en() (time.Time, time.Time) {
	layout := "2006-01-02 15:04:05"

	// 获取用户输入的起始时间年月日时分秒
	var sYear, sMonth, sDay, sHour, sMinute, sSecond string

	fmt.Println("请输入起始时间（格式：2006 01 02 15 04 05）：")
	fmt.Scanf("%s %s %s %s %s %s", &sYear, &sMonth, &sDay, &sHour, &sMinute, &sSecond)
	sTime := fmt.Sprintf("%v-%v-%v %v:%v:%v", sYear, sMonth, sDay, sHour, sMinute, sSecond)

	// 获取用户结束时间
	var eYear, eMonth, eDay, eHour, eMinute, eSecond string

	fmt.Println("请输入结束时间（格式：2006 01 02 15 04 05）：")
	fmt.Scanf("%s %s %s %s %s %s", &eYear, &eMonth, &eDay, &eHour, &eMinute, &eSecond)
	eTime := fmt.Sprintf("%v-%v-%v %v:%v:%v", eYear, eMonth, eDay, eHour, eMinute, eSecond)

	// 上海时区
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println("无效的时区")
		os.Exit(1)
	}

	start, _ := time.ParseInLocation(layout, sTime, loc)
	end, _ := time.ParseInLocation(layout, eTime, loc)

	return start, end
}

func main() {
	//自定义用户输入 Prometheus API
	PApi := ""
	fmt.Println("请输入 Prometheus 地址如（127.0.0.1:9090） :")
	fmt.Scanln(&PApi)
	newPApi := "http://" + PApi

	// 用户输入查询PromQL
	PromQl := ""
	fmt.Println("请输入查询 PromQL 如（up）:")
	fmt.Scanln(&PromQl)

	// 调用起始和结束时间函数
	start, end := st_en()

	// 创建 Prometheus 客户端配置
	cfg := api.Config{
		Address:      newPApi, // 替换为 Prometheus 服务器地址
		RoundTripper: api.DefaultRoundTripper,
	}

	// 创建 Prometheus 客户端
	client, err := api.NewClient(cfg)
	if err != nil {
		log.Fatal(err)
	}

	// 创建 Prometheus API
	Papi := v1.NewAPI(client)

	// 执行查询
	query := PromQl // 替换为你的查询表达式
	result, warnings, err := Papi.QueryRange(context.Background(), query, v1.Range{
		Start: start,
		End:   end,
		Step:  time.Second * 30,
	})
	if err != nil {
		log.Fatal(err)
	}
	if len(warnings) > 0 {
		log.Printf("Warnings: %v", warnings)
	}
	if result.Type() != model.ValMatrix {
		log.Fatalf("Expected a matrix result, got %v", result.Type())
	}

	// 处理查询结果
	matrix := result.(model.Matrix)
	if matrix.Len() == 0 {
		fmt.Println("注意：当前查询时间段指标没有数据！")
	} else {
		for _, sample := range matrix {
			if len(sample.Values) > 0 {
				value := sample.Values[len(sample.Values)-1].Value
				fmt.Printf("查询时间：%v\n查询指标：%s 查询数据：%s\n", start, sample.Metric, value)
			}
		}
	}

}
