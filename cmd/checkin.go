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
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func checkIn() error {
	apiURL := "https://people.zoho.eu/people/api/attendance"
	authToken := viper.GetString("zoho_people_token")
	emailID := viper.GetString("email")
	t := time.Now()
	checkInTime := fmt.Sprintf("%02d-%02d-%d %02d:%02d:%02d",
		t.Day(), t.Month(), t.Year(),
		t.Hour(), t.Minute(), t.Second())
	query := fmt.Sprintf("?authtoken=%s&emailId=%s&checkIn=%s", authToken, emailID, url.QueryEscape(checkInTime))
	resp, err := http.Get(apiURL + query)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	_, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	return nil
}

// createCmd represents the create command
var checkinCmd = &cobra.Command{
	Use:   "checkin",
	Short: "This will check you in",
	Long: `This command is a helper to check you in into Zoho People,

Be aware that you'll need to create and store a API Token in your config file

You can get your token here:

https://accounts.zoho.eu/apiauthtoken/create?SCOPE=zohopeople/peopleapi

Copy and set this with

	$ 42nerdsctl config set zoho_people_token <YOUR TOKEN>

Also set your e-mail:

	$ 42nerdsctl config set email <you>@42nerds.com
`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Checkin ya' in")
		err := checkIn()
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(checkinCmd)
}
