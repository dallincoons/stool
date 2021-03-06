/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
	"surgio-tools/stool"
)

// explainCmd represents the explain command
var explainCmd = &cobra.Command{
	Use:   "explain_parents",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		view_root := viper.GetString("views.root")
		view_name, _ := cmd.Flags().GetString("view")
		show_parents, _ := cmd.Flags().GetBool("file-names")

		explainer := stool.GetExplainer(filepath.Join(view_root, "resources", "views"))

		if show_parents {
			showParents(explainer, view_name)

			return
		}

		variables := explainer.CollectVariablesFromParents(view_name)

		for v, count := range variables {
			fmt.Fprintf(os.Stdout, "%-2d %-8s\n", count, v)
		}
	},
}

func showParents(explainer stool.ViewExplainer, view_name string) {
	parents := explainer.CollectParentsFrom(view_name)

	for parent, _ := range parents {
		fmt.Println(parent)
	}
}

func init() {
	explainCmd.Flags().String("view", "", "specify name of view to explain")
	explainCmd.MarkFlagRequired("view")

	explainCmd.Flags().Bool("file-names", false, "specify whether to show all parent view names")

	rootCmd.AddCommand(explainCmd)
}
