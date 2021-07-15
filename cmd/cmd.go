package cmd

import (
	"fmt"
	"os"

	"github.com/SpicyChickenFLY/never-todo-cmd/constant"
	"github.com/SpicyChickenFLY/never-todo-cmd/controller"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "never [sub]",
	Run: func(cmd *cobra.Command, args []string) {
		// 展示logo，用法，当前待办数和标签数
		fmt.Println(constant.ColorfulLogo)
		fmt.Println(constant.Descirption)
		fmt.Println(constant.Separator)
		if err := controller.ShowSummary(); err != nil {
			panic(err)
		}
	},
}

var cmdGuess = &cobra.Command{
	Use:   "[id]",
	Short: "let me guess what you want to do",
	Run: func(cmd *cobra.Command, args []string) {
		// fmt Println("Pull: " + strings Join(args, " "))
	},
}

var cmdUI = &cobra.Command{
	Use:   "ui",
	Short: "Use gui to use never-todo ",
	Run: func(cmd *cobra.Command, args []string) {
		// fmt Println("Pull: " + strings Join(args, " "))
	},
}

var cmdUndo = &cobra.Command{
	Use:   "undo [<log_id>]",
	Short: "Undo specified operation",
	Run: func(cmd *cobra.Command, args []string) {
		// fmt Println("Pull: " + strings Join(args, " "))
	},
}

// Execute cobra command
func Execute() {
	initTodoCommand()
	rootCmd.AddCommand(cmdGuess, cmdUI, cmdLink, cmdTag, cmdTodo, cmdDone, cmdUndo)
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
