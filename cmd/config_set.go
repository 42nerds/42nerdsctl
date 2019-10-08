/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

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
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// configSetCmd represents the configSet command
var configSetCmd = &cobra.Command{
	Use:                   "set KEY VALUE",
	Short:                 "Set a value in the config, eg. `42nerdsctl config set <key> <value>`",
	Long:                  `If no config file is given with --config this will write the config to $HOME/.42nerdsctl.yaml`,
	Args:                  cobra.ExactArgs(2),
	Example:               "42nerdsctl config set email info@42nerds.com",
	DisableFlagsInUseLine: true,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Setting %s: %s\n", args[0], args[1])
		viper.Set(args[0], args[1])
		err := viper.WriteConfig()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

	},
}

func init() {
	configCmd.AddCommand(configSetCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// configSetCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// configSetCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
