package cmd

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var forceFlag bool

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize AgentKit prompt structure",
	Long: `Initialize creates the .github directory structure with copilot-instructions.md
and prompt templates for the AgentKit workflow.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Implement actual file creation in next commit
		green := color.New(color.FgGreen)
		green.Println("Init command called")
		if forceFlag {
			fmt.Println("Force flag is set")
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
	initCmd.Flags().BoolVarP(&forceFlag, "force", "f", false, "Force overwrite existing files")
}
