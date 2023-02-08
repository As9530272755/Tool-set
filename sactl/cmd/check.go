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
	"github.com/spf13/cobra"
	"sactl/controller"
	"sactl/model"
)

// checkCmd represents the check command
var checkCmd = &cobra.Command{
	Use:   `check`,
	Short: "该子命令用于验证 SA",
	Long: `验证示例：
sactl check -n namespace sa api`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 || len(args) != 3 {
			fmt.Println(`语法错误:
查看帮助: sactl check -h `)
			return
		}
		fstatus, _ := cmd.Flags().GetBool("namespace")
		if fstatus {
			appoint_NS(args)
		} else {
			check_SA(args)
		}
	},
}

func init() {
	rootCmd.AddCommand(checkCmd)
	checkCmd.Flags().BoolP("namespace", "n", false, "指定 namespace 验证 SA,默认情况下访问 default NS")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// checkCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// checkCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func check_SA(args []string) {
	newSA := model.NewSa("default", args[0], args[1])

	controller.Get_Secrets(newSA)
	controller.Get_Token(newSA)
	controller.Get_RS(newSA)
}

func appoint_NS(args []string) {
	newSA := model.NewSa(args[0], args[1], args[2])
	controller.Get_Secrets(newSA)
	controller.Get_Token(newSA)
	controller.Get_RS(newSA)
}
