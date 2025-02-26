package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"os/exec"
)

func main() {
	var rootCmd = &cobra.Command{Use: "backstage-cli"}

	var upCmd = &cobra.Command{
		Use:   "up",
		Short: "Sobe a aplicação Backstage (Lab da aninha) usando yarn dev",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Subindo Lab da Aninha...")

			// Comando para rodar yarn dev
			yarnDev := exec.Command("yarn", "dev")
			yarnDev.Dir = "adicionar path"     // Definir o diretório do projeto Backstage do lab
			yarnDev.Stdout = cmd.OutOrStdout() // Exibindo a saída no terminal
			yarnDev.Stderr = cmd.ErrOrStderr() // Exibindo erros no terminal

			err := yarnDev.Run()
			if err != nil {
				fmt.Println("Erro ao subir Backstage:", err)
				return
			}
		},
	}

	rootCmd.AddCommand(upCmd)
	rootCmd.Execute()
}
