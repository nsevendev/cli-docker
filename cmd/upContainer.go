package cmd

import "github.com/spf13/cobra"

var upContainerCmd = &cobra.Command{
	Use: "up",
	Short: "🐳 Lance les conteneurs",
	Long: "Lance les conteneurs",
	Run: func (cmd *cobra.Command, args []string)  {
		
	},
}

func init() {
	rootCmd.AddCommand(upContainerCmd)
}