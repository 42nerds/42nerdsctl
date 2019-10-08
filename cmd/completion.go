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

package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// createCmd represents the create command
var completionCmd = &cobra.Command{
	Use:   "completion",
	Short: "Generates bash completion scripts",
	Long: `
Output shell completion code for bash. The shell code must be evaluated
to provide interactive completion of 42nerdsctl commands.  This can be done by sourcing it from the
.bash_profile.

Examples:
	# Installing bash completion on macOS using homebrew
	## If running Bash 3.2 included with macOS
	brew install bash-completion
	## or, if running Bash 4.1+
	brew install bash-completion@2
	## If you've installed, you may need add the completion to your completion
directory
	42nerdsctl completion > $(brew --prefix)/etc/bash_completion.d/42nerdsctl


	# Installing bash completion on Linux
	## Load the 42nerdsctl completion code for bash into the current shell
	source <(42nerdsctl completion bash)
`,
	Run: func(cmd *cobra.Command, args []string) {
		rootCmd.GenBashCompletion(os.Stdout)
	},
}

func init() {
	rootCmd.AddCommand(completionCmd)
}
