package cmd

import (
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var nodetach bool

func executeShellCommandUp() error {

	detach := "-d"

	if nodetach {
		detach = ""
	}

	cmd := exec.Command("sh", "-c", "docker compose up " + detach)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

var upContainerCmd = &cobra.Command{
	Use: "up",
	Short: "🐳 Lance les conteneurs (mode détaché par defaut)",
	Long: "Lance les conteneurs",
	Run: func (cmd *cobra.Command, args []string)  {
		executeShellCommandUp()
	},
}

func init() {
	upContainerCmd.Flags().BoolVarP(&nodetach, "nodetach", "n", false, "no detach (default detach)")

	rootCmd.AddCommand(upContainerCmd)
}