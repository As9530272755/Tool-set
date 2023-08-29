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
	node          string
	nodeStartTime string
	nodeEndTime   string
)

// nodeCmd represents the node command
var nodeCmd = &cobra.Command{
	Use:   "node",
	Short: `该命令指定查询 node 内置相关语句`,
	Long: `该命令指定查询 node 相关语句包含以下功能：
		1.输入 node 名查询指定 node 是否存活
		2.输入 node 名查询指定 node cpu 使用率百分比
		3.输入 node 名查询指定 node mem 使用率百分比
  示例: promctl node -Flags --n node:9100 --s 起始时间 --e 结束时间`,
	Run: func(cmd *cobra.Command, args []string) {
		UpStatus, _ := cmd.Flags().GetBool("up")
		if UpStatus {
			NodeUP(node, nodeStartTime, nodeEndTime)
		}

		CpuStatus, _ := cmd.Flags().GetBool("cpu")
		if CpuStatus {
			NodeCpu(node, nodeStartTime, nodeEndTime)
		}

		MemStatus, _ := cmd.Flags().GetBool("mem")
		if MemStatus {
			NodeMem(node, nodeStartTime, nodeEndTime)
		}

	},
}

func init() {
	rootCmd.AddCommand(nodeCmd)
	nodeCmd.Flags().BoolP("up", "u", false, "这是查询 node up 状态子命令")
	nodeCmd.Flags().BoolP("cpu", "c", false, "这是查询 node cpu 使用率子命令")
	nodeCmd.Flags().BoolP("mem", "m", false, "这是查询 node 内存使用率子命令")

	nodeCmd.Flags().StringVarP(&node, "n", "", "", "查询 node (instance:9100 或 hostname:9100)")
	nodeCmd.Flags().StringVarP(&nodeStartTime, "s", "", "", "输入查询起始时间格式如：(2023-01-01 00:00:00)")
	nodeCmd.Flags().StringVarP(&nodeEndTime, "e", "", "", "输入查询结束时间格式如：(2023-01-01 00:00:00)")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// nodeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// nodeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func NodeUP(node, nodeStartTime, nodeEndTime string) {
	if node == "" || nodeStartTime == "" || nodeEndTime == "" {
		fmt.Println("语法错误：请输入 promctl node -h 查看帮助")
		return
	}
	controller.Node_UP_Query(node, nodeStartTime, nodeEndTime)
}

func NodeCpu(node, nodeStartTime, nodeEndTime string) {
	if node == "" || nodeStartTime == "" || nodeEndTime == "" {
		fmt.Println("语法错误：请输入 promctl node -h 查看帮助")
		return
	}
	controller.Node_CPU_Query(node, nodeStartTime, nodeEndTime)
}

func NodeMem(node, nodeStartTime, nodeEndTime string) {
	if node == "" || nodeStartTime == "" || nodeEndTime == "" {
		fmt.Println("语法错误：请输入 promctl node -h 查看帮助")
		return
	}
	controller.Node_Mem_Query(node, nodeStartTime, nodeEndTime)
}
