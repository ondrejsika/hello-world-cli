package cmd

import (
	"github.com/ondrejsika/hello-world-cli/cmd/root"
	"github.com/spf13/cobra"
)

func Execute() {
	cobra.CheckErr(root.RootCmd.Execute())
}
