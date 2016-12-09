package container

import (
	"fmt"

	"github.com/Devatoria/rockedcli/client"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List existing containers",
	Run: func(cmd *cobra.Command, args []string) {
		containers, err := client.Get().ListContainers()
		if err != nil {
			panic(err)
		}

		for _, container := range containers {
			fmt.Println(container)
		}
	},
}
