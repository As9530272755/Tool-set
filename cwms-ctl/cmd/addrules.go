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
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

// addrulesCmd represents the addrules command
var addrulesCmd = &cobra.Command{
	Use: `addrules 告警名称 告警优先级 邮箱(多用户 "张三,李四") 短信 企微 抑制时段 可发送告警数量
  某字段指定为空值如下示例:
	例子: cwms-ctl addrules 告警名称 告警优先级 邮箱 "" "" 抑制时段 可发送告警数量
	解释: 将用户短信和用户企微设置为空值
  单条告警规则多用户示例:
	例子: cwms-ctl addrules 告警名称 告警优先级 "张三,李四" "张三,李四" "张三,李四" 抑制时段 可发送告警数量
	解释: 设置邮箱短信和企微为 张三、李四
  命令说明:
	(默认必须使用 7 个参数分别为顺序分别是：告警名称, 告警优先级,邮箱(多用户创建方式 "张桂元,张禹"),短信,企微,抑制时段,可发送告警数量.但是可以通过子参数进行单独创建)`,
	Short: `该子命令用于创建对应的告警规则`,
	Long:  `该子命令用于创建对应的告警规则`,
	Run: func(cmd *cobra.Command, args []string) {
		email, _ := cmd.Flags().GetBool("email")
		if email {
			if len(args) == 0 || len(args) != 5 {
				fmt.Println(`语法错误:
查看帮助: cwms-ctl addrules -h `)
				return
			}
			REmail(args)
		} else if message, _ := cmd.Flags().GetBool("message"); message {
			if len(args) == 0 || len(args) != 5 {
				fmt.Println(`语法错误:
查看帮助: cwms-ctl addrules -h `)
				return
			}
			RMessage(args)
		} else if wechat, _ := cmd.Flags().GetBool("wechat"); wechat {
			if len(args) == 0 || len(args) != 5 {
				fmt.Println(`语法错误:
查看帮助: cwms-ctl addrules -h `)
				return
			}
			RWchat(args)
		} else {
			if len(args) == 0 || len(args) != 7 {
				fmt.Println(`语法错误:
查看帮助: cwms-ctl addrules -h `)
				return
			}
			DefaultRules(args)
		}
	},
}

func init() {
	RootCmd.AddCommand(addrulesCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addrulesCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	addrulesCmd.Flags().BoolP("email", "e", false, "默认创建规则名,邮箱字段;短信,企微 字段为空\n用法: cwms-ctl addrules -e 规则名称 邮箱 优先级 抑制时段 可发送告警 ")
	addrulesCmd.Flags().BoolP("message", "m", false, "默认创建规则名,短信字段;邮箱,企微 字段为空\n用法: cwms-ctl addrules -m 规则名称 短信 优先级 抑制时段 可发送告警")
	addrulesCmd.Flags().BoolP("wechat", "w", false, "默认创建规则名,企微字段;短信,邮箱 字段为空\n用法: cwms-ctl addrules -w 规则名称 企微 优先级 抑制时段 可发送告警")
}

func RuleApi() (api string) {
	// Api 接口
	web := controller.Config()
	paht := `/warnrule/update`

	return web + paht
}

func DefaultRules(args []string) {
	// 将 string 转为 init 来定义优先级
	warn_priority, _ := strconv.Atoi(args[1])
	// 定义抑制时段
	restrain_stage, _ := strconv.Atoi(args[5])
	// 定义可发送告警数量
	restrain_max, _ := strconv.Atoi(args[6])

	payload := model.NewAlarmRules()
	payload.Warn_name = args[0]
	payload.Warn_priority = warn_priority
	payload.Email_receiver = args[2]
	payload.Message_receiver = args[3]
	payload.Wechat_receiver = args[4]
	payload.Restrain_stage = restrain_stage
	payload.Restrain_max = restrain_max

	controller.Put(payload, RuleApi())
}

func REmail(args []string) {
	// 将 string 转为 init 来定义优先级
	warn_priority, _ := strconv.Atoi(args[2])
	// 定义抑制时段
	restrain_stage, _ := strconv.Atoi(args[3])
	// 定义可发送告警数量
	restrain_max, _ := strconv.Atoi(args[4])

	payload := model.NewAlarmRules()
	payload.Warn_name = args[0]
	payload.Email_receiver = args[1]
	payload.Warn_priority = warn_priority
	payload.Restrain_stage = restrain_stage
	payload.Restrain_max = restrain_max

	controller.Put(payload, RuleApi())
}

func RMessage(args []string) {
	// 将 string 转为 init 来定义优先级
	warn_priority, _ := strconv.Atoi(args[2])
	// 定义抑制时段
	restrain_stage, _ := strconv.Atoi(args[3])
	// 定义可发送告警数量
	restrain_max, _ := strconv.Atoi(args[4])

	payload := model.NewAlarmRules()
	payload.Warn_name = args[0]
	payload.Message_receiver = args[1]
	payload.Warn_priority = warn_priority
	payload.Restrain_stage = restrain_stage
	payload.Restrain_max = restrain_max

	controller.Put(payload, RuleApi())
}

func RWchat(args []string) {
	// 将 string 转为 init 来定义优先级
	warn_priority, _ := strconv.Atoi(args[2])
	// 定义抑制时段
	restrain_stage, _ := strconv.Atoi(args[3])
	// 定义可发送告警数量
	restrain_max, _ := strconv.Atoi(args[4])

	payload := model.NewAlarmRules()
	payload.Warn_name = args[0]
	payload.Wechat_receiver = args[1]
	payload.Warn_priority = warn_priority
	payload.Restrain_stage = restrain_stage
	payload.Restrain_max = restrain_max

	controller.Put(payload, RuleApi())
}
