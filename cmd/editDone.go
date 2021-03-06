package cmd

import (
	"github.com/hypotheticalco/tracker-client/todo"
	"github.com/hypotheticalco/tracker-client/utils"
	"github.com/spf13/cobra"
)

var editDone = &cobra.Command{
	Use:     "editDone",
	Short:   "Open the done file using $EDITOR",
	Args:    cobra.NoArgs,
	Aliases: []string{"ed"},
	Run: func(cmd *cobra.Command, args []string) {
		utils.EditFilePath(todo.DoneFilePath())
	},
}

func init() {
	RootCmd.AddCommand(editDone)
}
