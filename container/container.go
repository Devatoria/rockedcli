package container

import (
	"github.com/spf13/cobra"
)

var ContainerCmd = &cobra.Command{
	Use:   "container",
	Short: "Manage Docker containers",
}

func init() {
	ContainerCmd.AddCommand(listCmd)
}
