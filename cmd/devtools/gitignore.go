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

	"github.com/spf13/cobra"
)

// createCmd represents the create command
var cmd = &cobra.Command{
	Use:   "gitignore",
	Short: "Generates a gitignore with useful defaults",
	Long:  `TODO: tbd`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("gitignore called")
	},
}

func init() {
	DevToolsCmd.AddCommand(cmd)
}
