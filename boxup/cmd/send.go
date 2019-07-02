// Copyright © 2018 Christopher Greaves <cjgreaves97@hotmail.co.uk>
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
	"errors"
	"fmt"

	"github.com/Chris-Greaves/boxup/boxup/stub"
	"github.com/spf13/cobra"
)

// sendCmd represents the get command
var sendCmd = &cobra.Command{
	Use:   "send [path]",
	Short: "Send a file to be stored on the server",
	Long:  `This will stream the contents of a file to the server and save it as a box for others to pull down.`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("at least one argument is required")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		host := cmd.Flag("host").Value.String()
		port := cmd.Flag("port").Value.String()
		url := fmt.Sprintf("%v:%v", host, port)
		fmt.Printf("Connecting to %v\n", url)

		client, err := stub.New(url)
		if err != nil {
			fmt.Print(err)
			return
		}

		fmt.Printf("Attempting to get %v from server\n", args[0])
		err = client.Send(args[0])
		if err != nil {
			fmt.Printf("Error getting box %v: %v", args[0], err)
		}
	},
}

func init() {
	rootCmd.AddCommand(sendCmd)
}
