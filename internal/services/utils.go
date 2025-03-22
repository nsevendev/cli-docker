package services

import (
	"fmt"
	"os"
	"os/exec"
)

func ExecuteShellCommand(command string) error {
	cmd := exec.Command("sh", "-c", command)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func QuestionStartCommand(question string) (string) {
	fmt.Print("\n✅ " + question + " (y/N) : ")

	var response string
	fmt.Scanln(&response)
	
	return response
}

func DisplayCommandsForExecute(commands *[]string) {
	fmt.Println(CYAN + "📌 Commandes à exécuter :" + RESET)
	for _, cmd := range *commands {
		fmt.Println("  " + cmd)
	}
}