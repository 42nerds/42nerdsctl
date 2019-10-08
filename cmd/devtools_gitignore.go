/*
Copyright © 2019 Julian Kaffke <julian@42nerds.com>

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
	"io/ioutil"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

var fileName string

// devtoolsGitignoreCmd represents the devtoolsGitignore command
var devtoolsGitignoreCmd = &cobra.Command{
	Use:   "gitignore TEMPLATE1,TEMPLATE2..TEMPLATEn",
	Short: "generates a `.gitignore` file in current directory",
	Long: `This command generates a '.gitignore' file from https://gitignore.io.
You can pass the needed templates in a commaseperated list (without whitespace`,
	Example: "42nerdsctl devtools gitignore macos,git,visualstudiocode,go",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		resp, err := http.Get("https://www.gitignore.io/api/" + args[0])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)

		ioutil.WriteFile(fileName, body, 0755)
	},
}

func init() {
	devtoolsCmd.AddCommand(devtoolsGitignoreCmd)
	devtoolsGitignoreCmd.Flags().StringVarP(&fileName, "filename", "f", ".gitignore", "--filename .custom_gitignore")
}
