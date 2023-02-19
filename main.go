package main

import (
	"fmt"

	"github.com/renmcc/go-cmdb/cmd"
)

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
