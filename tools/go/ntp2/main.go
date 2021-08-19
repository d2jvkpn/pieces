package main

import (
	"os"

	"ntp/cmd"

	"github.com/d2jvkpn/pieces/pkg/go/misc"
	"github.com/spf13/cobra"
)

func init() {
	if os.Getenv("TZ") == "" {
		os.Setenv("TZ", "Asia/Shanghai")
	}

	os.Setenv("APP_Name", "ntp")
	misc.SetLogRFC3339()
}

func main() {
	rootCmd := &cobra.Command{Use: "ntp"}

	rootCmd.AddCommand(cmd.NewServerCmd("server"))
	rootCmd.AddCommand(cmd.NewClientCmd("client"))

	rootCmd.Execute()
}
