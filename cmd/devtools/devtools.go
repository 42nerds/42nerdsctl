// Copyright © 2019 Julian Kaffke <julian@42nerds.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package devtools

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var cfgFile string

// DevToolsCmd represents the base command when called without any subcommands
var DevToolsCmd = &cobra.Command{
	Use:   "devtools",
	Short: "Developer tools for convenience",
	Long:  `TODO: tbd`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the DevToolsCmd.
func Execute() {
	if err := DevToolsCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
