package controller

import (
	"context"
	"demo3/config"
	"demo3/logs"
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
func Node_UP_Query(node, nodeStartTime, nodeEndTime string) {
	layout := "2006-01-02 15:04:05"

	// 上海时区
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println("无效的时区")
		os.Exit(1)
	}

	// 时区转换
	start, _ := time.ParseInLocation(layout, nodeStartTime, loc)
	end, _ := time.ParseInLocation(layout, nodeEndTime, loc)

	// 读取配置文件中的地址
	API, err := config.ParseConfig()
	if err != nil {
		log.Panic(err)
	}

	// Prometheus API 拼接
	promURL := API.Prometheus.Addr + ":" + API.Prometheus.Port

	fmt.Println(start, end)
	fmt.Println("Prometheus 地址：", promURL)

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

	// 传入 node 名称进查询
	nodeNameQL := fmt.Sprintf(`up{instance="%s"}`, node)

	// 执行查询
	result, warnings, err := Papi.QueryRange(context.Background(), nodeNameQL, v1.Range{
		Start: start,
		End:   end,
		Step:  time.Second * 30, // 计步 30s 做一次间隔
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

	// 创建一个新的Excel文件
	f := excelize.NewFile()
	// 设置表头
	// 设置单元格A1查询时间
	f.SetCellValue("Sheet1", "A1", "查询时间")
	// 设置单元格B1查询语句
	f.SetCellValue("Sheet1", "B1", "查询语句")
	// 设置单元格C1查询数据
	f.SetCellValue("Sheet1", "C1", "查询数据")

	// 设置表列 index
	index := 0

	// 处理查询结果
	matrix := result.(model.Matrix)
	if matrix.Len() == 0 {
		fmt.Println("注意：当前查询时间段指标没有数据!")
		return
	} else {
		for _, sample := range matrix {
			if len(sample.Values) > 0 {
				value := sample.Values[len(sample.Values)-1].Value
				fmt.Println("指标：", len(sample.Values))
				fmt.Printf("查询时间：%v\n查询指标：%s 查询数据：%s\n", start, sample.Metric, value)

				//日志调用
				logs.WithFields(sample.Metric, value)
			}

			// 每循环一次表列 +1
			index += 1

			// 数据指标从第二列开始写入,并转换数据类型
			sindex := strconv.Itoa(index + 1)
			fmt.Println("查询指标条数:", sindex)
			// 查询时间
			f.SetCellValue("Sheet1", "A"+sindex, start)
			// 设置单元格的值
			f.SetCellValue("Sheet1", "B"+sindex, nodeNameQL)
			// 设置单元格的值
			f.SetCellValue("Sheet1", "C"+sindex, sample.Values[len(sample.Values)-1].Value)
		}
	}
	// 保存文件
	err = f.SaveAs("Prometheus查询.xlsx")
	if err != nil {
		fmt.Println("保存文件出错：", err)
		return
	}
	fmt.Println("Prometheus查询.xlsx Excel 文件创建并保存成功。")
}

// 基于起始时间和结束时间做全量查询
func Node_CPU_Query(node, nodeStartTime, nodeEndTime string) {
	layout := "2006-01-02 15:04:05"

	// 上海时区
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println("无效的时区")
		os.Exit(1)
	}

	// 时区转换
	start, _ := time.ParseInLocation(layout, nodeStartTime, loc)
	end, _ := time.ParseInLocation(layout, nodeEndTime, loc)

	// 读取配置文件中的地址
	API, err := config.ParseConfig()
	if err != nil {
		log.Panic(err)
	}

	// Prometheus API 拼接
	promURL := API.Prometheus.Addr + ":" + API.Prometheus.Port

	fmt.Println(start, end)
	fmt.Println("Prometheus 地址: ", promURL)

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

	// 传入 node 名称进查询
	nodeNameQL := fmt.Sprintf(`100-(avg(irate(node_cpu_seconds_total{mode="idle",instance="%s"}[5m])) by(instance) *100)`, node)

	// 执行查询
	result, warnings, err := Papi.QueryRange(context.Background(), nodeNameQL, v1.Range{
		Start: start,
		End:   end,
		Step:  time.Second * 30, // 计步 30s 做一次间隔
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

	// 创建一个新的Excel文件
	f := excelize.NewFile()
	// 设置表头
	// 设置单元格A1查询时间
	f.SetCellValue("Sheet1", "A1", "查询时间")
	// 设置单元格B1查询语句
	f.SetCellValue("Sheet1", "B1", "查询语句")
	// 设置单元格C1查询数据
	f.SetCellValue("Sheet1", "C1", "查询数据")

	// 设置表列 index
	index := 0

	// 处理查询结果
	matrix := result.(model.Matrix)
	if matrix.Len() == 0 {
		fmt.Println("注意：当前查询时间段指标没有数据!")
		return
	} else {
		for _, sample := range matrix {
			if len(sample.Values) > 0 {
				value := sample.Values[len(sample.Values)-1].Value
				fmt.Println("指标：", len(sample.Values))
				fmt.Printf("查询时间：%v\n查询指标：%s\n查询数据：%s\n", start, nodeNameQL, value)

				//日志调用
				logs.WithFields(sample.Metric, value)
			}

			// 每循环一次表列 +1
			index += 1

			// 数据指标从第二列开始写入,并转换数据类型
			sindex := strconv.Itoa(index + 1)
			fmt.Println("查询指标条数:", sindex)
			// 查询时间
			f.SetCellValue("Sheet1", "A"+sindex, start)
			// 设置单元格的值
			f.SetCellValue("Sheet1", "B"+sindex, nodeNameQL)
			// 设置单元格的值
			f.SetCellValue("Sheet1", "C"+sindex, sample.Values[len(sample.Values)-1].Value)
		}
	}
	// 保存文件
	err = f.SaveAs("Prometheus查询.xlsx")
	if err != nil {
		fmt.Println("保存文件出错：", err)
		return
	}
	fmt.Println("Prometheus查询.xlsx Excel 文件创建并保存成功。")
}

func Node_Mem_Query(node, nodeStartTime, nodeEndTime string) {
	layout := "2006-01-02 15:04:05"

	// 上海时区
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println("无效的时区")
		os.Exit(1)
	}

	// 时区转换
	start, _ := time.ParseInLocation(layout, nodeStartTime, loc)
	end, _ := time.ParseInLocation(layout, nodeEndTime, loc)

	// 读取配置文件中的地址
	API, err := config.ParseConfig()
	if err != nil {
		log.Panic(err)
	}

	// Prometheus API 拼接
	promURL := API.Prometheus.Addr + ":" + API.Prometheus.Port

	fmt.Println(start, end)
	fmt.Println("Prometheus 地址: ", promURL)

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

	// 传入 node 名称进查询
	nodeNameQL := fmt.Sprintf(`avg(((node_memory_MemTotal_bytes{instance="%s"} - node_memory_MemFree_bytes - node_memory_Buffers_bytes - node_memory_Cached_bytes) / (node_memory_MemTotal_bytes)) * 100) by (instance)`, node)

	// 执行查询
	result, warnings, err := Papi.QueryRange(context.Background(), nodeNameQL, v1.Range{
		Start: start,
		End:   end,
		Step:  time.Second * 30, // 计步 30s 做一次间隔
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

	// 创建一个新的Excel文件
	f := excelize.NewFile()
	// 设置表头
	// 设置单元格A1查询时间
	f.SetCellValue("Sheet1", "A1", "查询时间")
	// 设置单元格B1查询语句
	f.SetCellValue("Sheet1", "B1", "查询语句")
	// 设置单元格C1查询数据
	f.SetCellValue("Sheet1", "C1", "查询数据")

	// 设置表列 index
	index := 0

	// 处理查询结果
	matrix := result.(model.Matrix)
	if matrix.Len() == 0 {
		fmt.Println("注意：当前查询时间段指标没有数据!")
		return
	} else {
		for _, sample := range matrix {
			if len(sample.Values) > 0 {
				value := sample.Values[len(sample.Values)-1].Value
				fmt.Println("指标：", len(sample.Values))
				fmt.Printf("查询时间：%v\n查询指标：%s\n查询数据：%s\n", start, nodeNameQL, value)

				//日志调用
				logs.WithFields(sample.Metric, value)
			}

			// 每循环一次表列 +1
			index += 1

			// 数据指标从第二列开始写入,并转换数据类型
			sindex := strconv.Itoa(index + 1)
			fmt.Println("查询指标条数:", sindex)
			// 查询时间
			f.SetCellValue("Sheet1", "A"+sindex, start)
			// 设置单元格的值
			f.SetCellValue("Sheet1", "B"+sindex, nodeNameQL)
			// 设置单元格的值
			f.SetCellValue("Sheet1", "C"+sindex, sample.Values[len(sample.Values)-1].Value)
		}
	}
	// 保存文件
	err = f.SaveAs("Prometheus查询.xlsx")
	if err != nil {
		fmt.Println("保存文件出错：", err)
		return
	}
	fmt.Println("Prometheus查询.xlsx Excel 文件创建并保存成功。")
}
