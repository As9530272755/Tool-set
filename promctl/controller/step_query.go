package controller

import (
	"context"
	"demo3/config"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	api "github.com/prometheus/client_golang/api"
	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
	"github.com/prometheus/common/model"
	"log"
	"os"
	"strconv"
	"time"
)

// 基于起始时间和结束时间做全量查询
func Step_Query(Step_query, Step_startTime, Step_number string) {
	layout := "2006-01-02 15:04:05"

	// 上海时区
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println("无效的时区")
		os.Exit(1)
	}

	// 时区转换
	start, _ := time.ParseInLocation(layout, Step_startTime, loc)

	// 读取配置文件中的地址
	API, err := config.ParseConfig()
	if err != nil {
		log.Panic(err)
	}

	// Prometheus API 拼接
	promURL := API.Prometheus.Addr + ":" + API.Prometheus.Port

	fmt.Println(promURL)

	// 创建 Prometheus 客户端配置
	cfg := api.Config{
		Address:      promURL, // 替换为 Prometheus 服务器地址
		RoundTripper: api.DefaultRoundTripper,
	}

	// 创建 Prometheus 客户端
	client, err := api.NewClient(cfg)
	if err != nil {
		log.Fatal(err)
	}

	// 创建 Prometheus API
	Papi := v1.NewAPI(client)

	// 对 step 查询次数做类型转换
	eStep, _ := strconv.Atoi(Step_number)

	// 创建一个新的Excel文件
	f := excelize.NewFile()
	// 设置表头
	// 设置单元格A1查询时间
	f.SetCellValue("Sheet1", "A1", "查询时间")
	// 设置单元格B1查询语句
	f.SetCellValue("Sheet1", "B1", "查询语句")
	// 设置单元格C1查询数据
	f.SetCellValue("Sheet1", "C1", "查询数据")

	// 执行查询
	for i := 0; i < eStep; i++ {
		endTime := i + 1
		fmt.Println("endTime:", endTime)

		result, warnings, err := Papi.QueryRange(context.Background(), Step_query, v1.Range{
			Start: start,
			End:   start.Add(time.Duration(10*endTime) * time.Second),
			Step:  time.Second * 10, // 计步 30s 做一次间隔
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
			fmt.Println("注意：当前查询时间段指标没有数据!")
			return
		} else {
			for _, sample := range matrix {
				if len(sample.Values) > 0 {
					value := sample.Values[len(sample.Values)-1].Value
					fmt.Printf("查询时间：%v\n查询指标：%s 查询数据：%s\n", start.Add(time.Duration(10*endTime)*time.Second), sample.Metric, value)
					index := strconv.Itoa(i + 2)
					fmt.Println("index:", index)
					// 查询时间
					f.SetCellValue("Sheet1", "A"+index, start.Add(time.Duration(10*endTime)*time.Second))
					// 设置单元格A1的值
					f.SetCellValue("Sheet1", "B"+index, sample.Metric)
					// 设置单元格B1的值
					f.SetCellValue("Sheet1", "C"+index, value)
					// 日志调用
					//logs.WithFields(sample.Metric, value)
				}
			}
		}
	}

	// 保存文件
	err = f.SaveAs("demo.xlsx")
	if err != nil {
		fmt.Println("保存文件出错：", err)
		return
	}

	fmt.Println("Excel文件创建并保存成功。")
}
