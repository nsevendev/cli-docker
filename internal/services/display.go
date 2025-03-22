package services

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func UpLine(){
	fmt.Println("")
}

func DisplayWithSpaceUpDown(callback func()){
	UpLine()
	callback()
	UpLine()
}


func DisplayMessageForCommandHelp() {
	UpLine()
	fmt.Printf("ℹ️  Si vous avez un binaire :\n")
	fmt.Printf("📌 Tapez %s`./ns --help`%s ou %s`./ns -h`%s pour voir la liste des commandes disponibles.\n", GREEN, RESET, GREEN, RESET)
	UpLine()
	fmt.Printf("ℹ️  Si vous utiliser le CLI :\n")
	fmt.Printf("📌 Tapez %s`ns --help`%s ou %s`ns -h`%s pour voir la liste des commandes disponibles.\n", GREEN, RESET, GREEN, RESET)
}

func DisplayMessageForHelpCommand() {
	DisplayWithSpaceUpDown(func() {
		fmt.Printf("%s %-10s %s\n", CYAN, "🚀 Voir les options de la commande :", RESET)
		UpLine()
		fmt.Printf("%s %-10s %s: %sAffiche les options de la commandes%s", GREEN, "-h [command], --help [command]", RESET, YELLOW, RESET)
	})
}

func DisplayCommandsOfCli(cmd *cobra.Command) {
	for _, command := range cmd.Commands() {
		if command.Use == "completion" || command.Use == "help [command]" {
			continue
		}
		fmt.Printf("%s  %-10s %s: %s%s%s\n", GREEN, command.Use, RESET, YELLOW, command.Short, RESET)
	}
}

func DisplayFlagForCommand(cmd *cobra.Command) {
	cmd.Flags().VisitAll(func(flag *pflag.Flag) {
		fmt.Printf("  %s-%s%s, %s--%s%s : %s\n",
			GREEN, flag.Shorthand, RESET,
			GREEN, flag.Name, RESET,
			YELLOW+flag.Usage+RESET,
		)
	})
}

func DisplayFlagPersitForCommand(cmd *cobra.Command) {
	cmd.PersistentFlags().VisitAll(func(flag *pflag.Flag) {
		fmt.Printf("  %s-%s%s, %s--%s%s : %s\n",
			GREEN, flag.Shorthand, RESET,
			GREEN, flag.Name, RESET,
			YELLOW+flag.Usage+RESET,
		)
	})
}