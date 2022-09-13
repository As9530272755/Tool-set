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
)

// adduserCmd represents the adduser command
var adduserCmd = &cobra.Command{
	Use: `adduser 用户名 邮件 手机 企微 角色 备注
	命令说明:
			默认必须使用 6 个参数分别为顺序分别是：用户名,邮件,手机,企微,角色,备注说明;但是可以通过子参数对单个字段进行单独创建
  如某字段指定 null 如下示例: 
	例子: cwms-ctl adduser 用户名 邮件 手机 企微 null null 
	解释: 对角色和备注字段定义为 null`,
	Short: `该子命令用于创建用户`,
	Long:  `该子命令用于创建用户`,
	Run: func(cmd *cobra.Command, args []string) {

		name, _ := cmd.Flags().GetBool("name")
		if name {
			if len(args) == 0 || len(args) != 1 {
				fmt.Println(`语法错误:
查看帮助: cwms-ctl adduser -h `)
				return
			}
			User(args)
		} else if email, _ := cmd.Flags().GetBool("email"); email {
			if len(args) == 0 || len(args) != 2 {
				fmt.Println(`语法错误:
查看帮助: cwms-ctl adduser -h `)
				return
			}
			Email(args)
		} else if phone, _ := cmd.Flags().GetBool("phone"); phone {
			if len(args) == 0 || len(args) != 2 {
				fmt.Println(`语法错误:
查看帮助: cwms-ctl adduser -h `)
				return
			}
			Phone(args)
		} else if wechat, _ := cmd.Flags().GetBool("wechat"); wechat {
			if len(args) == 0 || len(args) != 2 {
				fmt.Println(`语法错误:
查看帮助: cwms-ctl adduser -h `)
				return
			}
			Wechat_id(args)
		} else if comment, _ := cmd.Flags().GetBool("comment"); comment {
			if len(args) == 0 || len(args) != 2 {
				fmt.Println(`语法错误:
查看帮助: cwms-ctl adduser -h `)
				return
			}
			Comment(args)
		} else if role, _ := cmd.Flags().GetBool("role"); role {
			if len(args) == 0 || len(args) != 2 {
				fmt.Println(`语法错误:
查看帮助: cwms-ctl adduser -h `)
				return
			}
			Role(args)
		} else {
			if len(args) == 0 || len(args) != 6 {
				fmt.Println(`语法错误:
查看帮助: cwms-ctl adduser -h `)
				return
			}
			DefaultCreateUser(args)
		}

	},
}

func init() {
	RootCmd.AddCommand(adduserCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// adduserCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	adduserCmd.Flags().BoolP("name", "n", false, "输入创建用户名字段;邮箱、电话、企微、角色、备注 默认为 null \n用法: cwms-ctl adduser -n UserName")
	adduserCmd.Flags().BoolP("email", "e", false, "输入用户名和邮箱;电话、企微、角色、备注 默认为 null \n用法: cwms-ctl adduser -e UserName Email")
	adduserCmd.Flags().BoolP("phone", "p", false, "输入用户名和手机号;邮箱、企微、角色、备注 默认为 null \n用法: cwms-ctl adduser -e -p UserName Phone")
	adduserCmd.Flags().BoolP("wechat", "w", false, "输入用户名和企微;邮箱、电话、角色、备注 默认为 null \n用法: cwms-ctl adduser -w UserName Wechat")
	adduserCmd.Flags().BoolP("comment", "c", false, "输入用户名和备注;邮箱、电话、企微、角色 默认为 null \n用法: cwms-ctl adduser -c UserName Comment")
	adduserCmd.Flags().BoolP("role", "r", false, "输入用户名和角色;邮箱、电话、企微、备注 默认为 null \n用法: cwms-ctl adduser -r UserName Role")
}

func UserApi() (api string) {
	// Api 接口
	web := controller.Config()
	paht := `/staff/update`

	return web + paht
}

func DefaultCreateUser(args []string) {

	payload := model.NewUser()

	payload.Name = args[0]
	payload.Email = args[1]
	payload.Phone = args[2]
	payload.Wechat_id = args[3]
	payload.Role = args[4]
	payload.Comment = args[5]

	controller.Put(payload, UserApi())
}

func User(args []string) {
	payload := model.NewUser()
	payload.Name = args[0]

	controller.Put(payload, UserApi())
}

func Email(args []string) {
	payload := model.NewUser()
	payload.Name = args[0]
	payload.Email = args[1]

	controller.Put(payload, UserApi())
}

func Phone(args []string) {
	payload := model.NewUser()
	payload.Name = args[0]
	payload.Phone = args[1]

	controller.Put(payload, UserApi())
}

func Wechat_id(args []string) {
	payload := model.NewUser()
	payload.Name = args[0]
	payload.Wechat_id = args[1]

	controller.Put(payload, UserApi())
}

func Role(args []string) {
	payload := model.NewUser()
	payload.Name = args[0]
	payload.Role = args[1]

	controller.Put(payload, UserApi())
}

func Comment(args []string) {
	payload := model.NewUser()
	payload.Name = args[0]
	payload.Comment = args[1]

	controller.Put(payload, UserApi())
}
