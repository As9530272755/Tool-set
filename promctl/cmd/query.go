/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"demo3/controller"
	"fmt"
	"github.com/spf13/cobra"
)

var (
	allstartTime string
	allendTime   string
	allquery     string
)

// queryCmd represents the query command
var queryCmd = &cobra.Command{

	Use: "query",
	Short: `该命令对应 promQL 全量语句查询	`,
	Long: `该命令对应 promQL 全量语句查询
  示例: promctl query --q "查询语句" --s "起始时间" --e "结束时间"`,
	Run: func(cmd *cobra.Command, args []string) {
		AllQuery(allquery, allstartTime, allendTime)
	},
}

func init() {
	rootCmd.AddCommand(queryCmd)
	// 输入多个 flags 参数
	queryCmd.Flags().StringVarP(&allquery, "q", "", "", "全量查询 promQL 语句")
	queryCmd.Flags().StringVarP(&allstartTime, "s", "", "", "输入查询起始时间格式如：(2023-01-01 00:00:00)")
	queryCmd.Flags().StringVarP(&allendTime, "e", "", "", "输入查询结束时间格式如：(2023-01-01 00:00:00)")

}

// 调用 Prometheus 查询函数
func AllQuery(allquery, allstartTime, allendTime string) {
	// 如果用户少输入一个指标直接提示帮助
	if allquery == "" || allstartTime == "" || allendTime == "" {
		fmt.Println("语法错误：请输入 promctl query -h 查看帮助")
		return
	}
	// 传入输入 query 和 startTime and endTime
	controller.AllQuery(allquery, allstartTime, allendTime)
}
