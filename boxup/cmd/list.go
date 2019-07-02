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
	"fmt"

	"github.com/Chris-Greaves/boxup/boxup/stub"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all boxes on the Server",
	Long:  `This will get all the currently active boxes hosted by the Server.`,
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

		err = client.List()
		if err != nil {
			fmt.Printf("Error listing boxes: %v", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
