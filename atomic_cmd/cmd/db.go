/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

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
	"atomic/atomic_model/blog"
	"atomic/atomic_model/category"
	"atomic/atomic_model/collection"
	"atomic/atomic_model/comment"
	"atomic/atomic_model/user"
	"atomic/atomic_store"
	"context"
	"github.com/spf13/cobra"
)

var (
	cg = "mysql"
)

// dbCmd represents the db command
var dbCmd = &cobra.Command{
	Use:   "db",
	Short: "自动建表",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if cg == "mysql" {
			mysql := &atomic_store.Mysql{}
			db, err := atomic_store.DefaultDatabase(context.Background(), mysql)
			if err != nil {
				panic(err)
			}
			err = db.AutoMigrate(
				// 用户
				&user.User{},
				// 博客
				&blog.CommonBlog{},
				// 目录
				&category.Category{},
				// 评论
				&comment.Comment{},
				// 收藏
				&collection.Collection{},
			)
			if err != nil {
				panic(err)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(dbCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// dbCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// dbCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	dbCmd.Flags().StringVarP(&cg, "category", "c", "mysql", "迁移数据库")
}
