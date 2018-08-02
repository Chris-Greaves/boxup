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

	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get [box names]",
	Short: "Get Box(s) from a BoxUp Server",
	Long: `This command will get the contents of boxs hosted on a BoxUp Server.
	
The files and directories will be place wherever the current working directory is unless specified using the -o flag
Will assume local BoxUp Server unless specifed using --host http://domain.com or --host http://10.0.0.1
To specify a port use --port 5656

Examples:

Get box and output to specific folder:
  boxup get -o C:\output test-folder

Get multiple boxes:
  boxup get folder1 folder2
	
Specify server:
  boxup get --host localhost --port 5950 test-folder`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires at least one arg")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Args:\n")
		for index, arg := range args {
			fmt.Printf("\t%v:%v\n", index, arg)
		}
		fmt.Printf("\n")

		hostFlag := cmd.Flag("host")
		fmt.Printf("host:%v\n", hostFlag.Value.String())
		portFlag := cmd.Flag("port")
		fmt.Printf("port:%v\n", portFlag.Value.String())
		outputFlag := cmd.Flag("output")
		fmt.Printf("output:%v\n", outputFlag.Value.String())

		for _, arg := range args {
			fmt.Printf("\nUnboxing %v\n", arg)

			err := getBox(hostFlag.Value.String(), portFlag.Value.String(), outputFlag.Value.String(), arg)
			if err != nil {
				fmt.Printf("Error unboxing %v: %v", arg, err)
				continue
			}

			fmt.Printf("\nSuccessfully unboxed %v\n", arg)
		}
	},
}

func init() {
	rootCmd.AddCommand(getCmd)

	// Here you will define your flags and configuration settings.

	getCmd.Flags().String("host", "localhost", "Specify the host to connect")
	getCmd.Flags().Int("port", 5950, "Specify the port to use")
	getCmd.Flags().StringP("output", "o", "", "Specify the output directory (default Current Directory)")
}
