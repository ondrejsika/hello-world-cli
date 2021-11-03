package sum

import (
	"fmt"

	"github.com/ondrejsika/hello-world-cli/cmd/root"
	"github.com/spf13/cobra"
)

var CmdFlagA int
var CmdFlagB int

var Cmd = &cobra.Command{
	Use:   "sum",
	Short: "Sum of 2 numbers",
	Args:  cobra.NoArgs,
	Run: func(c *cobra.Command, args []string) {
		fmt.Println(CmdFlagA + CmdFlagB)
	},
}

func init() {
	root.RootCmd.AddCommand(Cmd)
	Cmd.PersistentFlags().IntVarP(
		&CmdFlagA,
		"a",
		"a",
		0,
		"1st nuber",
	)
	Cmd.MarkPersistentFlagRequired("a")
	Cmd.PersistentFlags().IntVarP(
		&CmdFlagB,
		"b",
		"b",
		0,
		"2nd nuber",
	)
	Cmd.MarkPersistentFlagRequired("b")
}
