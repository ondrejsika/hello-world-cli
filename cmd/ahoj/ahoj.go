package ahoj

import (
	"fmt"

	"github.com/ondrejsika/hello-world-cli/cmd/root"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:     "ahoj",
	Short:   "Say: Ahoj!",
	Aliases: []string{"h"},
	Args:    cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		fmt.Println("Ahoj!")
	},
}

func init() {
	root.RootCmd.AddCommand(Cmd)
}
