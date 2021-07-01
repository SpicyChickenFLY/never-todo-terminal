package cmd

import (
	"github.com/SpicyChickenFLY/never-todo-cmd/logic"
	"github.com/spf13/cobra"
)

var cmdTodo = &cobra.Command{
	Use:   "todo [<id>|<keyword>]",
	Short: "deal with task(s)",
	Run: func(cmd *cobra.Command, args []string) {
		if err := logic.ShowTasks(args); err != nil {
			// cmd.PrintErr(err)
			panic(err)
		}
	},
}

var cmdTodoAdd = &cobra.Command{
	Use:   "add <content> [,<content>]",
	Short: "add new task(s)",
	Run: func(cmd *cobra.Command, args []string) {
		// fmt Println("Pull: " + strings Join(args, " "))
		if err := logic.AddTasks(args); err != nil {
			panic(err)
		}
	},
}

var cmdTodoDel = &cobra.Command{
	Use:   "del <id> [,<id>]",
	Short: "delete task(s)",
	Run: func(cmd *cobra.Command, args []string) {
		// fmt Println("Pull: " + strings Join(args, " "))
	},
}

var cmdTodoSet = &cobra.Command{
	Use:   "[set] <id> <content>",
	Short: "update task(s)",
	Run: func(cmd *cobra.Command, args []string) {
		// fmt Println("Pull: " + strings Join(args, " "))
	},
}

var cmdTag = &cobra.Command{
	Use:   "tag [<id>|<keyword>]",
	Short: "deal with tag(s)",
	Run: func(cmd *cobra.Command, args []string) {
		// fmt Println("Pull: " + strings Join(args, " "))
	},
}

var cmdTagAdd = &cobra.Command{
	Use:   "add <content> [,<content>]",
	Short: "add new tag(s)",
	Run: func(cmd *cobra.Command, args []string) {
		// fmt Println("Pull: " + strings Join(args, " "))
	},
}

var cmdTagDel = &cobra.Command{
	Use:   "del <id> [,<id>]",
	Short: "delete tag(s)",
	Run: func(cmd *cobra.Command, args []string) {
		// fmt Println("Pull: " + strings Join(args, " "))
	},
}

var cmdTagSet = &cobra.Command{
	Use:   "[set] <id> <content>",
	Short: "Update tag(s)",
	Run: func(cmd *cobra.Command, args []string) {
		// fmt Println("Pull: " + strings Join(args, " "))
	},
}

var cmdDone = &cobra.Command{
	Use:   "done <todo_id> [,<todo_id>]",
	Short: "complete task(s)",
	Run: func(cmd *cobra.Command, args []string) {
		// fmt Println("Pull: " + strings Join(args, " "))
	},
}

var cmdLink = &cobra.Command{
	Use:   "link [<todo_id>|<keyword>]",
	Short: "deal with relationship between task(s) and tag(s)",
	Run: func(cmd *cobra.Command, args []string) {
		// fmt Println("Pull: " + strings Join(args, " "))
	},
}

var cmdLinkSet = &cobra.Command{
	Use:   "link <todo_id> [set] <tag_id> [,<tag_id>]",
	Short: "set tag(s) for task(s)",
	Run: func(cmd *cobra.Command, args []string) {
		// fmt Println("Pull: " + strings Join(args, " "))
	},
}

var cmdLinkAdd = &cobra.Command{
	Use:   "link <todo_id> add <tag_id> [,<tag_id>]",
	Short: "add tag(s) for task(s)",
	Run: func(cmd *cobra.Command, args []string) {
		// fmt Println("Pull: " + strings Join(args, " "))
	},
}

var cmdLinkDel = &cobra.Command{
	Use:   "link <todo_id> del <tag_id> [,<tag_id>]",
	Short: "delete tag(s) for task(s)",
	Run: func(cmd *cobra.Command, args []string) {
		// fmt Println("Pull: " + strings Join(args, " "))
	},
}

func initTodoCommand() {
	cmdTodo.AddCommand(cmdTodoAdd, cmdTodoDel, cmdTodoSet)
	cmdTag.AddCommand(cmdTagAdd, cmdTagDel, cmdTagSet)
	cmdLink.AddCommand(cmdLinkAdd, cmdLinkDel, cmdLinkSet)
}
