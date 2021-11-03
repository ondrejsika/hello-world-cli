package ahoj

import (
	"fmt"

	"github.com/ondrejsika/hello-world-cli/cmd/say"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:     "ahoj",
	Short:   "Say: Ahoj!",
	Aliases: []string{"a"},
	Args:    cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		fmt.Println("Ahoj!")
	},
}

func init() {
	say.Cmd.AddCommand(Cmd)
}
