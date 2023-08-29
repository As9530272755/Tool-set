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
	pod          string
	popStartTime string
	podEndTime   string
)

// podCmd represents the pod command
var podCmd = &cobra.Command{
	Use:   "pod",
	Short: `该命令指定查询 pod 内置相关语句`,
	Long: `该命令指定查询 pod 相关语句包含以下功能：
		1.输入 pod 名查询指定 pod 状态
		2.输入 pod 名查询指定 pod cpu 的运行态使用情况
		3.输入 pod 名查询指定 pod mem 的运行态使用情况
  示例: promctl pod -Flags --n podName --s 起始时间 --e 结束时间`,
	Run: func(cmd *cobra.Command, args []string) {
		PodStatus, _ := cmd.Flags().GetBool("up")
		if PodStatus {
			Pod_State(pod, popStartTime, podEndTime)
		}

		PodCpu, _ := cmd.Flags().GetBool("cpu")
		if PodCpu {
			Pod_Cpu(pod, popStartTime, podEndTime)
		}

		PodMem, _ := cmd.Flags().GetBool("mem")
		if PodMem {
			Pod_Mem(pod, popStartTime, podEndTime)
		}

	},
}

func init() {
	rootCmd.AddCommand(podCmd)
	podCmd.Flags().BoolP("up", "u", false, "查询 pod 否存活子命令")
	podCmd.Flags().BoolP("cpu", "c", false, "查询 pod cpu 使用率子命令")
	podCmd.Flags().BoolP("mem", "m", false, "查询 pod 内存使用率子命令")

	podCmd.Flags().StringVarP(&pod, "n", "", "", "输入指定查询 Pod 名称进行查询")
	podCmd.Flags().StringVarP(&popStartTime, "s", "", "", "输入查询起始时间格式(2023-01-01 00:00:00)")
	podCmd.Flags().StringVarP(&podEndTime, "e", "", "", "输入查询结束时间格式(2023-01-01 00:00:00)")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// podCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// podCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func Pod_State(pod, popStartTime, podEndTime string) {

	if pod == "" || popStartTime == "" || podEndTime == "" {
		fmt.Println("语法错误：请输入 promctl pod -h 查看帮助")
		return
	}
	controller.Pod_State(pod, popStartTime, podEndTime)
}

func Pod_Cpu(pod, popStartTime, podEndTime string) {

	if pod == "" || popStartTime == "" || podEndTime == "" {
		fmt.Println("语法错误：请输入 promctl pod -h 查看帮助")
		return
	}
	controller.Pod_Cpu(pod, popStartTime, podEndTime)
}

func Pod_Mem(pod, popStartTime, podEndTime string) {

	if pod == "" || popStartTime == "" || podEndTime == "" {
		fmt.Println("语法错误：请输入 promctl pod -h 查看帮助")
		return
	}
	controller.Pod_Mem(pod, popStartTime, podEndTime)
}
