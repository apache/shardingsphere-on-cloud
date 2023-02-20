/*
* Licensed to the Apache Software Foundation (ASF) under one or more
* contributor license agreements.  See the NOTICE file distributed with
* this work for additional information regarding copyright ownership.
* The ASF licenses this file to You under the Apache License, Version 2.0
* (the "License"); you may not use this file except in compliance with
* the License.  You may obtain a copy of the License at
*
*     http://www.apache.org/licenses/LICENSE-2.0
*
* Unless required by applicable law or agreed to in writing, software
* distributed under the License is distributed on an "AS IS" BASIS,
* WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
* See the License for the specific language governing permissions and
* limitations under the License.
 */

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const (
	host      = "host"
	port      = "port"
	username  = "username"
	password  = "password"
	agentPort = "agent-port"
)

var Root = &cobra.Command{
	Use:   "gs_pitr",
	Short: "PITR tools for openGauss",
	Run: func(cmd *cobra.Command, args []string) {
		host, err := cmd.Flags().GetString(host)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(fmt.Sprintf("flags:host:%s", host))

		port, err := cmd.Flags().GetUint16(port)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(fmt.Sprintf("flags:port:%d", port))

		un, err := cmd.Flags().GetString(username)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(fmt.Sprintf("flags:username:%s", un))

		pw, err := cmd.Flags().GetString(password)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(fmt.Sprintf("flags:password:%s", pw))

		agentPort, err := cmd.Flags().GetUint16(agentPort)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(fmt.Sprintf("flags:agentPort:%d", agentPort))
	},
	CompletionOptions: cobra.CompletionOptions{
		DisableDefaultCmd: true,
		HiddenDefaultCmd:  true,
	},
}

func init() {
	Root.PersistentFlags().StringP(host, "H", "", "shardingsphere proxy server host")
	Root.PersistentFlags().Uint16P(port, "P", 1, "shardingsphere proxy service port")
	Root.PersistentFlags().StringP(username, "u", "", "shardingsphere proxy username")
	Root.PersistentFlags().StringP(password, "p", "", "shardingsphere proxy password")
	Root.PersistentFlags().Uint16P(agentPort, "", 443, "agent server port")

	Root.AddCommand(Backup, Restore, Show)
}
