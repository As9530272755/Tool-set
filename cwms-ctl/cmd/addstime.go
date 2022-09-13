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
	"github.com/spf13/cobra"
	"strconv"
)

// addstimeCmd represents the addstime command
var addstimeCmd = &cobra.Command{
	Use: `addstime 
	例子：cwms-ctl addstime -o test "2022-09-08 18:00:00" "2022-09-09 09:00:00"
	解释：指定给 test 告警规则添加单次告警静默时间，告警静默时间为当天 18 点到第二天 9 点`,
	Short: "该子命令用于新增告警静默时间",
	Long:  `该子命令用于新增告警静默时间`,
	Run: func(cmd *cobra.Command, args []string) {
		once, _ := cmd.Flags().GetBool("once")
		if once {
			OnceStime(args)
		} else if everyday, _ := cmd.Flags().GetBool("everyday"); everyday {
			EverydayStime(args)
		} else {
			fmt.Println("失败")
			return
		}
	},
}

func init() {
	RootCmd.AddCommand(addstimeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addstimeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	addstimeCmd.Flags().BoolP("once", "o", false, `单次静默: cwms-ctl addstime -o WarnRuleName BeginTime EndTime `)
	addstimeCmd.Flags().BoolP("everyday", "e", false, "每天静默: cwms-ctl addstime -e WarnRuleName BeginTime EndTime ")

}

func stimeApi() (api string) {
	// Api 接口
	web := controller.Config()
	paht := `/stoptime/update`

	return web + paht
}

func OnceStime(args []string) {

	// 获取id
	id, _ := strconv.Atoi(controller.QueryWarn(args[0]))

	payload := model.NewStopTime(args[0], args[1], args[2], id, 0)

	controller.Put(payload, stimeApi())
}

func EverydayStime(args []string) {
	// 获取id
	id, _ := strconv.Atoi(controller.QueryWarn(args[0]))

	payload := model.NewStopTime(args[0], args[1], args[2], id, 1)

	controller.Put(payload, stimeApi())
}
