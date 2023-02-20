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
