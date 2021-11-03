package say

import (
	"github.com/ondrejsika/hello-world-cli/cmd/root"
	"github.com/spf13/cobra"
)

var Cmd = &cobra.Command{
	Use:   "say",
	Short: "Say something",
}

func init() {
	root.RootCmd.AddCommand(Cmd)
}
