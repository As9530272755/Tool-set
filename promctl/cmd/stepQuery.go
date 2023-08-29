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
	Step_startTime string
	Step_number    string
	Step_query     string
)

// stepQueryCmd represents the stepQuery command
var stepQueryCmd = &cobra.Command{
	Use:   "stepQuery",
	Short: "基于步长做查询(待开发)",
	Long:  `基于步长做查询`,
	Run: func(cmd *cobra.Command, args []string) {
		StepQuery(Step_query, Step_startTime, Step_number)
	},
}

func init() {
	rootCmd.AddCommand(stepQueryCmd)

	stepQueryCmd.Flags().StringVarP(&Step_query, "q", "", "", "查询 promQL 语句")
	stepQueryCmd.Flags().StringVarP(&Step_startTime, "s", "", "", "起始时间")
	stepQueryCmd.Flags().StringVarP(&Step_number, "t", "", "", "查询次数以 10s 为单位")
}

// 调用 Prometheus 查询函数
func StepQuery(Step_query, Step_startTime, Step_number string) {
	// 如果用户少输入一个指标直接提示帮助
	if Step_query == "" || Step_startTime == "" || Step_number == "" {
		fmt.Println("语法错误：请输入 promctl query -h 查看帮助")
		return
	}
	// 传入输入 query 和 startTime and endTime
	controller.Step_Query(Step_query, Step_startTime, Step_number)
}
