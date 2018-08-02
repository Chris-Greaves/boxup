// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
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

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a new Box to the Server",
	Long: `Calling this command will request a new Box to be added to the Server.
	
	NOTE: The path is relative to the server location, not where this command is being called.`,
	Args: func(cmd *cobra.Command, args []string) error {
		nameFlag := cmd.Flag("name")
		pathFlag := cmd.Flag("path")
		if nameFlag.Value.String() == "" || pathFlag.Value.String() == "" {
			return errors.New("You must specify both --name and --path")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		host := cmd.Flag("host").Value.String()
		port := cmd.Flag("port").Value.String()
		name := cmd.Flag("name").Value.String()
		path := cmd.Flag("path").Value.String()
		fmt.Printf("host:%v\n", host)
		fmt.Printf("port:%v\n", port)
		fmt.Printf("name:%v\n", name)
		fmt.Printf("path:%v\n", path)

		err := addBox(host, port, name, path)
		if err != nil {
			fmt.Printf("Error adding new Box: %v", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.
	addCmd.Flags().StringP("name", "n", "", "Specify thw name of new Box")
	addCmd.Flags().StringP("path", "p", "", "Specify the path of the Box")
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
