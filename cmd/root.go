package cmd

import (
	"fmt"

	"github.com/renmcc/go-cmdb/version"
	"github.com/spf13/cobra"
)

var vers bool

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "cmdb-api",
	Short: "资源管理系统",
	Long:  "资源管理系统",
	RunE: func(cmd *cobra.Command, args []string) error {
		if vers {
			fmt.Println(version.FullVersion())
			return nil
		}
		return nil
	},
}

func init() {
	RootCmd.PersistentFlags().BoolVarP(&vers, "version", "v", false, "print cmdb-api version")
}
