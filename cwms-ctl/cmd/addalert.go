// Copyright © 2022 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"cwms-ctl/controller"
	"cwms-ctl/model"
	"github.com/spf13/cobra"
	"strconv"
)

// addalertCmd represents the addalert command
var addalertCmd = &cobra.Command{
	Use: `addalert 告警名称 告警字段(如namespace,cluster 等) 告警字段值(如 dev namespace,devk8s cluster 等)
  某字段指定为空值如下示例:
	例子: cwms-ctl addalert test namespace dev
	解释: 对 test 告警规则添加告警信息,这里是匹配 namespace=dev`,
	Short: "该子命令用于新增告警字段",
	Long:  `该子命令用于新增告警字段`,
	Run: func(cmd *cobra.Command, args []string) {
		//		if len(args) == 0 || len(args) != 3 {
		//			fmt.Println(`语法错误:
		//查看帮助: cwms-ctl addalert -h `)
		//		} else {
		Alert(args)
		//}

	},
}

func init() {
	RootCmd.AddCommand(addalertCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addalertCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addalertCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func Alert(args []string) {
	// 获取 id
	id, _ := strconv.Atoi(controller.QueryWarn(args[0]))

	payload := model.NewRuleDetail()
	payload.RuleId = id
	payload.WarnField = args[1]
	payload.WarnValue = args[2]

	controller.QueryWarnRuleDetail(controller.QueryWarn(args[0]), payload, args[1])

}
