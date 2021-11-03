package cmd

import (
	_ "github.com/ondrejsika/hello-world-cli/cmd/ahoj"
	_ "github.com/ondrejsika/hello-world-cli/cmd/hello"
	"github.com/ondrejsika/hello-world-cli/cmd/root"
	_ "github.com/ondrejsika/hello-world-cli/cmd/say"
	_ "github.com/ondrejsika/hello-world-cli/cmd/sum"
	"github.com/spf13/cobra"
)

func Execute() {
	cobra.CheckErr(root.RootCmd.Execute())
}
