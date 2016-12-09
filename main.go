package main

import (
	"fmt"
	"os"

	"github.com/Devatoria/rockedcli/container"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "rockedcli",
	Short: "rockedcli is a CLI using rocked to get information about your Docker containers without using the daemon",
}

func init() {
	rootCmd.AddCommand(container.ContainerCmd)

	rootCmd.PersistentFlags().StringP("docker", "d", "/var/lib/docker", "Docker home directory")
	viper.BindPFlag("docker", rootCmd.PersistentFlags().Lookup("docker"))
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
