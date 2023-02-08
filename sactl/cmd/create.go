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
	"fmt"
	"sactl/controller"
	"sactl/model"

	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "该子命令用于创建 SA",
	Long: `验证示例：
sactl create -n namespace sa`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 || len(args) != 2 {
			fmt.Println(`语法错误:
查看帮助: sactl create -h `)
			return
		}
		fstatus, _ := cmd.Flags().GetBool("namespace")
		if fstatus {
			Create_Sa_Rolebind(args)
		} else {
			fmt.Println(`语法错误:
查看帮助: sactl create -h `)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(createCmd)
	createCmd.Flags().BoolP("namespace", "n", false, "指定 namespace 创建 SA,实现 rolebind")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func Create_Sa_Rolebind(args []string) {
	newSA := model.NewSa(args[0], args[1], "")
	newRoleBind := model.NewRoleBind(args[0], args[1])

	controller.Create_SA(newSA)
	controller.Role_Bind(newRoleBind)
}
