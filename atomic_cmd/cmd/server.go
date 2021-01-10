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
	"atomic/atomic_server/atomic_handler"
	"atomic/atomic_server/register"
	"atomic/atomic_store"
	"context"
	"github.com/facebookarchive/inject"
	"github.com/spf13/cobra"
)

var (
	port          = 8090
	insideService = "user"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "启动分布式服务",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		db, err := atomic_store.DefaultDatabase(context.Background(), &atomic_store.Mysql{})
		if err != nil {
			panic(err)
		}
		graph := &inject.Graph{}
		err = graph.Provide(&inject.Object{Value: db})
		if err != nil {
			panic(err)
		}
		if insideService == "user" {
			handler := &atomic_handler.UserHandler{}
			err = graph.Provide(&inject.Object{Value: handler})
			if err != nil {
				panic(err)
			}
			err = graph.Populate()
			if err != nil {
				panic(err)
			}
			register.UserServiceRegister(port, handler)
		}

		if insideService == "blog" {
			handler := &atomic_handler.BlogHandler{}
			err = graph.Provide(&inject.Object{Value: handler})
			if err != nil {
				panic(err)
			}
			err = graph.Populate()
			if err != nil {
				panic(err)
			}
			register.BlogServiceRegistry(port, handler)
		}
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serverCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serverCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	//serverCmd.Flags().StringP("",)

	// 端口号
	serverCmd.Flags().IntVarP(&port, "port", "p", 8090, "服务所占用的端口")
	// 服务名
	serverCmd.Flags().StringVarP(&insideService, "register", "r", "user", "所需要注册的服务")

}
