package root

import (
	"github.com/ondrejsika/hello-world-cli/version"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "hello-world-cli",
	Short: "Hello World CLI, " + version.Version,
}
